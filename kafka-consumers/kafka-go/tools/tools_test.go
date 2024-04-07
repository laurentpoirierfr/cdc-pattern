package tools_test

import (
	"kafka-go-consumer/tools"
	"testing"
)

func TestNewId(t *testing.T) {
	id := tools.NewId()
	t.Log(id)
}
