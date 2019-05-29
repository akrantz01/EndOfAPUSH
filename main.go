package main

import (
	"fmt"
	"github.com/akrantz01/EndOfAPUSH/database"
	"github.com/akrantz01/EndOfAPUSH/routes/auth"
	"github.com/akrantz01/EndOfAPUSH/routes/messages"
	"github.com/akrantz01/EndOfAPUSH/routes/user"
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
	viper.SetDefault("database.wipe", true)

	// Read configuration
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read configuration file: %v", err)
	}

	// Setup the database
	db = database.SetupDatabase()

	// Setup handlers & routes
	router := mux.NewRouter()

	// API sub-router
	api := router.PathPrefix("/api").Subrouter()
	// User routes
	api.HandleFunc("/users", user.Create(db)).Methods("POST")
	api.HandleFunc("/users/{username}", user.Delete(db)).Methods("DELETE")
	api.HandleFunc("/users/{username}", user.Update(db)).Methods("PUT")
	api.HandleFunc("/users/{username}", user.Read(db)).Methods("GET")
	api.HandleFunc("/users/search", user.Search(db)).Methods("GET")
	// Authentication routes
	api.HandleFunc("/auth/login", auth.Login(db)).Methods("POST")
	api.HandleFunc("/auth/logout", auth.Logout(db)).Methods("GET")
	// Message routes
	api.HandleFunc("/messages", messages.Create(db)).Methods("POST")
	api.HandleFunc("/messages", messages.List(db)).Methods("GET")
	api.HandleFunc("/messages/{id}", messages.Read(db)).Methods("GET")

	// View route
	router.Handle("/", http.FileServer(http.Dir("./public")))
	http.Handle("/", handlers.LoggingHandler(os.Stdout, router))

	// Start the server
	log.Printf("Listening on %s:%s...", viper.GetString("server.host"), viper.GetString("server.port"))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port")), nil); err != nil {
		log.Fatalf("Falied to listen on %s:%s: %v", viper.GetString("server.host"), viper.GetString("server.port"), err)
	}
}
