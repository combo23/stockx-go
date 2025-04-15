package stockxgo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	DeleteListingEndpoint = "https://api.stockx.com/v2/selling/listings/%v"
)

func (s *stockXClient) DeleteListing(listingID string) (ListingModificationResponse, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf(DeleteListingEndpoint, listingID), nil)
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
