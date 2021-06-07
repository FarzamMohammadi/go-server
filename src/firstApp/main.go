package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Visits struct {
	gorm.Model
	ID   uint
	Time string
}

var db *gorm.DB
var err error

func main() {

	fmt.Println("App Started")
	initConnection()
	http.HandleFunc("/hello", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("World"))
	w.WriteHeader(http.StatusOK)

	if db != nil {
		addNewVisitor(db)
	}
}

func initConnection() {

	var (
		dbUser = mustGetenv("DB_USER")
		dbPwd  = mustGetenv("DB_PASS")
		host   = "localhost"
		port   = "5432"
		dbName = mustGetenv("DB_NAME")
	)

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s sslmode=disable", host, dbUser, dbName, port, dbPwd)

	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		fmt.Println("DB Not Connected!: ", err)
	} else {
		fmt.Println("DB Connected!")
	}

	db.AutoMigrate(&Visits{})
}

func addNewVisitor(db *gorm.DB) {
	currentTime := time.Now()
	db.Create(&Visits{Time: currentTime.Format("2006-01-02 15:04:05")})
	fmt.Println("row added!")
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}
