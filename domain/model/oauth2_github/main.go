package oauth2_github

func NewClient(clientId string, clientSecret string) *Client {
	return &Client{clientId, clientSecret}
}

type Client struct {
	clientId     string
	clientSecret string
}
