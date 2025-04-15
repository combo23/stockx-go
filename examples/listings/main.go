package main

import (
	"fmt"
	"log"

	stockxgo "github.com/combo23/stockx-go"
)

func main() {
	client := stockxgo.NewClient("code", "client_id", "client_secret", "")

	err := client.Authenticate()
	if err != nil {
		log.Fatal(err)
	}

	// get a single listing by its id
	listing, err := client.GetListing("1234567890")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(listing)

	// get all listings
	listings, err := client.GetAllListings()
	if err != nil {
		log.Fatal(err)
	}

	for _, listing := range listings.Listings {
		fmt.Println(listing)
	}
}
