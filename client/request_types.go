package client

import (
	"io"
	"log"
	"net/http"
	"strings"
)

// SendHttpRequest to http service
func (c *Client) SendHttpRequest(contentType, method string, body io.Reader) error {
	req, err := c.prepareHttpRequest(method, body)
	if err != nil {
		log.Println("failed to send http request", err)
		return err
	}

	addHeaders(req, contentType)
	resp, err := c.client.Do(req)
	if err != nil {
		log.Println("failed to send http request", err)
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println("falied to close response", err)
		}
	}()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("the error", err)
	}

	log.Println("http request completed with response,", "status:", string(content), resp.StatusCode)
	return nil
}

// prepareHttpRequest according to request method
func (c *Client) prepareHttpRequest(method string, body io.Reader) (*http.Request, error) {
	var req *http.Request
	var err error
	switch {
	case strings.ToUpper(method) == http.MethodPut:
		req, err = http.NewRequest(http.MethodPost, c.TargetUrl, body)
		if err != nil {
			log.Println("failed to create post request", err)
			return nil, err
		}
	default:
		req, err = http.NewRequest(http.MethodPost, c.TargetUrl, body)
		if err != nil {
			log.Println("failed to create post request", err)
			return nil, err
		}
	}
	return req, nil
}

// addHeaders to http request
func addHeaders(req *http.Request, contentType string) {
	req.Header.Add("Content-Type", contentType)
}
