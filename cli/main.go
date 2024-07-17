package main

import (
	"flag"
	"log"
	"os"

	"github.com/nadavbm/goploader/client"
)

func main() {
	log.Println("starting goploader")
	for i, cmd := range os.Args {
		log.Println("\n index: ", i, "arg: ", cmd)
	}
	log.Println("args", os.Args)

	args := getArgs()
	flag.Parse()
	log.Println("dir arg", *args.dir)
	log.Println("file arg", *args.file)
	log.Println("url arg", *args.url)
	log.Println("method arg", *args.method)
	log.Println("content arg", *args.content)

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
		if err := c.Send(string(*args.method), writer, buff); err != nil {
			log.Printf("failed to send http request %s", err)
		}
	}
}
