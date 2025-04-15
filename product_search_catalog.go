package stockxgo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type SearchCatalogRequest struct {
	Query      string
	PageNumber int
	PageSize   int
}

type SearchCatalogOption func(*SearchCatalogRequest)

func WithSearchCatalogPageNumber(pageNumber int) SearchCatalogOption {
	return func(r *SearchCatalogRequest) {
		r.PageNumber = pageNumber
	}
}

func WithSearchCatalogPageSize(pageSize int) SearchCatalogOption {
	return func(r *SearchCatalogRequest) {
		r.PageSize = pageSize
	}
}

func WithSearchCatalogQuery(query string) SearchCatalogOption {
	return func(r *SearchCatalogRequest) {
		r.Query = query
	}
}

func (s *stockXClient) SearchCatalog(opts ...SearchCatalogOption) (SearchCatalogResponse, error) {
	request := &SearchCatalogRequest{
		PageNumber: 1,
		PageSize:   10,
	}

	for _, opt := range opts {
		opt(request)
	}

	queryParams := url.Values{}
	queryParams.Add("query", request.Query)
	queryParams.Add("pageNumber", fmt.Sprintf("%d", request.PageNumber))
	queryParams.Add("pageSize", fmt.Sprintf("%d", request.PageSize))

	url := fmt.Sprintf("https://api.stockx.com/v2/catalog/products/search?%s", queryParams.Encode())

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return SearchCatalogResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.session.AccessToken))
	req.Header.Set("x-api-key", s.clientID)

	resp, err := s.client.Do(req)
	if err != nil {
		return SearchCatalogResponse{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return SearchCatalogResponse{}, err
	}

	var response SearchCatalogResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return SearchCatalogResponse{}, err
	}

	return response, nil
}

type SearchCatalogResponse struct {
	Count       int       `json:"count"`
	PageSize    int       `json:"pageSize"`
	PageNumber  int       `json:"pageNumber"`
	HasNextPage bool      `json:"hasNextPage"`
	Products    []Product `json:"products"`
}
