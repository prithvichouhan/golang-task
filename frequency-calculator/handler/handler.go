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
