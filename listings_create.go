package stockxgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	CreateListingEndpoint = "https://api.stockx.com/v2/selling/listings"
)

type CreateLisingPayload struct {
	Amount       string    `json:"amount"`
	VariantID    string    `json:"variantId"`
	CurrencyCode string    `json:"currencyCode"`
	ExpiresAt    time.Time `json:"expiresAt"`
	Active       bool      `json:"active"`
}

type CreateListingOption func(*CreateLisingPayload)

func NewCreateListingPayload(amount, variantID string, opts ...CreateListingOption) CreateLisingPayload {
	payload := CreateLisingPayload{
		Amount:    amount,
		VariantID: variantID,
	}

	for _, opt := range opts {
		opt(&payload)
	}

	return payload
}

func WithCurrencyCode(currencyCode string) CreateListingOption {
	return func(payload *CreateLisingPayload) {
		payload.CurrencyCode = currencyCode
	}
}

func WithExpiresAt(expiresAt time.Time) CreateListingOption {
	return func(payload *CreateLisingPayload) {
		payload.ExpiresAt = expiresAt
	}
}

func WithActive(active bool) CreateListingOption {
	return func(payload *CreateLisingPayload) {
		payload.Active = active
	}
}

func (s *stockXClient) CreateListing(payload CreateLisingPayload) (ListingModificationResponse, error) {
	payloadRaw, err := json.Marshal(payload)
	if err != nil {
		return ListingModificationResponse{}, err
	}

	req, err := http.NewRequest("POST", CreateListingEndpoint, bytes.NewBuffer(payloadRaw))
	if err != nil {
		return ListingModificationResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.session.AccessToken))
	req.Header.Set("x-api-key", s.apiKey)

	resp, err := s.client.Do(req)
	if err != nil {
		return ListingModificationResponse{}, err
	}

	defer resp.Body.Close()

	if err := statusCode(resp.StatusCode); err != nil {
		return ListingModificationResponse{}, err
	}

	var response ListingModificationResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return ListingModificationResponse{}, err
	}

	return response, nil
}

type ListingModificationResponse struct {
	ListingID             string    `json:"listingId"`
	OperationID           string    `json:"operationId"`
	OperationType         string    `json:"operationType"`
	OperationStatus       string    `json:"operationStatus"`
	OperationURL          string    `json:"operationUrl"`
	OperationInitiatedBy  string    `json:"operationInitiatedBy"`
	OperationInitiatedVia string    `json:"operationInitiatedVia"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
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
	Error interface{} `json:"error"`
}
