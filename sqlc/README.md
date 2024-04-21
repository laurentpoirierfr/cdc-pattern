# README

```bash
curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET "http://localhost:7070/api/customers/?offset=0&limit=500000"
curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET "http://localhost:7070/api/customers/1"


curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET "http://localhost:7070/api/countries/2"

curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET "http://localhost:7070/api/customers/address/cities/countries/2"

```


## References

* https://www.postgresqltutorial.com/postgresql-getting-started/load-postgresql-sample-database/
* https://www.postgresqltutorial.com/postgresql-getting-started/postgresql-sample-database/