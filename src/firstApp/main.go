package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

type Visit struct {
	gorm.Model

	Time string
}

func main() {

	//Server
	fmt.Println("App Started")
	http.HandleFunc("/hello", handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	//Database
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	db, err := gorm.Open("postgres", "host=db port=7567 user=postgres dbname=postgres sslmode=disable password=postgres")
	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("World"))
	w.WriteHeader(http.StatusOK)

	router := mux.NewRouter()
	router.HandleFunc("/hello", AddTime).Methods("POST", "")
}
