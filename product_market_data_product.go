package stockxgo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	ProductMarketDataProductEndpoint = "https://api.stockx.com/v2/catalog/products/%v/market-data?currencyCode=%v"
)

func (s *stockXClient) GetProductMarketData(productID, currencyCode string) ([]MarketData, error) {
	url := fmt.Sprintf(ProductMarketDataProductEndpoint, productID, currencyCode)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []MarketData{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.Session.AccessToken))
	req.Header.Set("x-api-key", s.clientID)

	resp, err := s.client.Do(req)
	if err != nil {
		return []MarketData{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []MarketData{}, err
	}

	var productMarketData []MarketData
	err = json.Unmarshal(body, &productMarketData)
	if err != nil {
		return []MarketData{}, err
	}

	return productMarketData, nil
}
