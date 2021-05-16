package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResponse(t *testing.T) {

	expectedResult := "World"
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	resp, err := http.Get(server.URL)

	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	}

	result, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}
	if string(result) != expectedResult {
		t.Fatal("Did Not Recieve Expected Byte: 'World'")
	}
}
