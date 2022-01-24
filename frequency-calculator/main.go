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

	r.HandleFunc("/api/v1/wordsCount", handler.CalculateWordsFrequency).Methods("POST")

	log.Printf("Server listening on port: %s", os.Getenv("PORT"))
	log.Fatal("Server has stopped", http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r))
}

func main() {
	initializeRouter()
}
