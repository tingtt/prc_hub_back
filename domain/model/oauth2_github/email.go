package oauth2_github

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type OwnerEmail struct {
	Email      string `json:"email"`
	Verified   bool   `json:"verified"`
	Primary    bool   `json:"primary"`
	Visivility bool   `json:"visibility"`
}

func (c *Client) GetOwnerEmails(token string) (emails []OwnerEmail, err error) {
	// GET github api
	req, err := http.NewRequest("GET", "https://api.github.com/user/emails", nil)
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

	// Read response body
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	// Check status code
	if res.StatusCode != http.StatusOK {
		err = errors.New(string(bodyBytes))
		return
	}

	// Unmarshal response body
	err = json.Unmarshal(bodyBytes, &emails)
	if err != nil {
		return
	}

	return
}

func (c *Client) GetOwnerPrimaryEmail(token string) (e OwnerEmail, err error) {
	// Get owner emails
	emails, err := c.GetOwnerEmails(token)
	if err != nil {
		return
	}

	// Find primary email
	found := false
	for _, email := range emails {
		if email.Primary {
			e = email
			found = true
			break
		}
	}
	if !found {
		err = errors.New("primary email not found in body of response from \"api.github.com\"")
		return
	}

	return e, nil
}
