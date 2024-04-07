package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	kafka "github.com/segmentio/kafka-go"
)

var (
	kafkaURL string
	topic    string
	groupID  string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get kafka reader using environment variables.
	kafkaURL = os.Getenv("KAFKA_URL")
	topic = os.Getenv("TOPIC_NAME")
	groupID = os.Getenv("GROUP_ID")
}

func main() {
	createTopic(topic, 1, 1)

	done := make(chan struct{})

	go func() {
		log.Println("Listening signals...")
		c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		close(done)
	}()

	go producer()
	go consumer()

	<-done

	log.Println("Done.")
}

func producer() {
	fmt.Println("start producing ... !!")

	kafkaWriter := kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	defer kafkaWriter.Close()

	for i := 0; ; i++ {
		key := fmt.Sprintf("Key-%d", i)
		msg := kafka.Message{
			Key:   []byte(key),
			Value: []byte(fmt.Sprint(uuid.New())),
		}
		err := kafkaWriter.WriteMessages(context.Background(), msg)
		if err != nil {
			log.Fatalln(err)
		}
		traceMessage("produce :", msg)
		// time.Sleep(1 * time.Second)
	}

}

func consumer() {
	fmt.Println("start consuming ... !!")

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaURL},
		// GroupID:     groupID,
		Topic:       topic,
		MinBytes:    10e3, // 10KB
		MaxBytes:    10e6, // 10MB
		StartOffset: kafka.FirstOffset,
		Partition:   0,
	})
	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		traceMessage("consumer :", msg)
	}
}

func createTopic(topic string, partitions, replications int) {
	conn, err := kafka.Dial("tcp", kafkaURL)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{{Topic: topic, NumPartitions: partitions, ReplicationFactor: replications}}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}
}

func traceMessage(title string, msg kafka.Message) {
	j, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(title, string(j))
}
