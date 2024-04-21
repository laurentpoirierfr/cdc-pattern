package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Customer struct {
	CustomerID int32
	StoreID    int16
	FirstName  string
	LastName   string
	Email      pgtype.Text

	// AddressID  int16
	Address Address

	Activebool bool
	CreateDate pgtype.Date
	LastUpdate pgtype.Timestamp
	Active     pgtype.Int4
}

type Address struct {
	AddressID int32
	Address   string
	Address2  pgtype.Text
	District  string

	// CityID     int16
	City City

	PostalCode pgtype.Text
	Phone      string
	LastUpdate pgtype.Timestamp
}

type City struct {
	CityID int32
	City   string

	// CountryID  int16
	Country Country

	LastUpdate pgtype.Timestamp
}

type Country struct {
	CountryID  int32
	Country    string
	LastUpdate pgtype.Timestamp
}

// type Actor struct {
// 	ActorID    int32
// 	FirstName  string
// 	LastName   string
// 	LastUpdate pgtype.Timestamp
// }

// type Category struct {
// 	CategoryID int32
// 	Name       string
// 	LastUpdate pgtype.Timestamp
// }

// type Film struct {
// 	FilmID          int32
// 	Title           string
// 	Description     pgtype.Text
// 	ReleaseYear     interface{}
// 	LanguageID      int16
// 	RentalDuration  int16
// 	RentalRate      pgtype.Numeric
// 	Length          pgtype.Int2
// 	ReplacementCost pgtype.Numeric
// 	Rating          interface{}
// 	LastUpdate      pgtype.Timestamp
// 	SpecialFeatures []string
// 	Fulltext        interface{}
// }

// type FilmActor struct {
// 	ActorID    int16
// 	FilmID     int16
// 	LastUpdate pgtype.Timestamp
// }

// type FilmCategory struct {
// 	FilmID     int16
// 	CategoryID int16
// 	LastUpdate pgtype.Timestamp
// }

// type Inventory struct {
// 	InventoryID int32
// 	FilmID      int16
// 	StoreID     int16
// 	LastUpdate  pgtype.Timestamp
// }

// type Language struct {
// 	LanguageID int32
// 	Name       string
// 	LastUpdate pgtype.Timestamp
// }

// type Payment struct {
// 	PaymentID   int32
// 	CustomerID  int16
// 	StaffID     int16
// 	RentalID    int32
// 	Amount      pgtype.Numeric
// 	PaymentDate pgtype.Timestamp
// }

// type Rental struct {
// 	RentalID    int32
// 	RentalDate  pgtype.Timestamp
// 	InventoryID int32
// 	CustomerID  int16
// 	ReturnDate  pgtype.Timestamp
// 	StaffID     int16
// 	LastUpdate  pgtype.Timestamp
// }

// type Staff struct {
// 	StaffID    int32
// 	FirstName  string
// 	LastName   string
// 	AddressID  int16
// 	Email      pgtype.Text
// 	StoreID    int16
// 	Active     bool
// 	Username   string
// 	Password   pgtype.Text
// 	LastUpdate pgtype.Timestamp
// 	Picture    []byte
// }

// type Store struct {
// 	StoreID        int32
// 	ManagerStaffID int16
// 	AddressID      int16
// 	LastUpdate     pgtype.Timestamp
// }
