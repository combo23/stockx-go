package stockxgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var (
	GetAllListingsEndpoint = "https://api.stockx.com/v2/selling/listings"
)

type GetAllListingsRequest struct {
	pageNumber                  int
	pageSize                    int
	productIDs                  []string
	variantIDs                  []string
	batchIDs                    []string
	fromDate                    time.Time
	toDate                      time.Time
	listingStatuses             []string
	inventoryTypes              []string
	initiatedShipmentDisplayIds []string
}

func WithGetAllListingsPageNumber(pageNumber int) GetAllListingsOption {
	return func(r *GetAllListingsRequest) {
		r.pageNumber = pageNumber
	}
}

func WithGetAllListingsPageSize(pageSize int) GetAllListingsOption {
	return func(r *GetAllListingsRequest) {
		r.pageSize = pageSize
	}
}

func WithGetAllListingsProductIDs(productIDs []string) GetAllListingsOption {
	return func(r *GetAllListingsRequest) {
		r.productIDs = productIDs
	}
}

func WithGetAllListingsVariantIDs(variantIDs []string) GetAllListingsOption {
	return func(r *GetAllListingsRequest) {
		r.variantIDs = variantIDs
	}
}

func WithGetAllListingsBatchIDs(batchIDs []string) GetAllListingsOption {
	return func(r *GetAllListingsRequest) {
		r.batchIDs = batchIDs
	}
}

func WithGetAllListingsFromDate(fromDate time.Time) GetAllListingsOption {
	return func(r *GetAllListingsRequest) {
		r.fromDate = fromDate
	}
}

func WithGetAllListingsToDate(toDate time.Time) GetAllListingsOption {
	return func(r *GetAllListingsRequest) {
		r.toDate = toDate
	}
}

func WithGetAllListingsListingStatuses(listingStatuses []string) GetAllListingsOption {
	return func(r *GetAllListingsRequest) {
		r.listingStatuses = listingStatuses
	}
}

func WithGetAllListingsInventoryTypes(inventoryTypes []string) GetAllListingsOption {
	return func(r *GetAllListingsRequest) {
		r.inventoryTypes = inventoryTypes
	}
}

func WithGetAllListingsInitiatedShipmentDisplayIds(initiatedShipmentDisplayIds []string) GetAllListingsOption {
	return func(r *GetAllListingsRequest) {
		r.initiatedShipmentDisplayIds = initiatedShipmentDisplayIds
	}
}

type GetAllListingsOption func(*GetAllListingsRequest)

func (s *stockXClient) GetAllListings(options ...GetAllListingsOption) (GetAllListingsResponse, error) {
	request := &GetAllListingsRequest{
		pageNumber: 1,
		pageSize:   100,
	}

	for _, option := range options {
		option(request)
	}

	queryParams := url.Values{}
	queryParams.Add("pageNumber", strconv.Itoa(request.pageNumber))
	queryParams.Add("pageSize", strconv.Itoa(request.pageSize))

	if len(request.productIDs) > 0 {
		queryParams.Add("productIds", strings.Join(request.productIDs, ","))
	}

	if len(request.variantIDs) > 0 {
		queryParams.Add("variantIds", strings.Join(request.variantIDs, ","))
	}

	if len(request.batchIDs) > 0 {
		queryParams.Add("batchIds", strings.Join(request.batchIDs, ","))
	}

	if !request.fromDate.IsZero() {
		queryParams.Add("fromDate", request.fromDate.Format(time.RFC3339))
	}

	if !request.toDate.IsZero() {
		queryParams.Add("toDate", request.toDate.Format(time.RFC3339))
	}

	if len(request.listingStatuses) > 0 {
		queryParams.Add("listingStatuses", strings.Join(request.listingStatuses, ","))
	}

	if len(request.inventoryTypes) > 0 {
		queryParams.Add("inventoryTypes", strings.Join(request.inventoryTypes, ","))
	}

	if len(request.initiatedShipmentDisplayIds) > 0 {
		queryParams.Add("initiatedShipmentDisplayIds", strings.Join(request.initiatedShipmentDisplayIds, ","))
	}

	url := fmt.Sprintf("%s?%s", GetAllListingsEndpoint, queryParams.Encode())

	req, err := http.NewRequest("GET", url, nil)
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
