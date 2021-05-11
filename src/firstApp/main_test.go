package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestResponse(t *testing.T) {
	url := "http://localhost:8080/hello"

	var client http.Client
	resp, err := client.Get(url)

	if err != nil {
		t.Error("ERRPR: App Is Not Running!")
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error("ERROR: Problem Reading Respone")
		}

		bodyString := string(bodyBytes)

		if bodyString != "World" {
			t.Error("ERROR: App Is Running But Did Not Receive Expcted Respone: 'World'")
		}
	}
}
