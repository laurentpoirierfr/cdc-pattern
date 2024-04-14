package main

//go:generate go-plantuml generate -f model/model.go -o model/graph.puml

import (
	"gorm-demo/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	dsn := "postgresql://postgres:postgres@localhost/demo?sslmode=disable" // Update with your database credentials
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

}

const (
	CREATE = "c"
	UPDATE = "u"
	DELETE = "d"
)

func main() {
	// Auto migrate the model
	db.AutoMigrate(&model.CreditCard{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.Order{})

	user := model.NewUser("first name", "last name", "adress", "test@gmail.com")
	var creditCards = []model.CreditCard{}
	creditCards = append(creditCards, model.CreditCard{Number: "132456789"})
	creditCards = append(creditCards, model.CreditCard{Number: "789546213"})
	user.CreditCards = creditCards

	db.Create(&user)

	user.Address = "nouvelle adresse"

	db.Updates(&user)

	db.Delete(&user, user.ID)

}
