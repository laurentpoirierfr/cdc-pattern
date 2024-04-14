package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/sijms/go-ora"
)

type Account struct {
	User     string
	Password string
}

func main() {

	accounts := []Account{
		{
			User:     "system",
			Password: "password",
		},
		{
			User:     "demo",
			Password: "demo",
		},
	}

	dbServices := []string{"FREE", "FREEPDB1"}

	for _, dbService := range dbServices {
		for _, account := range accounts {
			// connectString format: [hostname]:[port]/[DB service name]
			dsn := fmt.Sprintf("user=%s password=%s connectString=localhost:1521/%s", account.User, account.Password, dbService)
			db, err := sql.Open("oracle", string(dsn))
			if err != nil {
				panic(err)
			}
			db.Close()
			msg := fmt.Sprintf("Connection localhost:1521/%s user <%s>  Ok !", dbService, account.User)
			log.Println(msg)
		}
	}

}
