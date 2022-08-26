package oauth2_github

func (c *Client) GetOwnerId(token string) (string, error) {
	o, err := c.GetOwner(token)
	if err != nil {
		return "", err
	}

	return string(rune(o.Id)), nil
}
