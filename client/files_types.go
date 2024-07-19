package client

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

// allowedFileSuffixes are the allowed suffix to be sent with goploader
// TODO: change this list to a map[string]string and define content length according to https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types
var allowedFileSuffixes = []string{
	".json", ".pdf", ".jpg", ".txt",
}

var allowedFileSuffixesAndContentLegnth = map[string]string{
	".json": "application/json",
	".png":  "image/png",
	".pdf":  "application/pdf",
	".txt":  "text/plain",
}

// PrepareFormFile to be sent later with closed multipart and data buffer
func PrepareFormFile(filename string) (*multipart.Writer, *bytes.Buffer, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		log.Printf("error writing to buffer %s", err)
		return nil, nil, err
	}

	file, err := getOsFile(filename)
	if err != nil {
		log.Printf("failed to get file %s", err)
		return nil, nil, err
	}

	if _, err := io.Copy(fileWriter, file); err != nil {
		log.Printf("error copy file content %s", err)
		return nil, nil, err
	}

	if err := bodyWriter.Close(); err != nil {
		log.Printf("error closing writer. failed to set multipart boundaries %s", err)
		return nil, nil, err
	}

	return bodyWriter, bodyBuf, nil
}

// GetAllFilesInDirectory returns a slice of file names (full path)
func GetAllFilesInDirectory(dir string) (map[string]string, error) {
	fileNames := make(map[string]string)
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Printf("failed to read directory %s", err)
		return nil, err
	}

	for _, file := range entries {
		if checkFileSuffix(file.Name()) {
			fileNames[path.Join(dir, file.Name())] = GetFileSuffixContentType(file.Name())
		}
	}
	return fileNames, nil
}

// getOsFile
func getOsFile(file string) (*os.File, error) {
	fileDir, err := os.Getwd()
	if err != nil {
		log.Printf("error get working directory %s", err)
		return nil, err
	}

	return os.Open(path.Join(fileDir, file))
}

// GetFileSuffixContentType
func GetFileSuffixContentType(filename string) string {
	for k, v := range allowedFileSuffixesAndContentLegnth {
		if strings.HasSuffix(filename, k) {
			return v
		}
	}
	return "application/octet-stream"
}

// checkFileSuffix check if file has allowed suffix
func checkFileSuffix(filename string) bool {
	for _, s := range allowedFileSuffixes {
		if strings.HasSuffix(filename, s) {
			return true
		}
	}
	return false
}
