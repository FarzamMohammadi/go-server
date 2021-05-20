package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("App Started")
	http.HandleFunc("/hello", handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("World - Testing"))
	w.WriteHeader(http.StatusOK)
}
