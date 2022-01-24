package main

import (
	"log"

	"github.com/joho/godotenv"
)
// Function init
func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error: Unable to load env")
	}
}
