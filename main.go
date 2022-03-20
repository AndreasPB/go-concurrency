package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	startTime := time.Now()
	defer printExecutionTime(startTime)
	getQuotesSeq(100)
}

func printExecutionTime(t time.Time) {
	fmt.Println("Execution time: ", time.Since(t))
}

// --- ChuckNorris --- //
type ChuckNorris struct {
	Quote      string   `json:"value"`
	Url        string   `json:"url"`
	Id         string   `json:"id"`
	Categories []string `json:"categories"`
}

const baseUrl = "https://api.chucknorris.io/jokes/random"

func getQuote() (quote *ChuckNorris, err error) {
	response, err := http.Get(baseUrl)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(response.Body).Decode(&quote)
	if err != nil {
		return nil, err
	}

	return quote, nil
}

func getQuotesSeq(numOfQuotes int) {
	quotesMap := make(map[int]*ChuckNorris, numOfQuotes)

	for i := 0; i < numOfQuotes; i++ {
		quote, err := getQuote()

		if err != nil {
			continue
		}

		quotesMap[i] = quote
		fmt.Println("New quote: ", quote.Quote)
	}

}
