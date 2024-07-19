package main

import (
	"log"

	"github.com/nadavbm/goploader/client"
)

func main() {
	log.Println("starting goploader")
	args := getArgs()
	c, err := setClient(args)
	if err != nil {
		panic(err)
	}

	log.Println("all files", c.Files)
	for k, v := range c.Files {
		log.Println("processing file", k, "content type", v)
		writer, buff, err := client.PrepareFormFile(k)
		if err != nil {
			log.Printf("failed to prepare file for multipart form %s", err)
		}
		log.Println("writer content type", writer.FormDataContentType())
		if err := c.SendHttpRequest(writer.FormDataContentType(), string(*args.method), buff); err != nil {
			log.Printf("failed to send http request %s", err)
		}
	}
}
