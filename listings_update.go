package stockxgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	UpdateListingEndpoint = "https://api.stockx.com/v2/selling/listings/%v"
)

type UpdateListingPayload struct {
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currencyCode"`
	ExpiresAt    string `json:"expiresAt"`
}

func NewUpdateListingPayload(amount, currencyCode, expiresAt string) UpdateListingPayload {
	return UpdateListingPayload{
		Amount:       amount,
		CurrencyCode: currencyCode,
		ExpiresAt:    expiresAt,
	}
}

func (s *stockXClient) UpdateListing(listingID string, payload UpdateListingPayload) (ListingModificationResponse, error) {
	payloadRaw, err := json.Marshal(payload)
	if err != nil {
		return ListingModificationResponse{}, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf(UpdateListingEndpoint, listingID), bytes.NewBuffer(payloadRaw))
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
