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
	for i, f := range c.Files {
		log.Println("processing file", i, f)
		writer, buff, err := client.PrepareFormFile(f)
		if err != nil {
			log.Printf("failed to prepare file for multipart form %s", err)
		}
		if err := c.SendHttpRequest(writer.FormDataContentType(), string(*args.method), buff); err != nil {
			log.Printf("failed to send http request %s", err)
		}
	}
}
