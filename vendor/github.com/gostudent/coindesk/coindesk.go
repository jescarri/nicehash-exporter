package coindesk

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Price struct {
	Date time.Time
	Price float64
}


// GetPrice fetches the latest Bitcoin Price from coindesk API in given currency.
// s is the ISO code of the currency. Defaults to USD.
// Supported Currencies - https://api.coindesk.com/v1/bpi/supported-currencies.json
func GetPrice(s ...string) float64 {
	curr := "USD"
	if len(s) > 0 {
		curr = strings.ToUpper(s[0])
	}
	url := fmt.Sprintf("https://api.coindesk.com/v1/bpi/currentPrice/%s.json", curr)

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	body := string(responseData)
	data := fmt.Sprintf("bpi.%s.rate_float", curr)
	Price := gjson.Get(body, data).String()
	output, err := strconv.ParseFloat(Price, 64)
	if err != nil {
		log.Fatal(err)
	}

	return output
}

// HistoryPrice takes in startDate and endDate as input
// It returns an array of Price struct
func HistoryPrice(startDate string, endDate string) []Price {
	url := "https://api.coindesk.com/v1/bpi/historical/close.json?start="+startDate+"&end="+endDate

	var history []Price //Array of struct containing Prices from startDate to endDate

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	body := string(responseData)

	var m map[string]gjson.Result
	m = gjson.Get(body, "bpi").Map()
	layout := "2006-01-02"

	for t, p := range(m) {
		Date, err := time.Parse(layout, t)
		if err != nil {
			fmt.Println(err)
		}
		history = append(history, Price{Date, p.Num})
	}

	return history
}

// CurrentPrice fetches current Price of Bitcoin from coindesk API.
// It returns the Price in USD, GBP and EUR.
func CurrentPrice() (float64, float64, float64) {
	url := "https://api.coindesk.com/v1/bpi/currentPrice.json"

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	body := string(responseData)
	priceUsd := gjson.Get(body, "bpi.USD.rate_float").String()
	USD, err := strconv.ParseFloat(priceUsd, 64)
	if err != nil {
		log.Fatal(err)
	}
	priceGbp := gjson.Get(body, "bpi.GBP.rate_float").String()
	GBP, err := strconv.ParseFloat(priceGbp, 64)
	if err != nil {
		log.Fatal(err)
	}
	priceEur := gjson.Get(body, "bpi.EUR.rate_float").String()
	EUR, err := strconv.ParseFloat(priceEur, 64)
	if err != nil {
		log.Fatal(err)
	}

	return USD, GBP, EUR
}

// Yesterday returns yesterday's Price
func Yesterday() float64{
	url := "https://api.coindesk.com/v1/bpi/historical/close.json?for=yesterday"

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	body := string(responseData)
	var m map[string]gjson.Result
	m = gjson.Get(body, "bpi").Map()
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range(m) {
		return p.Num //Returns the yesterdays Price
	}

	return 0
}
