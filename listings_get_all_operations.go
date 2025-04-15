package stockxgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	GetAllListingOperationsEndpoint = "https://api.stockx.com/v2/selling/listings/%v/operations"
)

func (s *stockXClient) GetAllListingOperations(listingID string) (GetAllListingOperationsResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(GetAllListingOperationsEndpoint, listingID), nil)
	if err != nil {
		return GetAllListingOperationsResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.Session.AccessToken))
	req.Header.Set("x-api-key", s.clientID)

	resp, err := s.client.Do(req)
	if err != nil {
		return GetAllListingOperationsResponse{}, err
	}

	defer resp.Body.Close()

	if err := statusCode(resp.StatusCode); err != nil {
		return GetAllListingOperationsResponse{}, err
	}

	var response GetAllListingOperationsResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return GetAllListingOperationsResponse{}, err
	}

	return response, nil
}

type GetAllListingOperationsResponse struct {
	NextCursor string `json:"nextCursor"`
	Operations []struct {
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
	} `json:"operations"`
}
