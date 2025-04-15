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

const (
	OrdersGetActiveEndpoint = "https://api.stockx.com/v2/selling/orders/active"
)

// OrderStatus represents possible order statuses
type OrderStatus string

const (
	OrderStatusCreated         OrderStatus = "CREATED"
	OrderStatusCCAuthFailed    OrderStatus = "CCAUTHORIZATIONFAILED"
	OrderStatusShipped         OrderStatus = "SHIPPED"
	OrderStatusReceived        OrderStatus = "RECEIVED"
	OrderStatusAuthenticating  OrderStatus = "AUTHENTICATING"
	OrderStatusAuthenticated   OrderStatus = "AUTHENTICATED"
	OrderStatusPayoutPending   OrderStatus = "PAYOUTPENDING"
	OrderStatusPayoutCompleted OrderStatus = "PAYOUTCOMPLETED"
	OrderStatusSystemFulfilled OrderStatus = "SYSTEMFULFILLED"
	OrderStatusPayoutFailed    OrderStatus = "PAYOUTFAILED"
	OrderStatusSuspended       OrderStatus = "SUSPENDED"
)

// SortField represents possible sorting fields
type SortField string

const (
	SortFieldCreatedAt  SortField = "CREATEDAT"
	SortFieldShipByDate SortField = "SHIPBYDATE"
)

// InventoryType represents possible inventory types
type InventoryType string

const (
	InventoryTypeStandard InventoryType = "STANDARD"
	InventoryTypeFlex     InventoryType = "FLEX"
)

type ActiveOrdersOption func(*ActiveOrdersRequest)

// ActiveOrdersRequest holds the parameters for the active orders request
type ActiveOrdersRequest struct {
	PageNumber                  int
	PageSize                    int
	OrderStatus                 OrderStatus
	ProductID                   string
	VariantID                   string
	SortField                   SortField
	InventoryTypes              []InventoryType
	InitiatedShipmentDisplayIDs []string
}

// WithPageNumber sets the page number for pagination
// Must be >= 1
func WithActivePageNumber(pageNumber int) ActiveOrdersOption {
	return func(r *ActiveOrdersRequest) {
		if pageNumber < 1 {
			pageNumber = 1
		}
		r.PageNumber = pageNumber
	}
}

// WithPageSize sets the page size for pagination
// Must be between 1 and 100
func WithActivePageSize(pageSize int) ActiveOrdersOption {
	return func(r *ActiveOrdersRequest) {
		if pageSize < 1 {
			pageSize = 1
		} else if pageSize > 100 {
			pageSize = 100
		}
		r.PageSize = pageSize
	}
}

// WithOrderStatus filters orders by status
func WithActiveOrderStatus(status OrderStatus) ActiveOrdersOption {
	return func(r *ActiveOrdersRequest) {
		r.OrderStatus = status
	}
}

// WithProductID filters orders by product ID
func WithActiveProductID(productID string) ActiveOrdersOption {
	return func(r *ActiveOrdersRequest) {
		r.ProductID = productID
	}
}

// WithVariantID filters orders by variant ID
func WithActiveVariantID(variantID string) ActiveOrdersOption {
	return func(r *ActiveOrdersRequest) {
		r.VariantID = variantID
	}
}

// WithSortField sets the sort field of the results
// Defaults to CREATEDAT if not provided
func WithActiveSortField(sortField SortField) ActiveOrdersOption {
	return func(r *ActiveOrdersRequest) {
		r.SortField = sortField
	}
}

// WithInventoryTypes filters orders by inventory types
// Valid values are STANDARD and FLEX
func WithActiveInventoryTypes(inventoryTypes ...InventoryType) ActiveOrdersOption {
	return func(r *ActiveOrdersRequest) {
		r.InventoryTypes = inventoryTypes
	}
}

// WithInitiatedShipmentDisplayIDs filters orders by shipment display IDs
func WithActiveInitiatedShipmentDisplayIDs(ids ...string) ActiveOrdersOption {
	return func(r *ActiveOrdersRequest) {
		r.InitiatedShipmentDisplayIDs = ids
	}
}

func (s *stockXClient) GetActiveOrders(opts ...ActiveOrdersOption) (OrdersResponse, error) {
	req := &ActiveOrdersRequest{
		PageNumber: 1,
		PageSize:   20,
		SortField:  SortFieldCreatedAt,
	}

	for _, opt := range opts {
		opt(req)
	}

	u, err := url.Parse(OrdersGetActiveEndpoint)
	if err != nil {
		return OrdersResponse{}, err
	}

	q := u.Query()
	q.Add("pageNumber", strconv.Itoa(req.PageNumber))
	q.Add("pageSize", strconv.Itoa(req.PageSize))

	if req.OrderStatus != "" {
		q.Add("orderStatus", string(req.OrderStatus))
	}

	if req.ProductID != "" {
		q.Add("productId", req.ProductID)
	}

	if req.VariantID != "" {
		q.Add("variantId", req.VariantID)
	}

	if req.SortField != "" {
		q.Add("sortOrder", string(req.SortField))
	}

	if len(req.InventoryTypes) > 0 {
		inventoryTypesStr := make([]string, len(req.InventoryTypes))
		for i, invType := range req.InventoryTypes {
			inventoryTypesStr[i] = string(invType)
		}
		q.Add("inventoryTypes", strings.Join(inventoryTypesStr, ","))
	}

	if len(req.InitiatedShipmentDisplayIDs) > 0 {
		q.Add("initiatedShipmentDisplayIds", strings.Join(req.InitiatedShipmentDisplayIDs, ","))
	}

	u.RawQuery = q.Encode()

	httpReq, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return OrdersResponse{}, err
	}

	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.session.AccessToken))
	httpReq.Header.Set("x-api-key", s.apiKey)

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return OrdersResponse{}, err
	}

	defer resp.Body.Close()

	if err := statusCode(resp.StatusCode); err != nil {
		return OrdersResponse{}, err
	}

	var response OrdersResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return OrdersResponse{}, err
	}

	return response, nil
}

type OrdersResponse struct {
	Count       int     `json:"count"`
	PageSize    int     `json:"pageSize"`
	PageNumber  int     `json:"pageNumber"`
	HasNextPage bool    `json:"hasNextPage"`
	Orders      []Order `json:"orders"`
}

type Order struct {
	OrderNumber  string    `json:"orderNumber"`
	ListingID    string    `json:"listingId"`
	AskID        string    `json:"askId"`
	Amount       string    `json:"amount"`
	CurrencyCode string    `json:"currencyCode"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Product      struct {
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
	InitiatedShipments struct {
		Inbound struct {
			DisplayID string `json:"displayId"`
		} `json:"inbound"`
	} `json:"initiatedShipments"`
	InventoryType string `json:"inventoryType"`
}
