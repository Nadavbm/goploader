package main

import (
	"flag"

	"github.com/nadavbm/goploader/client"
)

// Args provided while executing the command. e.g. ./goploader --file=example/files/testfile.json --url=http://localhost:8080/upload --method=post
type Args struct {
	dir     *string // directory (multipart req)
	file    *string // filename (single req)
	url     *string // server url
	method  *string // request method
	content *string // mime type https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types
}

// setClient create a client based on cli arguments
func setClient(args Args) (*client.Client, error) {
	var files []string
	var err error
	if args.dir != nil {
		if string(*args.dir) != "" {
			files, err = client.GetAllFilesInDirectory(*args.dir)
			if err != nil {
				return nil, err
			}
		}
	}

	if args.file != nil {
		if string(*args.file) != "" {
			files = append(files, string(*args.file))
		}
	}

	return client.NewClient(string(*args.url), files), nil
}

// getArgs from command line
func getArgs() Args {
	args := Args{
		dir:     flag.String("dir", "", "Go uploader files directory"),
		file:    flag.String("file", "", "Go uploader file"),
		url:     flag.String("url", "", "api server url for goploader"),
		method:  flag.String("method", "", "http request method"),
		content: flag.String("content", "", "request mime type"),
	}
	flag.Parse()
	return args
}
