package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	handler "frequency-calculator/handler"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error: Unable to load env")
	}
}

func initializeRouter() {
	r := mux.NewRouter()

	log.Fatal("Stopped ", http.ListenAndServe(":4000", r))

}

func main() {
	initializeRouter()
}
