package client

import "testing"

func TestListDirectoryFiles(t *testing.T) {
	dir := "../example/files"
	file := "../example/files/testfile.json"
	GetAllFilesInDirectory(dir)
	getOsFile(file)
}
