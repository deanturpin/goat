package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)

func main() {

	// Step 1: Make the HTTP GET request
	resp, err := http.Get("https://min-api.cryptocompare.com/data/price?fsym=BTC&tsyms=USD")
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}
	defer resp.Body.Close() // Always close the response body

	// Step 2: Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	fmt.Println(string(body))

	// print price
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	fmt.Println("Price of BTC in USD: ", result["USD"])
}
