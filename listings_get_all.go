package stockxgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	GetAllListingsEndpoint = "https://api.stockx.com/v2/selling/listings"
)

func (s *stockXClient) GetAllListings() (GetAllListingsResponse, error) {
	req, err := http.NewRequest("GET", GetAllListingsEndpoint, nil)
	if err != nil {
		return GetAllListingsResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.session.AccessToken))
	req.Header.Set("x-api-key", s.clientID)

	resp, err := s.client.Do(req)
	if err != nil {
		return GetAllListingsResponse{}, err
	}

	defer resp.Body.Close()

	if err := statusCode(resp.StatusCode); err != nil {
		return GetAllListingsResponse{}, err
	}

	var response GetAllListingsResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return GetAllListingsResponse{}, err
	}

	return response, nil
}

type GetAllListingsResponse struct {
	Count       int       `json:"count"`
	PageSize    int       `json:"pageSize"`
	PageNumber  int       `json:"pageNumber"`
	HasNextPage bool      `json:"hasNextPage"`
	Listings    []Listing `json:"listings"`
}

type Listing struct {
	ListingID     string    `json:"listingId"`
	Status        string    `json:"status"`
	Amount        string    `json:"amount"`
	CurrencyCode  string    `json:"currencyCode"`
	InventoryType string    `json:"inventoryType"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	Batch         struct {
		BatchID string `json:"batchId"`
		TaskID  string `json:"taskId"`
	} `json:"batch"`
	Ask struct {
		AskID        string    `json:"askId"`
		AskCreatedAt time.Time `json:"askCreatedAt"`
		AskUpdatedAt time.Time `json:"askUpdatedAt"`
		AskExpiresAt time.Time `json:"askExpiresAt"`
	} `json:"ask"`
	AuthenticationDetails struct {
		Status       string `json:"status"`
		FailureNotes string `json:"failureNotes"`
	} `json:"authenticationDetails"`
	Order struct {
		OrderNumber    string    `json:"orderNumber"`
		OrderCreatedAt time.Time `json:"orderCreatedAt"`
		OrderStatus    string    `json:"orderStatus"`
	} `json:"order"`
	Product struct {
		ProductID   string `json:"productId"`
		ProductName string `json:"productName"`
		StyleID     string `json:"styleId"`
	} `json:"product"`
	InitiatedShipments struct {
		Inbound struct {
			DisplayID string `json:"displayId"`
		} `json:"inbound"`
	} `json:"initiatedShipments"`
	Variant struct {
		VariantID    string `json:"variantId"`
		VariantName  string `json:"variantName"`
		VariantValue string `json:"variantValue"`
	} `json:"variant"`
}
