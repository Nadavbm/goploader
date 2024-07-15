package client

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
)

type Client struct {
	File      string
	TargetUrl string
	client    *http.Client
}

func NewClient(file, url string) *Client {
	client := &http.Client{}
	return &Client{
		File:      file,
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

	return os.Open(path.Join(fileDir, c.File))
}

func (c *Client) PostFormDataContentRequest(writer *multipart.Writer, body io.Reader) error {
	// create api request
	fmt.Println("target url", c.TargetUrl)
	req, err := http.NewRequest(http.MethodPost, c.TargetUrl, body)
	if err != nil {
		log.Println("failed to create post request", err)
		return err
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
