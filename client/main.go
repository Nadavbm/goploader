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
	"path/filepath"
)

type Client struct {
	client *http.Client
}

func main() {
	// prepare file for upload post request
	fileDir, err := os.Getwd()
	if err != nil {
		log.Printf("error get directory %s", err)
	}
	fileName := "file.txt"
	filePath := path.Join(fileDir, fileName)
	log.Printf("open file %s", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("error open file %s", err)
	}
	defer file.Close()

	buf := make([]byte, 1024)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		log.Printf("error create form file %s", err)
	}
	log.Printf("create form file part %s", part)

	if _, err := io.Copy(part, bytes.NewReader(buf)); err != nil {
		log.Printf("error copy file content %s", err)
	}
	if err := writer.Close(); err != nil {
		log.Printf("error closing writer %s", err)
	}

	fmt.Printf("part and buf %s %s", part, string(buf))

	// create api request
	target_url := "http://localhost:8080/upload"
	req, err := http.NewRequest(http.MethodPost, target_url, body)
	if err != nil {
		log.Println("failed to create post request", err)
	}

	// auth
	c := Client{}
	c.client = &http.Client{}

	// content type
	req.Header.Add("Content-Type", writer.FormDataContentType())

	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Println("error1", err)
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
}
