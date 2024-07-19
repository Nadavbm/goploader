package client

import (
	"net/http"
)

// Files slice of file names inside a directory
type Files map[string]string

// Client send files to a server by target url
type Client struct {
	Files     Files
	TargetUrl string
	client    *http.Client
}

// NewClient creates a new instance of client
func NewClient(url string, files Files) *Client {
	client := &http.Client{}
	return &Client{
		Files:     files,
		TargetUrl: url,
		client:    client,
	}
}
