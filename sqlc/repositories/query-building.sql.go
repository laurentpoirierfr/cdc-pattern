// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: query-building.sql

package repo

import (
	"context"
)

const getAddress = `-- name: GetAddress :one
SELECT address_id, address, address2, district, city_id, postal_code, phone, last_update FROM address WHERE address_id=$1 LIMIT 1
`

func (q *Queries) GetAddress(ctx context.Context, addressID int32) (Address, error) {
	row := q.db.QueryRow(ctx, getAddress, addressID)
	var i Address
	err := row.Scan(
		&i.AddressID,
		&i.Address,
		&i.Address2,
		&i.District,
		&i.CityID,
		&i.PostalCode,
		&i.Phone,
		&i.LastUpdate,
	)
	return i, err
}

const getCity = `-- name: GetCity :one
SELECT city_id, city, country_id, last_update FROM city WHERE city_id=$1 LIMIT 1
`

func (q *Queries) GetCity(ctx context.Context, cityID int32) (City, error) {
	row := q.db.QueryRow(ctx, getCity, cityID)
	var i City
	err := row.Scan(
		&i.CityID,
		&i.City,
		&i.CountryID,
		&i.LastUpdate,
	)
	return i, err
}

const getCountry = `-- name: GetCountry :one
SELECT country_id, country, last_update FROM country WHERE country_id=$1 LIMIT 1
`

func (q *Queries) GetCountry(ctx context.Context, countryID int32) (Country, error) {
	row := q.db.QueryRow(ctx, getCountry, countryID)
	var i Country
	err := row.Scan(&i.CountryID, &i.Country, &i.LastUpdate)
	return i, err
}

const getCustomer = `-- name: GetCustomer :one
SELECT customer_id, store_id, first_name, last_name, email, address_id, activebool, create_date, last_update, active FROM customer WHERE customer_id=$1 LIMIT 1
`

func (q *Queries) GetCustomer(ctx context.Context, customerID int32) (Customer, error) {
	row := q.db.QueryRow(ctx, getCustomer, customerID)
	var i Customer
	err := row.Scan(
		&i.CustomerID,
		&i.StoreID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.AddressID,
		&i.Activebool,
		&i.CreateDate,
		&i.LastUpdate,
		&i.Active,
	)
	return i, err
}

const getCustomers = `-- name: GetCustomers :many
SELECT customer_id, store_id, first_name, last_name, email, address_id, activebool, create_date, last_update, active FROM customer LIMIT $1 OFFSET $2
`

type GetCustomersParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) GetCustomers(ctx context.Context, arg GetCustomersParams) ([]Customer, error) {
	rows, err := q.db.Query(ctx, getCustomers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Customer
	for rows.Next() {
		var i Customer
		if err := rows.Scan(
			&i.CustomerID,
			&i.StoreID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.AddressID,
			&i.Activebool,
			&i.CreateDate,
			&i.LastUpdate,
			&i.Active,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCustomersByAddress = `-- name: GetCustomersByAddress :many
SELECT customer.customer_id, customer.store_id, customer.first_name, customer.last_name, customer.email, customer.address_id, customer.activebool, customer.create_date, customer.last_update, customer.active 
FROM customer, address, city, country
WHERE customer.address_id = address.address_id 
AND address.address_id  = $1
`

func (q *Queries) GetCustomersByAddress(ctx context.Context, addressID int32) ([]Customer, error) {
	rows, err := q.db.Query(ctx, getCustomersByAddress, addressID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Customer
	for rows.Next() {
		var i Customer
		if err := rows.Scan(
			&i.CustomerID,
			&i.StoreID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.AddressID,
			&i.Activebool,
			&i.CreateDate,
			&i.LastUpdate,
			&i.Active,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCustomersByCity = `-- name: GetCustomersByCity :many
SELECT customer.customer_id, customer.store_id, customer.first_name, customer.last_name, customer.email, customer.address_id, customer.activebool, customer.create_date, customer.last_update, customer.active 
FROM customer, address, city, country
WHERE customer.address_id = address.address_id 
AND address.city_id = city.city_id
AND city.city_id = $1
`

func (q *Queries) GetCustomersByCity(ctx context.Context, cityID int32) ([]Customer, error) {
	rows, err := q.db.Query(ctx, getCustomersByCity, cityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Customer
	for rows.Next() {
		var i Customer
		if err := rows.Scan(
			&i.CustomerID,
			&i.StoreID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.AddressID,
			&i.Activebool,
			&i.CreateDate,
			&i.LastUpdate,
			&i.Active,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCustomersByCountry = `-- name: GetCustomersByCountry :many
SELECT customer.customer_id, customer.store_id, customer.first_name, customer.last_name, customer.email, customer.address_id, customer.activebool, customer.create_date, customer.last_update, customer.active 
FROM customer, address, city, country
WHERE customer.address_id = address.address_id 
AND address.city_id = city.city_id
AND city.country_id = country.country_id AND country.country_id = $1
`

func (q *Queries) GetCustomersByCountry(ctx context.Context, countryID int32) ([]Customer, error) {
	rows, err := q.db.Query(ctx, getCustomersByCountry, countryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Customer
	for rows.Next() {
		var i Customer
		if err := rows.Scan(
			&i.CustomerID,
			&i.StoreID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.AddressID,
			&i.Activebool,
			&i.CreateDate,
			&i.LastUpdate,
			&i.Active,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
