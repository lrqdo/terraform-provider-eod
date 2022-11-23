package db

import (
	"net/http"
	"time"
)

func NewClient(url string) *Client {
	return &Client{
		url: url,
		http: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

type Client struct {
	url  string
	http *http.Client
}
