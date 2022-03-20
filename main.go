package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	defer printExecutionTime(startTime)
	getQuotesConcurrently(100)
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

func getQuotesSequencially(numOfQuotes int) {
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

func getQuotesConcurrently(numOfQuotes int) {
	var quotesMap sync.Map
	wg := sync.WaitGroup{}

	for i := 0; i < numOfQuotes; i++ {
		wg.Add(1)
		go func(idx int) {
			quote, err := getQuote()

			if err != nil {
				panic(err)
			}

			quotesMap.Store(idx, quote)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
