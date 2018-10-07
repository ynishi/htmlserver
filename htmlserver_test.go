package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestInitHtmlServer(t *testing.T) {
	htmlServer := initHtmlServer(8001, "/html")
	if htmlServer == nil {
		t.Fatal("failed to initServer")
	}
	if htmlServer.Handler == nil {
		t.Fatal("failed to make Handler")
	}
	if htmlServer.Server == nil {
		t.Fatal("failed to make Server")
	}
	expectedPort := "8001"
	if htmlServer.Port != expectedPort {
		t.Fatalf("Port not matched.\nwant: %v\nhave: %v\n", expectedPort, htmlServer.Port)
	}
	expectedPath := "/html"
	if htmlServer.DocPath != expectedPath {
		t.Fatalf("DocPath not matched.\nwant: %v\nhave: %v\n", expectedPath, htmlServer.DocPath)
	}
}

func TestHandler(t *testing.T) {
	testServer := initHtmlServer(8001, "./test/html")
    server := httptest.NewServer(testServer.Handler)
    defer server.Close()
    resp, err := http.Get(server.URL)
    if err != nil {
    	t.Fatal(err)
	}
    if resp.StatusCode != 200 {
		t.Fatalf("StatusCode is not 200: %d\n", resp.StatusCode)
	}
	actual, err :=  ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(  string(actual), "test") {
		t.Errorf("Body not contains test: %s\n", actual)
	}
}