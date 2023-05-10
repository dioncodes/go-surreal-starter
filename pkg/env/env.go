package env

import (
	"time"

	"github.com/dioncodes/go-surreal-starter/internal/db"
	"github.com/surrealdb/surrealdb.go"
)

type Env struct {
	DB *surrealdb.DB
}

func Init() *Env {
	return &Env{
		DB: db.SurrealConnect(),
	}
}

func (env *Env) destructDb() {
	env.DB.Close()
}

// send a ping to the database to check if it is reachable
func (env *Env) CheckDbConnection() {
	c1 := make(chan string, 1)

	go func() {
		_, err := env.DB.Query("INFO FOR DB;", map[string]any{})
		if err != nil {
			// log.Println("Connection to database lost. Reconnecting", err)
			env.destructDb()
			env.DB = db.SurrealConnect()
		}
		c1 <- "ping"
	}()

	select {
	case <- c1:
		// nothing to do
	case <- time.After(time.Second * 15):
		// log.Println("Timeout on db.Ping - killing connection and reconnecting")
		env.destructDb()
		env.DB = db.SurrealConnect()
	}
}
