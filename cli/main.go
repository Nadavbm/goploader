package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"

	"github.com/nadavbm/goploader/client"
)

type Args struct {
	dir    *string
	file   *string
	url    *string
	method *string
	form   *string
}

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
	log.Println("form arg", *args.form)

	c := client.NewClient(string(*args.file), string(*args.url))

	file, err := c.GetFile()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("file", c.File)
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
	if err := c.PostFormDataContentRequest(bodyWriter, bodyBuf); err != nil {
		panic(err)
	}
}

func getArgs() Args {
	args := Args{
		dir:    flag.String("dir", "", "Go uploader files directory"),
		file:   flag.String("file", "", "Go uploader file"),
		url:    flag.String("url", "", "api server url for goploader"),
		method: flag.String("method", "", "http request method"),
		form:   flag.String("form", "", "request form type"),
	}
	flag.Parse()
	return args
}
