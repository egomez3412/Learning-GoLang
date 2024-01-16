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
type subscriptionDetails struct {
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
	subscriptionDetails // we do not need to specifiey a field name because we are embedding the struct
}

func main() {
	// create a subscription: all values are zero values (i.e. service name is empty string, price is 0, etc.)
	var sub1 subscriptionDetails
	
	// set the values of the struct
	sub1.serviceName = "Netflix"	

	// print the struct; result: {Netflix  0  0 0}
	fmt.Println(sub1)

	// print the struct with the values; result: {serviceName:Netflix country: price:0 currency: conversionRate:0 priceInUSD:0}
	fmt.Printf("%+v\n", sub1)

	// create a service 
	netflix := service{
		name: "Netflix",
		costPerMonth: 8.99,
		catalogSize: 10000,
		subscriptionDetails: subscriptionDetails{
			country: "USA",
			price: 8.99,
			currency: "USD",
			conversionRate: 1,
			priceInUSD: 8.99,
		},
	}

	// print the service; result: {name:Netflix costPerMonth:8.99 catalogSize:10000 subscriptionDetails:{serviceName: country:USA price:8.99 currency:USD conversionRate:1 priceInUSD:8.99}}
	netflix.print()
	
	netflixPointer := &netflix
	// update the price of the service
	netflixPointer.updatePrice(9.99)

	// print the service; result: {name:Netflix costPerMonth:8.99 catalogSize:10000 subscriptionDetails:{serviceName: country:USA price:8.99 currency:USD conversionRate:1 priceInUSD:8.99}}
	netflix.print()

	// Theres no need to deference the pointer to the struct because Go does it for us if the function is a pointer receiver
	netflixPointer.updatePrice(20.99)
	netflix.print()
}

// there is difference between a type of struct and a operation on a struct
func (pointerToService *service) updatePrice(newPrice float64) { // creates a copy of newPrice
	(*pointerToService).costPerMonth = newPrice
}

// a method for the service struct to print the struct
func (s service) print() {
	fmt.Printf("%+v\n", s)
}

// a method for the service struct to change the country of the subscription
func (s *service) countryChange(country string) {
	(*s).country = country
	(*s).priceInUSD = (*s).price * (*s).conversionRate
}