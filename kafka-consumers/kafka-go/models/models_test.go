package models_test

import (
	"kafka-go-consumer/models"
	"testing"
)

func TestNewItem(t *testing.T) {
	item := models.NewItem()
	t.Log(item)
}
