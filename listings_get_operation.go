package stockxgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	GetListingOperation = "https://api.stockx.com/v2/selling/listings/%v/operations/%v"
)

func (s *stockXClient) GetListingOperation(listingID, operationID string) (GetListingOperationResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(GetListingOperation, listingID, operationID), nil)
	if err != nil {
		return GetListingOperationResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.Session.AccessToken))
	req.Header.Set("x-api-key", s.clientID)

	resp, err := s.client.Do(req)
	if err != nil {
		return GetListingOperationResponse{}, err
	}

	defer resp.Body.Close()

	if err := statusCode(resp.StatusCode); err != nil {
		return GetListingOperationResponse{}, err
	}

	var response GetListingOperationResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return GetListingOperationResponse{}, err
	}

	return response, nil
}

type GetListingOperationResponse struct {
	ListingID             string    `json:"listingId"`
	OperationID           string    `json:"operationId"`
	OperationType         string    `json:"operationType"`
	OperationStatus       string    `json:"operationStatus"`
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
