package main

import (
	"bytes"
	"encoding/json"
	"log"
	"sqlc-demo/services"
)

//go:generate sqlc generate

func main() {
	srv, err := services.NewService()
	OnErr(err)
	customer, err := srv.GetCustomer(1)
	OnErr(err)
	pretty, err := PrettyJson(customer)
	OnErr(err)
	log.Println(pretty)
}

func OnErr(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	empty = ""
	tab   = "\t"
)

func PrettyJson(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)

	err := encoder.Encode(data)
	if err != nil {
		return empty, err
	}
	return buffer.String(), nil
}
