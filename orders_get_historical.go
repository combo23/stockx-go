package stockxgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	OrdersGetHistoricalEndpoint = "https://api.stockx.com/v2/selling/orders/history"
)

// HistoricalOrdersRequest holds the parameters for the active orders request
type HistoricalOrdersRequest struct {
	FromDate                    string
	ToDate                      string
	PageNumber                  int
	PageSize                    int
	OrderStatus                 OrderStatus
	ProductID                   string
	VariantID                   string
	InventoryType               []InventoryType
	InitiatedShipmentDisplayIds []string
}

type HistoricalOrdersOption func(*HistoricalOrdersRequest)

func WithHistoricalFromDate(fromDate string) HistoricalOrdersOption {
	return func(r *HistoricalOrdersRequest) {
		r.FromDate = fromDate
	}
}

func WithHistoricalToDate(toDate string) HistoricalOrdersOption {
	return func(r *HistoricalOrdersRequest) {
		r.ToDate = toDate
	}
}

func WithHistoricalPageNumber(pageNumber int) HistoricalOrdersOption {
	return func(r *HistoricalOrdersRequest) {
		r.PageNumber = pageNumber
	}
}

func WithHistoricalPageSize(pageSize int) HistoricalOrdersOption {
	return func(r *HistoricalOrdersRequest) {
		r.PageSize = pageSize
	}
}

func WithHistoricalOrderStatus(orderStatus OrderStatus) HistoricalOrdersOption {
	return func(r *HistoricalOrdersRequest) {
		r.OrderStatus = orderStatus
	}
}

func WithHistoricalProductID(productID string) HistoricalOrdersOption {
	return func(r *HistoricalOrdersRequest) {
		r.ProductID = productID
	}
}

func WithHistoricalVariantID(variantID string) HistoricalOrdersOption {
	return func(r *HistoricalOrdersRequest) {
		r.VariantID = variantID
	}
}

func WithHistoricalInventoryType(inventoryType InventoryType) HistoricalOrdersOption {
	return func(r *HistoricalOrdersRequest) {
		r.InventoryType = append(r.InventoryType, inventoryType)
	}
}

func WithHistoricalInitiatedShipmentDisplayIds(initiatedShipmentDisplayIds ...string) HistoricalOrdersOption {
	return func(r *HistoricalOrdersRequest) {
		r.InitiatedShipmentDisplayIds = initiatedShipmentDisplayIds
	}
}

func (s *stockXClient) GetHistoricalOrders(opts ...HistoricalOrdersOption) (OrdersResponse, error) {
	req := &HistoricalOrdersRequest{
		PageNumber: 1,
		PageSize:   20,
	}

	for _, opt := range opts {
		opt(req)
	}

	u, err := url.Parse(OrdersGetHistoricalEndpoint)
	if err != nil {
		return OrdersResponse{}, err
	}

	q := u.Query()
	q.Add("pageNumber", strconv.Itoa(req.PageNumber))
	q.Add("pageSize", strconv.Itoa(req.PageSize))

	if req.FromDate != "" {
		q.Add("fromDate", req.FromDate)
	}

	if req.ToDate != "" {
		q.Add("toDate", req.ToDate)
	}

	if req.OrderStatus != "" {
		q.Add("orderStatus", string(req.OrderStatus))
	}

	if req.ProductID != "" {
		q.Add("productId", req.ProductID)
	}

	if req.VariantID != "" {
		q.Add("variantId", req.VariantID)
	}

	if len(req.InventoryType) > 0 {
		inventoryTypesStr := make([]string, len(req.InventoryType))
		for i, invType := range req.InventoryType {
			inventoryTypesStr[i] = string(invType)
		}
		q.Add("inventoryTypes", strings.Join(inventoryTypesStr, ","))
	}

	if len(req.InitiatedShipmentDisplayIds) > 0 {
		q.Add("initiatedShipmentDisplayIds", strings.Join(req.InitiatedShipmentDisplayIds, ","))
	}

	u.RawQuery = q.Encode()

	httpReq, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return OrdersResponse{}, err
	}

	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.Session.AccessToken))
	httpReq.Header.Set("x-api-key", s.clientID)

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
