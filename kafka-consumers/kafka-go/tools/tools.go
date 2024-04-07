package tools

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

func TraceMessage(title string, msg kafka.Message) {
	j, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(title, string(j))
}

func NewId() string {
	return uuid.New().String()
}
