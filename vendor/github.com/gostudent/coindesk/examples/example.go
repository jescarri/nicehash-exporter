package main

import (
	"fmt"
	"github.com/gostudent/coindesk"
)

func main() {
	priceUsd := coindesk.GetPrice()
	fmt.Println(priceUsd)

	priceInr := coindesk.GetPrice("INR")
	fmt.Println(priceInr)

	USD, GBP, EUR := coindesk.CurrentPrice()
	fmt.Println(USD, GBP, EUR)
	
	historyPrice := coindesk.HistoryPrice("2013-09-01", "2013-09-05")
	fmt.Println("Date\t\tPrice")
	for _, i := range historyPrice {
		fmt.Print(i.Date.Format("2006-01-02"), "\t")
		fmt.Println(i.Price)
	}

	yesterdayPrice := coindesk.Yesterday()
	fmt.Println(yesterdayPrice)
}


