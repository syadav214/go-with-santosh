package config

import (
	"log"
	"os"

	"github.com/go-pg/pg/v9"
	controllers "github.com/syadav214/todo/controllers"
)

// Connecting to db
func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "sekrit",
		Addr:     "localhost:5432",
		Database: "godb",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	controllers.CreateTodoTable(db)
	controllers.InitiateDB(db)
	return db
}
