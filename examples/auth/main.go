package main

import (
	"log"

	stockxgo "github.com/combo23/stockx-go"
)

func main() {
	// code - value from the redirect uri
	// client_id - your client id
	// client_secret - your client secret
	// reference: https://developer.stockx.com/portal/authentication/
	client := stockxgo.NewClient("code", "client_id", "client_secret")

	err := client.Authenticate()
	if err != nil {
		log.Fatal(err)
	}
}
