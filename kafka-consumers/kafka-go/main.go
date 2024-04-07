package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"kafka-go-consumer/kafka"

	"github.com/joho/godotenv"
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
	kafka.CreateTopic(kafkaURL, topic, 1, 1)

	done := make(chan struct{})

	go func() {
		log.Println("Listening signals...")
		c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		close(done)
	}()

	for i := 0; i < 5; i++ {
		go kafka.Producer(kafkaURL, topic)
	}

	for i := 0; i < 10; i++ {
		go kafka.Consumer(kafkaURL, topic, kafka.FIRST_OFFSET)
	}

	<-done

	log.Println("Done.")
}
