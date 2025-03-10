package stockxgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	GetOrdersEndpoint = "https://api.stockx.com/v2/selling/orders/%v"
)

func (s *stockXClient) GetOrder(orderNumber string) (GetSingleOrderResponse, error) {
	url := fmt.Sprintf(GetOrdersEndpoint, orderNumber)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return GetSingleOrderResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.session.AccessToken))
	req.Header.Set("x-api-key", s.clientID)

	resp, err := s.client.Do(req)
	if err != nil {
		return GetSingleOrderResponse{}, err
	}

	defer resp.Body.Close()

	if err := statusCode(resp.StatusCode); err != nil {
		return GetSingleOrderResponse{}, err
	}

	var response GetSingleOrderResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return GetSingleOrderResponse{}, err
	}

	return response, nil
}

type GetSingleOrderResponse struct {
	AskID        string    `json:"askId"`
	OrderNumber  string    `json:"orderNumber"`
	ListingID    string    `json:"listingId"`
	Amount       string    `json:"amount"`
	CurrencyCode string    `json:"currencyCode"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Variant      struct {
		VariantID    string `json:"variantId"`
		VariantName  string `json:"variantName"`
		VariantValue string `json:"variantValue"`
	} `json:"variant"`
	Product struct {
		ProductID   string `json:"productId"`
		ProductName string `json:"productName"`
		StyleID     string `json:"styleId"`
	} `json:"product"`
	Status   string `json:"status"`
	Shipment struct {
		ShipByDate          string `json:"shipByDate"`
		TrackingNumber      string `json:"trackingNumber"`
		TrackingURL         string `json:"trackingUrl"`
		CarrierCode         string `json:"carrierCode"`
		ShippingLabelURL    string `json:"shippingLabelUrl"`
		ShippingDocumentURL string `json:"shippingDocumentUrl"`
	} `json:"shipment"`
	InitiatedShipments struct {
		Inbound struct {
			DisplayID string `json:"displayId"`
		} `json:"inbound"`
	} `json:"initiatedShipments"`
	InventoryType         string `json:"inventoryType"`
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
}
