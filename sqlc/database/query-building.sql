-- name: GetCustomer :one
SELECT * FROM customer WHERE customer_id=$1 LIMIT 1;

-- name: GetCustomers :many
SELECT * FROM customer LIMIT $1 OFFSET $2;

-- name: GetAddress :one
SELECT * FROM address WHERE address_id=$1 LIMIT 1;

-- name: GetCity :one
SELECT * FROM city WHERE city_id=$1 LIMIT 1;

-- name: GetCountry :one
SELECT * FROM country WHERE country_id=$1 LIMIT 1;

