package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Print("App Started")
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("World"))
	})
	http.ListenAndServe(":8080", nil)

	fmt.Print("End")

}
