package stockxgo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	ProductVariantGetSingleEndpoint = "https://api.stockx.com/v2/catalog/products/%v/variants/%v"
)

func (s *stockXClient) GetSingleProductVariant(productID, variantID string) (ProductVariant, error) {
	url := fmt.Sprintf(ProductVariantGetSingleEndpoint, productID, variantID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ProductVariant{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.session.AccessToken))
	req.Header.Set("x-api-key", s.clientID)

	resp, err := s.client.Do(req)
	if err != nil {
		return ProductVariant{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ProductVariant{}, err
	}

	var productVariant ProductVariant
	err = json.Unmarshal(body, &productVariant)
	if err != nil {
		return ProductVariant{}, err
	}

	return productVariant, nil
}

type ProductVariant struct {
	ProductID    string `json:"productId"`
	VariantID    string `json:"variantId"`
	VariantName  string `json:"variantName"`
	VariantValue string `json:"variantValue"`
	SizeChart    struct {
		AvailableConversions []struct {
			Size string `json:"size"`
			Type string `json:"type"`
		} `json:"availableConversions"`
		DefaultConversion struct {
			Size string `json:"size"`
			Type string `json:"type"`
		} `json:"defaultConversion"`
	} `json:"sizeChart"`
}
