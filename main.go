package main

import (
	"fmt"
	"github.com/akrantz01/EndOfAPUSH/database"
	"github.com/akrantz01/EndOfAPUSH/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

var db *gorm.DB

func main() {
	// Setup configuration file
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	// Setup default values
	viper.SetDefault("server.host", "127.0.0.1")
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("database.host", "127.0.0.1")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.ssl", "disable")
	viper.SetDefault("database.username", "postgres")
	viper.SetDefault("database.password", "postgres")
	viper.SetDefault("database.database", "postgres")

	// Read configuration
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read configuration file: %v", err)
	}

	// Setup the database
	db = database.SetupDatabase()

	// Setup handlers & routes
	router := mux.NewRouter()
	router.HandleFunc("/users", routes.Create(db)).Methods("POST")
	router.HandleFunc("/users/{username}", routes.Delete(db)).Methods("DELETE")
	router.HandleFunc("/users/search", routes.Search(db)).Methods("GET")
	router.HandleFunc("/users/login", routes.Login(db)).Methods("POST")
	router.HandleFunc("/users/logout", routes.Logout(db)).Methods("GET")
	http.Handle("/", handlers.LoggingHandler(os.Stdout, router))

	// Start the server
	log.Printf("Listening on %s:%s...", viper.GetString("server.host"), viper.GetString("server.port"))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port")), nil); err != nil {
		log.Fatalf("Falied to listen on %s:%s: %v", viper.GetString("server.host"), viper.GetString("server.port"), err)
	}
}
