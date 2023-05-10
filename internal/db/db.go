package db

import (
	"os"

	"github.com/surrealdb/surrealdb.go"
)

type SurrealResponse []struct {
	Result []map[string]any `json:"result"`
	Status string           `json:"status"`
	Time   string           `json:"time"`
}

func SurrealConnect() *surrealdb.DB {
	db, err := surrealdb.New(os.Getenv("SURREAL_HOST"))

	if err != nil {
		panic(err.Error())
	}

	db.Use(os.Getenv("SURREAL_NS"), os.Getenv("SURREAL_DB"))
	db.Signin(map[string]any{
		"user": os.Getenv("SURREAL_USER"),
		"pass": os.Getenv("SURREAL_PASS"),
	})

	return db
}
