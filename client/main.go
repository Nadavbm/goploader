package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func main() {
	c := newClient("file.txt", "http://localhost:8080/upload")

	file, err := c.getFile()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, 1024)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		log.Printf("error create form file %s", err)
	}

	if _, err := io.Copy(part, bytes.NewReader(buf)); err != nil {
		log.Printf("error copy file content %s", err)
	}
	if err := writer.Close(); err != nil {
		log.Printf("error closing writer %s", err)
	}

	if err := c.postFormDataContentRequest(writer, body); err != nil {
		panic(err)
	}
}

type Client struct {
	file      string
	targetUrl string
	client    *http.Client
}

func newClient(file, url string) *Client {
	client := &http.Client{}
	return &Client{
		file:      file,
		targetUrl: url,
		client:    client,
	}
}

func (c *Client) getFile() (*os.File, error) {
	fileDir, err := os.Getwd()
	if err != nil {
		log.Printf("error get working directory %s", err)
		return nil, err
	}

	return os.Open(path.Join(fileDir, c.file))
}

func (c *Client) postFormDataContentRequest(writer *multipart.Writer, body io.Reader) error {
	// create api request
	req, err := http.NewRequest(http.MethodPost, c.targetUrl, body)
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

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("the error", err)
	}
	log.Printf("\ncontent: %s\nstatus: %d", string(content), resp.StatusCode)
	return nil
}
