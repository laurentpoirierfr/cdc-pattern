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


-- name: GetCustomersByCountry :many
SELECT customer.* 
FROM customer, address, city, country
WHERE customer.address_id = address.address_id 
AND address.city_id = city.city_id
AND city.country_id = country.country_id AND country.country_id = $1 ;


-- name: GetCustomersByCity :many
SELECT customer.* 
FROM customer, address, city, country
WHERE customer.address_id = address.address_id 
AND address.city_id = city.city_id
AND city.city_id = $1 ;


-- name: GetCustomersByAddress :many
SELECT customer.* 
FROM customer, address, city, country
WHERE customer.address_id = address.address_id 
AND address.address_id  = $1 ;
