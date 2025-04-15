package stockxgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

var (
	AuthEndpoint = "https://accounts.stockx.com/oauth/token"
)

func (s *stockXClient) Authenticate() error {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", s.clientID)
	data.Set("client_secret", s.clientSecret)
	data.Set("code", s.code)
	data.Set("redirect_uri", "https://localhost:3000")

	req, err := http.NewRequest(http.MethodPost, AuthEndpoint, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to authenticate: %s", string(body))
	}

	var authResp AuthResponse
	if err := json.Unmarshal(body, &authResp); err != nil {
		return err
	}

	s.Session.AccessToken = authResp.AccessToken
	s.Session.RefreshToken = authResp.RefreshToken
	s.Session.ExpiresIn = authResp.ExpiresIn

	go func() {
		// automatically refresh the token when it expires
		for {
			time.Sleep(time.Duration(s.Session.ExpiresIn) * time.Second)
			if err := s.RefreshToken(); err != nil {
				log.Printf("failed to refresh token: %s", err)
			}
		}
	}()

	return nil
}

func (s *stockXClient) RefreshToken() error {
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("client_id", s.clientID)
	data.Set("client_secret", s.clientSecret)
	data.Set("refresh_token", s.Session.RefreshToken)
	data.Set("audience", "gateway.stockx.com")

	req, err := http.NewRequest(http.MethodPost, AuthEndpoint, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to refresh token: %s", string(body))
	}

	var refreshResp RefreshResponse
	if err := json.Unmarshal(body, &refreshResp); err != nil {
		return err
	}

	s.Session.AccessToken = refreshResp.AccessToken
	s.Session.ExpiresIn = refreshResp.ExpiresIn

	return nil
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	IDToken     string `json:"id_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}
