package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Response struct {
	Success     bool   `json:"success"`
	LowestPrice string `json:"lowest_price"`
	Volume      string `json:"volume"`
	MedianPrice string `json:"median_price"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	message := "This HTTP triggered function executed successfully. Pass a name in the query string for a personalized response.\n"
	name := r.URL.Query().Get("name")
	if name != "" {
		message = fmt.Sprintf("Hello, %s. This HTTP triggered function executed successfully.\n", name)
	}
	fmt.Fprint(w, message)
}

func steamHandler(w http.ResponseWriter, r *http.Request) {
	message := "Steam trigger has been executed successfully"
	const itemUrl = "https://steamcommunity.com/market/priceoverview/?appid=730&currency=3&market_hash_name=Revolution%20Case"

	c := http.Client{Timeout: time.Duration(1) * time.Second}

	resp, err := c.Get(itemUrl)

	if err != nil {
		fmt.Printf("Error %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var response Response
	json.Unmarshal([]byte(body), &response)

	lowestPriceFloat, err := strconv.ParseFloat(normalizeGermanFloatString(strings.TrimSuffix(response.LowestPrice, "€")), 32)

	message = fmt.Sprintf("Revolution case a current supply of %s and the lowest price is %f€", response.Volume, lowestPriceFloat)

	fmt.Fprint(w, message)
}

func normalizeGermanFloatString(old string) string {
	s := strings.Replace(old, ",", ".", -1)
	s = strings.Replace(s, "--", "00", -1)
	return strings.Replace(s, ".", "", 1)
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/sample", helloHandler)
	http.HandleFunc("/api/steam", steamHandler)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
