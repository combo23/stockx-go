package main

import (
	"fmt"
	"log"

	stockxgo "github.com/combo23/stockx-go"
)

func main() {
	client := stockxgo.NewClient("code", "client_id", "client_secret", "api_key")

	err := client.Authenticate()
	if err != nil {
		log.Fatal(err)
	}

	// get a single order by its order number
	order, err := client.GetOrder("1234567890")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(order)

	// get all active orders with specific filters
	orders, err := client.GetActiveOrders(
		stockxgo.WithActivePageNumber(1),
		stockxgo.WithActivePageSize(10),
		stockxgo.WithActiveOrderStatus(stockxgo.OrderStatusAuthenticated),
	)

	if err != nil {
		log.Fatal(err)
	}

	for _, order := range orders.Orders {
		fmt.Println(order)
	}
}
