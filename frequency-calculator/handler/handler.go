package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"

	errors "frequency-calculator/shared/error"
)

type request struct {
	Input string
}

type wordFrequency struct {
	Word      string
	Frequency int
}

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

// Function to calculate input words frequency
func CalculateWordsFrequency(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body request
	var topTenEle []wordFrequency
	_ = json.NewDecoder(r.Body).Decode(&body)

	_, err := validateInput(body.Input)
	if err != nil {
		printHeader(w, http.StatusBadRequest, err)
		return
	}

	input := strings.Fields(body.Input)
	wordMap := make(map[string]int)

	for _, v := range input {
		_, ok := wordMap[v]
		if ok {
			wordMap[v] += 1
		} else {
			wordMap[v] = 1
		}
	}

	for k, v := range wordMap {
		topTenEle = append(topTenEle, wordFrequency{k, v})
	}

	sort.Slice(topTenEle, func(i, j int) bool {
		return topTenEle[i].Frequency > topTenEle[j].Frequency
	})

	if len(topTenEle) > 10 {
		topTenEle = topTenEle[:10]
	}

	json.NewEncoder(w).Encode(topTenEle)
}
