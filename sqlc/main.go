package main

import (
	"context"
	"log"
	"os"
	repo "sqlc-demo/repositories"

	"github.com/jackc/pgx/v5"
)

//go:generate sqlc generate

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	OnErr(err)

	defer conn.Close(ctx)

	queries := repo.New(conn)

	arg := repo.GetCustomersParams{
		Limit:  10,
		Offset: 1,
	}

	address, err := queries.GetCustomers(ctx, arg)
	OnErr(err)

	log.Println(address)

}

func OnErr(err error) {
	if err != nil {
		panic(err)
	}
}
