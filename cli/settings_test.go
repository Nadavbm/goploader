package main

import "testing"

func TestSetClient(t *testing.T) {
	dir := "../example/files"
	file := "../example/files/testfile3.json"
	url := "http://localhost:8080/upload"
	args := Args{
		url: &url,
	}
	args.dir = &dir
	client, err := setClient(args)
	if err != nil {
		t.Fatal("could not set client", err)
	}

	if client.Files[0] != "../example/files/testfile.json" {
		t.Errorf("expected file name to be ../example/files/testfile.json instead got %s", client.Files[0])
	}

	if client.Files[1] != "../example/files/testfile2.json" {
		t.Errorf("expected file name to be ../example/files/testfile2.json instead got %s", client.Files[1])
	}

	client = nil
	args.dir = nil
	args.file = &file
	client, err = setClient(args)
	if err != nil {
		t.Fatal("could not set client", err)
	}
	if client.Files[0] != "../example/files/testfile3.json" {
		t.Errorf("expected file name to be ../example/files/testfile3.json instead got %s", client.Files[0])
	}
}
