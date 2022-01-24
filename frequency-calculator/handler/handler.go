package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
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

// Function to get word count
func GetWordCounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body request
	_ = json.NewDecoder(r.Body).Decode(&body)

	_, err := validateInput(body.Input)
	if err != nil {
		printHeader(w, http.StatusBadRequest, err)
		return
	}

	result := sortElements(parseInputContent(body.Input))

	json.NewEncoder(w).Encode(result)
}

// Function to parse input string
func parseInputContent(fileContents string) map[string]int {
	re := regexp.MustCompile(`(?m)[\pL_]`)
	wCol := strings.Split(string(fileContents), " ")
	wordMap := make(map[string]int)
	for _, word := range wCol {
		if re.MatchString(word) {
			if val, found := wordMap[word]; found {
				wordMap[word] = val + 1
			} else {
				wordMap[word] = 1
			}
		}
	}

	return wordMap
}

// Funcion to sort the elements
func sortElements(wordsMap map[string]int) []wordFrequency {
	var sortSlice []wordFrequency
	for k, v := range wordsMap {
		sortSlice = append(sortSlice, wordFrequency{k, v})
	}

	sort.Slice(sortSlice, func(i int, j int) bool {
		return sortSlice[i].Frequency > sortSlice[j].Frequency
	})

	if len(sortSlice) > 10 {
		sortSlice = sortSlice[:10]
	}

	return sortSlice
}
