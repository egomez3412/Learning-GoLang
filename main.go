package main

import ( 
	"fmt"
)

/*
Project: Find out which country has the cheapest subscription for an online services with the best benefits (Netflix, Spotify, etc.)
 
- some services may contain less content in some countries than others
- A VPN may be required to access the content
- the content may be specific to the country

*/

// Phase 1: Learn how to use Go using this project as a guide

// a struct for locations where a subscription for an online service is cheaper at a particular country
// some properties include the country, the service, price, currency, conversion rate, and the price in USD
type subscription struct {
	serviceName string
	country string
	price float64
	currency string
	conversionRate float64
	priceInUSD float64
}

type service struct {
	name string
	costPerMonth float64
	catalogSize int
}

func main() {
	// create a subscription: all values are zero values (i.e. service name is empty string, price is 0, etc.)
	var sub1 subscription
	
	// set the values of the struct
	sub1.serviceName = "Netflix"	

	// print the struct; result: {Netflix  0  0 0}
	fmt.Println(sub1)

	// print the struct with the values; result: {serviceName:Netflix country: price:0 currency: conversionRate:0 priceInUSD:0}
	fmt.Printf("%+v\n", sub1)
}