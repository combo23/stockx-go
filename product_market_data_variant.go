package stockxgo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	ProductMarketDataVariantEndpoint = "https://api.stockx.com/v2/catalog/products/%v/variants/%v/market-data?currencyCode=%v"
)

func (s *stockXClient) GetProductMarketDataForVariant(productID, variantID, currencyCode string) (MarketData, error) {
	url := fmt.Sprintf(ProductMarketDataVariantEndpoint, productID, variantID, currencyCode)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return MarketData{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.session.AccessToken))
	req.Header.Set("x-api-key", s.apiKey)

	resp, err := s.client.Do(req)
	if err != nil {
		return MarketData{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return MarketData{}, err
	}

	var productMarketDataVariant MarketData
	err = json.Unmarshal(body, &productMarketDataVariant)
	if err != nil {
		return MarketData{}, err
	}

	return productMarketDataVariant, nil
}

type MarketData struct {
	ProductID           string `json:"productId"`
	VariantID           string `json:"variantId"`
	CurrencyCode        string `json:"currencyCode"`
	LowestAskAmount     string `json:"lowestAskAmount"`
	HighestBidAmount    string `json:"highestBidAmount"`
	SellFasterAmount    string `json:"sellFasterAmount"`
	EarnMoreAmount      string `json:"earnMoreAmount"`
	FlexLowestAskAmount string `json:"flexLowestAskAmount"`
}
