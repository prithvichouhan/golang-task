package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Function init
func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error: Unable to load env")
	}
}

// Function to get the top ten element
func getTopTenElement(input string) {
	//Encode the data
	postBody, _ := json.Marshal(map[string]string{
		"input": input,
	})

	body := bytes.NewBuffer(postBody)
	resp, err := http.Post(os.Getenv("URL"), "application/json", body)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	fmt.Println("Top Ten element - ", string(responseData))
}

// Function Main
func main() {
	content, err := ioutil.ReadFile("data/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	getTopTenElement(string(content))
}
