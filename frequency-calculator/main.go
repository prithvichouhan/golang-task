package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	log.Fatal("Stopped ", http.ListenAndServe(":4000", r))

}

func main() {
	initializeRouter()
}
