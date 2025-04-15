package stockxgo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	ProductVariantGetAllEndpoint = "https://api.stockx.com/v2/catalog/products/%v/variants"
)

func (s *stockXClient) GetAllProductVariants(productID string) ([]ProductVariant, error) {
	url := fmt.Sprintf(ProductVariantGetAllEndpoint, productID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.session.AccessToken))
	req.Header.Set("x-api-key", s.apiKey)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if err := statusCode(resp.StatusCode); err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var productVariants []ProductVariant
	err = json.Unmarshal(body, &productVariants)
	if err != nil {
		return nil, err
	}

	return productVariants, nil
}
