package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName   string       `json:"first_name"`
	LastName    string       `json:"last_name"`
	Address     string       `json:"address"`
	Email       string       `json:"email"`
	CreditCards []CreditCard `json:"credit_cards" gorm:"foreignKey:user_id"`
	Orders      []Order      `json:"orders gorm:"foreignKey:user_id"`
}

type CreditCard struct {
	gorm.Model
	Number string `json:"number"`
	UserID uint   `json:"user_id"`
}

type Order struct {
	gorm.Model
	TotalPaid   float64   `json:"total_paid"`
	Products    []Product `json:"products" gorm:"foreignKey:order_id"`
	OperationAt time.Time `json:"op_at"`
	UserID      uint      `json:"user_id"`
}

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	OrderId     uint    `json:"order_id"`
}

func NewUser(firstName, lastName, address, email string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Address:   address,
		Email:     email,
	}
}
