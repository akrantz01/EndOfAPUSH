package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"log"
)

// Stores the user's information in a PostgreSQL table
type User struct {
	gorm.Model
	Name     string
	Username string
	Password string
	Inbox    []Message `gorm:"foreignkey:ToID"`
	Outbox   []Message `gorm:"foreignkey:FromID"`
}

// Stores a message in a PostgreSQL table
type Message struct {
	gorm.Model
	Subject   string
	Message   string
	Algorithm uint
	ToID      uint
	FromID    uint
}

// Setup the database by connecting to it and building the schema if it does not already exist
func setupDatabase() *gorm.DB {
	log.Print("Connecting to database...")
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", viper.GetString("database.host"), viper.GetString("database.port"), viper.GetString("database.username"), viper.GetString("database.password"), viper.GetString("database.database"), viper.GetString("database.ssl")))
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	log.Print("Successfully connected to database")

	log.Print("Building database schema...")
	for _, model := range []interface{}{&User{}, &Message{}} {
		if !db.HasTable(model) {
			db.CreateTable(model)
		}
	}
	log.Print("Successfully built database schema")

	return db
}
