package stockxgo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	ProductGetSingleEndpoint = "https://api.stockx.com/v2/catalog/products/%v"
)

func (s *stockXClient) GetSingleProduct(productID string) (Product, error) {
	url := fmt.Sprintf(ProductGetSingleEndpoint, productID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Product{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.session.AccessToken))
	req.Header.Set("x-api-key", s.apiKey)

	resp, err := s.client.Do(req)
	if err != nil {
		return Product{}, err
	}

	defer resp.Body.Close()

	if err := statusCode(resp.StatusCode); err != nil {
		return Product{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Product{}, err
	}

	var product Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

type Product struct {
	ProductID         string `json:"productId"`
	URLKey            string `json:"urlKey"`
	StyleID           string `json:"styleId"`
	ProductType       string `json:"productType"`
	Title             string `json:"title"`
	Brand             string `json:"brand"`
	ProductAttributes struct {
		Gender      string `json:"gender"`
		Season      string `json:"season"`
		ReleaseDate string `json:"releaseDate"`
		RetailPrice int    `json:"retailPrice"`
		Colorway    string `json:"colorway"`
		Color       string `json:"color"`
	} `json:"productAttributes"`
	SizeChart struct {
		AvailableConversions []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"availableConversions"`
		DefaultConversion struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"defaultConversion"`
	} `json:"sizeChart"`
}
