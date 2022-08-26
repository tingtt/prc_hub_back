package oauth2_github

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Owner struct {
	Name      string `json:"login"`
	Id        uint64 `json:"id"`
	AvatarUrl string `json:"avatar_url"`
}

func (c *Client) GetOwner(token string) (o Owner, err error) {
	// GET github api
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	// Check status code
	if res.StatusCode != http.StatusOK {
		return Owner{}, errors.New("failed to get owner informations")
	}

	// Read response body
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	// Check status code
	if res.StatusCode != http.StatusOK {
		return Owner{}, errors.New(string(bodyBytes))
	}

	// Unmarshal response body
	err = json.Unmarshal(bodyBytes, &o)
	if err != nil {
		return
	}

	return
}
