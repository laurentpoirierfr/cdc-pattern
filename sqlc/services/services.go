package services

import (
	"context"
	"os"
	"sqlc-demo/models"
	repo "sqlc-demo/repositories"

	"github.com/jackc/pgx/v5"
)

type Service struct {
	ctx     context.Context
	conn    *pgx.Conn
	queries *repo.Queries
}

func NewService() (*Service, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return &Service{}, err
	}

	queries := repo.New(conn)
	return &Service{
		ctx:     ctx,
		conn:    conn,
		queries: queries,
	}, nil
}

func (s *Service) GetCustomer(id int32) (models.Customer, error) {
	var repoCustomer repo.Customer
	repoCustomer, err := s.queries.GetCustomer(s.ctx, id)
	if err != nil {
		return models.Customer{}, err
	}

	address, err := s.GetAddress(int32(repoCustomer.AddressID))
	if err != nil {
		return models.Customer{}, err
	}

	return models.Customer{
		CustomerID: repoCustomer.CustomerID,
		StoreID:    repoCustomer.StoreID,
		FirstName:  repoCustomer.FirstName,
		LastName:   repoCustomer.LastName,
		Email:      repoCustomer.Email,
		Address:    address,
		Activebool: repoCustomer.Activebool,
		CreateDate: repoCustomer.CreateDate,
		LastUpdate: repoCustomer.LastUpdate,
		Active:     repoCustomer.Active,
	}, nil
}

func (s *Service) GetAddress(id int32) (models.Address, error) {
	var repoAddress repo.Address
	repoAddress, err := s.queries.GetAddress(s.ctx, id)
	if err != nil {
		return models.Address{}, err
	}

	city, err := s.GetCity(int32(repoAddress.CityID))
	if err != nil {
		return models.Address{}, err
	}

	return models.Address{
		AddressID:  repoAddress.AddressID,
		Address:    repoAddress.Address,
		Address2:   repoAddress.Address2,
		District:   repoAddress.District,
		City:       city,
		PostalCode: repoAddress.PostalCode,
		Phone:      repoAddress.Phone,
		LastUpdate: repoAddress.LastUpdate,
	}, nil
}

func (s *Service) GetCity(id int32) (models.City, error) {
	var repoCity repo.City
	repoCity, err := s.queries.GetCity(s.ctx, id)
	if err != nil {
		return models.City{}, err
	}

	country, err := s.GetCountry(int32(repoCity.CountryID))
	if err != nil {
		return models.City{}, err
	}

	return models.City{
		CityID:     repoCity.CityID,
		City:       repoCity.City,
		Country:    country,
		LastUpdate: repoCity.LastUpdate,
	}, nil
}

func (s *Service) GetCountry(id int32) (models.Country, error) {
	var repoCountry repo.Country
	repoCountry, err := s.queries.GetCountry(s.ctx, id)
	if err != nil {
		return models.Country{}, err
	}
	return models.Country{
		CountryID:  repoCountry.CountryID,
		Country:    repoCountry.Country,
		LastUpdate: repoCountry.LastUpdate,
	}, nil
}
