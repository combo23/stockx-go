package stockxgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	GetListingsEndpoint = "https://api.stockx.com/v2/selling/listings/%v"
)

func (s *stockXClient) GetListing(listingID string) (GetListingResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(GetListingsEndpoint, listingID), nil)
	if err != nil {
		return GetListingResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.Session.AccessToken))
	req.Header.Set("x-api-key", s.clientID)

	resp, err := s.client.Do(req)
	if err != nil {
		return GetListingResponse{}, err
	}

	defer resp.Body.Close()

	if err := statusCode(resp.StatusCode); err != nil {
		return GetListingResponse{}, err
	}

	var response GetListingResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return GetListingResponse{}, err
	}

	return response, nil
}

type GetListingResponse struct {
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
	Variant struct {
		VariantID    string `json:"variantId"`
		VariantName  string `json:"variantName"`
		VariantValue string `json:"variantValue"`
	} `json:"variant"`
	AuthenticationDetails struct {
		Status       string `json:"status"`
		FailureNotes string `json:"failureNotes"`
	} `json:"authenticationDetails"`
	Payout struct {
		TotalPayout      float64 `json:"totalPayout"`
		SalePrice        int     `json:"salePrice"`
		TotalAdjustments int     `json:"totalAdjustments"`
		CurrencyCode     string  `json:"currencyCode"`
		Adjustments      []struct {
			AdjustmentType string  `json:"adjustmentType"`
			Amount         float64 `json:"amount"`
			Percentage     float64 `json:"percentage"`
		} `json:"adjustments"`
	} `json:"payout"`
	LastOperation struct {
		OperationID           string    `json:"operationId"`
		OperationType         string    `json:"operationType"`
		OperationStatus       string    `json:"operationStatus"`
		OperationInitiatedBy  string    `json:"operationInitiatedBy"`
		OperationInitiatedVia string    `json:"operationInitiatedVia"`
		OperationCreatedAt    time.Time `json:"operationCreatedAt"`
		OperationUpdatedAt    time.Time `json:"operationUpdatedAt"`
		Changes               struct {
			Additions struct {
				Active  bool `json:"active"`
				AskData struct {
					Amount    string    `json:"amount"`
					Currency  string    `json:"currency"`
					ExpiresAt time.Time `json:"expiresAt"`
				} `json:"askData"`
			} `json:"additions"`
			Updates struct {
				UpdatedAt time.Time `json:"updatedAt"`
			} `json:"updates"`
			Removals struct {
			} `json:"removals"`
		} `json:"changes"`
		Error string `json:"error"`
	} `json:"lastOperation"`
	InitiatedShipments struct {
		Inbound struct {
			DisplayID string `json:"displayId"`
		} `json:"inbound"`
	} `json:"initiatedShipments"`
}
