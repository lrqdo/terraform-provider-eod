package db

func NewClient(url string) *Client {
	return &Client{
		url: url,
	}
}

type Client struct {
	url string
}
