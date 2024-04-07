package models

import (
	"encoding/json"
	"kafka-go-consumer/tools"

	"github.com/bxcodec/faker"
)

type Item struct {
	Id          string  `json:"id"`
	Name        string  `faker:"name" json:"name"`
	Description string  `faker:"paragraph" json:"description"`
	Price       float64 `faker:"amount" json:"price"`
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
