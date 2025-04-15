package stockxgo

import "net/http"

type StockXClient interface {
	GetOrder(orderNumber string) (GetSingleOrderResponse, error)
	GetActiveOrders(options ...ActiveOrdersOption) (OrdersResponse, error)
	GetHistoricalOrders(options ...HistoricalOrdersOption) (OrdersResponse, error)
	Authenticate() error
	RefreshToken() error
	CreateListing(payload CreateLisingPayload) (ListingModificationResponse, error)
	GetAllListings(options ...GetAllListingsOption) (GetAllListingsResponse, error)
	GetListing(listingID string) (GetListingResponse, error)
	GetAllListingOperations(listingID string) (GetAllListingOperationsResponse, error)
	GetListingOperation(listingID, operationID string) (GetListingOperationResponse, error)
	ActivateListing(listingID string, payload ActivateListingPayload) (ListingModificationResponse, error)
	DeactivateListing(listingID string) (ListingModificationResponse, error)
	UpdateListing(listingID string, payload UpdateListingPayload) (ListingModificationResponse, error)
	DeleteListing(listingID string) (ListingModificationResponse, error)
	SearchCatalog(opts ...SearchCatalogOption) (SearchCatalogResponse, error)
	GetSingleProduct(productID string) (Product, error)
	GetAllProductVariants(productID string) ([]ProductVariant, error)
	GetSingleProductVariant(productID, variantID string) (ProductVariant, error)
	GetProductMarketData(productID, currencyCode string) ([]MarketData, error)
	GetProductMarketDataForVariant(productID, variantID, currencyCode string) (MarketData, error)
	GetAccessToken() string
	GetRefreshToken() string
	GetExpiresIn() int
}

type stockXClient struct {
	client       *http.Client
	code         string
	clientID     string
	clientSecret string
	session      Session
	apiKey       string
}

type Session struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int
}

func NewClient(code, clientID, clientSecret, apiKey string) StockXClient {
	return &stockXClient{
		code:         code,
		clientID:     clientID,
		clientSecret: clientSecret,
		apiKey:       apiKey,
		client:       &http.Client{},
	}
}

func NewClientWithSession(session Session, clientID, clientSecret, apiKey string) StockXClient {
	return &stockXClient{
		session:      session,
		clientID:     clientID,
		clientSecret: clientSecret,
		apiKey:       apiKey,
		client:       &http.Client{},
	}
}
