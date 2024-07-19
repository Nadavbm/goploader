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

	if client.Files["../example/files/testfile.json"] != "application/json" {
		t.Errorf("expected content type to be application/json instead got %s", client.Files["../example/files/testfile.json"])
	}

	if client.Files["../example/files/testfile2.json"] != "application/json" {
		t.Errorf("expected content type to be application/json instead got %s", client.Files["../example/files/testfile2.json"])
	}

	client = nil
	args.dir = nil
	args.file = &file
	client, err = setClient(args)
	if err != nil {
		t.Fatal("could not set client", err)
	}
	if client.Files["../example/files/testfile3.json"] != "application/json" {
		t.Errorf("expected content type to be application/json instead got %s", client.Files["../example/files/testfile3.json"])
	}
}
