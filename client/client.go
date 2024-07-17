package client

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
)

type Files []string

type Client struct {
	Files     Files
	TargetUrl string
	client    *http.Client
}

func NewClient(url string, files Files) *Client {
	client := &http.Client{}
	return &Client{
		Files:     files,
		TargetUrl: url,
		client:    client,
	}
}

func (c *Client) GetFile() (*os.File, error) {
	fileDir, err := os.Getwd()
	if err != nil {
		log.Printf("error get working directory %s", err)
		return nil, err
	}

	return os.Open(path.Join(fileDir, c.Files[0]))
}

func (c *Client) Send(method string, writer *multipart.Writer, body io.Reader) error {
	var req *http.Request
	var err error
	switch {
	case strings.ToUpper(method) == http.MethodPut:
		req, err = http.NewRequest(http.MethodPost, c.TargetUrl, body)
		if err != nil {
			log.Println("failed to create post request", err)
			return err
		}
	default:
		req, err = http.NewRequest(http.MethodPost, c.TargetUrl, body)
		if err != nil {
			log.Println("failed to create post request", err)
			return err
		}
	}

	// content type
	req.Header.Add("Content-Type", writer.FormDataContentType())

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
	log.Printf("\ncontent: %s\nstatus: %d", string(content), resp.StatusCode)
	return nil
}
