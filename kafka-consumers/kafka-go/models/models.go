package models

import (
	"encoding/json"
	"kafka-go-consumer/tools"

	"github.com/bxcodec/faker"
)

type Item struct {
	Id          string
	Name        string  `faker:"name"`
	Description string  `faker:"paragraph"`
	Price       float64 `faker:"amount"`
}

func NewItem() *Item {
	item := Item{}
	faker.FakeData(&item)
	item.Id = tools.NewId()
	return &item
}

func (item *Item) Json() []byte {
	j, _ := json.Marshal(item)
	return j
}
