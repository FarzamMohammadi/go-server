package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Visitors struct {
	gorm.Model
	ID   uint
	Time string
}

var db *gorm.DB

var err error

func main() {

	connectDB()
	//Server
	fmt.Println("App Started")
	http.HandleFunc("/hello", handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func connectDB() {

	dsn := "host=localhost user=postgres password=password dbname=db1 port=6666 sslmode=disable TimeZone=US/Eastern"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Visitors{}) //Database migration

}

func addNewVisitor(db *gorm.DB) {
	currentTime := time.Now()
	db.Create(&Visitors{Time: currentTime.Format("2006-01-02 15:04:05")})
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("World"))
	w.WriteHeader(http.StatusOK)

	addNewVisitor(db)
}
