package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
)

func main() {
	c := newClient("file.txt", "http://localhost:8080/upload")

	file, err := c.getFile()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("file", c.file)
	if err != nil {
		fmt.Println("error writing to buffer")
	}

	if _, err := io.Copy(fileWriter, file); err != nil {
		log.Printf("error copy file content %s", err)
	}
	if err := bodyWriter.Close(); err != nil {
		log.Printf("error closing writer %s", err)
	}
	fmt.Println("body", string(bodyBuf.Bytes()))
	if err := c.postFormDataContentRequest(bodyWriter, bodyBuf); err != nil {
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
