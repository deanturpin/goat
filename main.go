package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)

func main() {

	// Make the HTTP GET request
	resp, err := http.Get("https://min-api.cryptocompare.com/data/v2/histoday?fsym=BTC&tsym=USD&limit=3")
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}
	defer resp.Body.Close() // Always close the response body

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)	
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Pretty print json
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	prettyJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatalf("Failed to generate pretty json: %v", err)
	}
	fmt.Println(string(prettyJSON))
}
