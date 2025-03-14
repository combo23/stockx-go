# stockx-go

An unofficial Go wrapper for the StockX API that allows you to interact with StockX's services programmatically.

## Installation

```bash
go get github.com/combo23/stockx-go
```

## Features

- Manage Listings
- Manage Orders

## Quick Start

```go
package main

import (
    "fmt"
    
    stockxgo "github.com/combo23/stockx-go"
)

func main() {
    client := stockx.NewClient("xxxxx", "xxxxx", "xxxxx")

    err := client.Authenticate()
    if err != nil {
        panic(err)
    }
    
    // Get all user's listings
    listings, err := client.GetAllListings()
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Found %d listings\n", len(listings.Listings))
}
```

## TODO

- Implement Catalog API Endpoints
- Improve Documentation

## API Reference

[Click Here]("https://developer.stockx.com/openapi/reference/overview/")

### Client Methods

-`GetOrder(orderNumber string) (GetSingleOrderResponse, error)`

-`GetActiveOrders(options ...ActiveOrdersOption) (OrdersResponse, error)`

-`GetHistoricalOrders(options ...HistoricalOrdersOption) (OrdersResponse, error)`

-`CreateListing(payload CreateLisingPayload) (ListingModificationResponse, error)`

-`GetAllListings() (GetAllListingsResponse, error)`

-`GetListing(listingID string) (GetListingResponse, error)`

-`GetAllListingOperations(listingID string) (GetAllListingOperationsResponse, error)`

-`GetListingOperation(listingID, operationID string) (GetListingOperationResponse, error)`

-`ActivateListing(listingID string, payload ActivateListingPayload) (ListingModificationResponse, error)`

-`DeactivateListing(listingID string) (ListingModificationResponse, error)`

-`UpdateListing(listingID string, payload UpdateListingPayload) (ListingModificationResponse, error)`

-`DeleteListing(listingID string) (ListingModificationResponse, error)`

## Error Handling

The library returns detailed errors that can be handled using standard Go error handling:

```go
if err != nil {
    switch e := err.(type) {
    case *stockx.ErrUnauthorized:
        // Handle authentication errors
    case *stockx.ErrInternal:
        // Handle API errors
    default:
        // Handle other errors
    }
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Disclaimer

This is an unofficial wrapper for the StockX API. This project is not affiliated with, maintained, authorized, endorsed, or sponsored by StockX.