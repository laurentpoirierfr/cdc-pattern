package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type CreditCard struct {
	gorm.Model
	Number string `json:"number"`
	UserID uint   `json:"user_id"`
}

type User struct {
	gorm.Model
	Name        string `json:"name"`
	Email       string `json:"email"`
	CreditCards []CreditCard
}

func init() {
	var err error
	dsn := "postgresql://postgres:postgres@localhost/demo?sslmode=disable" // Update with your database credentials
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

}

func main() {
	// Auto migrate the model
	db.AutoMigrate(&CreditCard{})
	db.AutoMigrate(&User{})

	user := new(User)
	user.Email = "test@gmail.com"
	user.Name = "Test"

	var creditCards = []CreditCard{}
	creditCards = append(creditCards, CreditCard{Number: "132456789"})
	creditCards = append(creditCards, CreditCard{Number: "789546213"})
	user.CreditCards = creditCards

	db.Create(&user)

	// db.Delete(&user, id)

}
