package oauth2_github

import (
	"fmt"
	"net/url"
)

func (c *Client) GetLoginUrl() string {
	return fmt.Sprintf("https://github.com/login/oauth/authorize?%s", url.QueryEscape(c.clientId))
}
