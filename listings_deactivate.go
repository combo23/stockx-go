package stockxgo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	DeactivateListingEndpoint = "https://api.stockx.com/v2/selling/listings/%v/deactivate"
)

func (s *stockXClient) DeactivateListing(listingID string) (ListingModificationResponse, error) {
	req, err := http.NewRequest("PUT", fmt.Sprintf(DeactivateListingEndpoint, listingID), nil)
	if err != nil {
		return ListingModificationResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.Session.AccessToken))
	req.Header.Set("x-api-key", s.clientID)

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
