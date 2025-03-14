package stockxgo

import "net/http"

type StockXClient interface {
	GetOrder(orderNumber string) (GetSingleOrderResponse, error)
	GetActiveOrders(options ...ActiveOrdersOption) (OrdersResponse, error)
	GetHistoricalOrders(options ...HistoricalOrdersOption) (OrdersResponse, error)
	Authenticate() error
	CreateListing(payload CreateLisingPayload) (ListingModificationResponse, error)
	GetAllListings() (GetAllListingsResponse, error)
	GetListing(listingID string) (GetListingResponse, error)
	GetAllListingOperations(listingID string) (GetAllListingOperationsResponse, error)
	GetListingOperation(listingID, operationID string) (GetListingOperationResponse, error)
	ActivateListing(listingID string, payload ActivateListingPayload) (ListingModificationResponse, error)
	DeactivateListing(listingID string) (ListingModificationResponse, error)
	UpdateListing(listingID string, payload UpdateListingPayload) (ListingModificationResponse, error)
	DeleteListing(listingID string) (ListingModificationResponse, error)
}

type stockXClient struct {
	client       *http.Client
	code         string
	clientID     string
	clientSecret string
	session      Session
}

type Session struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int
}

func NewClient(code, clientID, clientSecret string) StockXClient {
	return &stockXClient{
		code:         code,
		clientID:     clientID,
		clientSecret: clientSecret,
		client:       &http.Client{},
	}
}
