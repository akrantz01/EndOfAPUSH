package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

func main() {
	// Setup configuration file
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	// Setup default values
	viper.SetDefault("server.host", "127.0.0.1")
	viper.SetDefault("server.port", 8080)

	// Read configuration
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read configuration file: %s", err)
	}

	// Setup handlers & routes
	router := mux.NewRouter()
	http.Handle("/", handlers.LoggingHandler(os.Stdout, router))

	// Start the server
	log.Printf("Listening on %s:%s...", viper.GetString("server.host"), viper.GetString("server.port"))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port")), nil); err != nil {
		log.Fatalf("Falied to listen on %s:%s: %s", viper.GetString("server.host"), viper.GetString("server.port"), err)
	}
}
