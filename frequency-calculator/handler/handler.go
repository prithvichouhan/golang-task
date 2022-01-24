package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"

	errors "frequency-calculator/shared/error"
)

// Function to validate input
func validateInput(input string) (string, error) {
	if input == "" {
		return "", errors.ErrInputNotEmpty
	}

	return input, nil
}

// Function to set header
func printHeader(w http.ResponseWriter, code int, result interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	post, _ := json.Marshal(result)
	fmt.Fprintf(w, "%v", string(post))
}
