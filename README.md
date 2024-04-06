# README


```sql
CREATE TABLE customers (id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY, name text);

ALTER TABLE customers REPLICA IDENTITY USING INDEX customers_pkey;

INSERT INTO customers (name) VALUES ('john'), ('jack'), ('jane');

SELECT * FROM customers;

TRUNCATE customers;
```

## References

* https://medium.com/geekculture/listen-to-database-changes-with-apache-kafka-35440a3344f0
