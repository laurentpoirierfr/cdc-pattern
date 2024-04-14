package main

import (
	"database/sql"
	"log"

	_ "github.com/sijms/go-ora"
)

func main() {

	// connectString format: [hostname]:[port]/[DB service name]

	dsn := `user="system" password="password" connectString="localhost:1521/FREE"`
	db, err := sql.Open("oracle", dsn)
	if err != nil {
		panic(err)
	}
	db.Close()
	log.Println("Connection localhost:1521/FREE user <system>  Ok !")

	dsn = `user="demo" password="demo" connectString="localhost:1521/FREE"`
	db, err = sql.Open("oracle", dsn)
	if err != nil {
		panic(err)
	}
	db.Close()
	log.Println("Connection localhost:1521/FREE user <demo>  Ok !")

	dsn = `user="system" password="password" connectString="localhost:1521/FREEPDB1"`
	db, err = sql.Open("oracle", dsn)
	if err != nil {
		panic(err)
	}
	log.Println("Connection localhost:1521/FREEPDB1 user <system> Ok !")
	db.Close()

	dsn = `user="demo" password="demo" connectString="localhost:1521/FREEPDB1"`
	db, err = sql.Open("oracle", dsn)
	if err != nil {
		panic(err)
	}
	log.Println("Connection localhost:1521/FREEPDB1 user <demo> Ok !")
	db.Close()

}
