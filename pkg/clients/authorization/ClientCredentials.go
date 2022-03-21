package authorization

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ClientCredentials struct {
	clientID     string
	clientSecret string
	tokenURL     string
}

func (c ClientCredentials) SetAuthorization() (Authorization, error) {
	payload := strings.NewReader(fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", c.clientID, c.clientSecret))

	req, err := http.NewRequest("POST", c.tokenURL, payload)
	if err != nil {
		return nil, fmt.Errorf("in SetAuthorization: %w", err)
	}

	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("in SetAuthorization: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("in SetAuthorization: %w", errors.New(resp.Status))
	}
	defer resp.Body.Close()

	byteValue, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("in SetAuthorization: %w", err)
	}

	var accessToken AccessToken
	err = json.Unmarshal(byteValue, &accessToken)
	if err != nil {
		return nil, fmt.Errorf("in SetAuthorization: %w", err)
	}

	return accessToken, nil
}

func NewClientCredentials(clientID, clientSecret, tokenURL string) ClientCredentials {
	return ClientCredentials{clientID, clientSecret, tokenURL}
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func (a AccessToken) Authorize(r *http.Request) {
	r.Header.Set("Authorization", a.TokenType+" "+a.AccessToken)
}
