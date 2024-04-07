package kafka

import (
	"context"
	"fmt"
	"kafka-go-consumer/models"
	"kafka-go-consumer/tools"
	"log"
	"net"
	"strconv"

	"github.com/segmentio/kafka-go"
)

const (
	FIRST_OFFSET = kafka.FirstOffset
	LAST_OFFSET  = kafka.LastOffset
)

func CreateTopic(kafkaURL, topic string, partitions, replications int) {
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

func Producer(kafkaURL, topic string) {
	fmt.Println("start producing ... !!")

	kafkaWriter := kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
		//AllowAutoTopicCreation: true,
	}

	defer kafkaWriter.Close()

	kafkaWriter.Async = true

	for {
		headers := []kafka.Header{
			{Key: "business.client-id", Value: []byte(tools.NewId())},
		}

		item := models.NewItem().Json()
		msg := kafka.Message{
			Key: []byte(tools.NewId()),
			// Value:   []byte(fmt.Sprintf("%v", models.NewItem())),
			Value:   item,
			Headers: headers,
		}
		err := kafkaWriter.WriteMessages(context.Background(), msg)
		if err != nil {
			log.Fatalln(err)
		}

		tools.TraceMessage("produce :", msg)
		// time.Sleep(1 * time.Second)
	}

}

func Consumer(kafkaURL, topic string, offset int64) {
	fmt.Println("start consuming ... !!")

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaURL},
		// GroupID:     groupID,
		Topic:     topic,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
		Partition: 0,
	})
	defer reader.Close()

	reader.SetOffset(offset)
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		tools.TraceMessage("consume :", msg)
	}
}
