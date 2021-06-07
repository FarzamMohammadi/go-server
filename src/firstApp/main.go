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

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

type visitors struct {
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
		dbUser                 = mustGetenv("DB_USER")
		dbPwd                  = mustGetenv("DB_PASS")
		instanceConnectionName = "simple-go-app-315700:us-central1:go-app-postgres"
		dbName                 = mustGetenv("DB_NAME")
	)

	var dbURI string
	dbURI = fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", instanceConnectionName, dbUser, dbName, dbPwd)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "cloudsqlpostgres",
		DSN:        dbURI,
	}), &gorm.Config{})

	if err != nil {
		fmt.Errorf("sql.Open: %v", err)
	} else {
		fmt.Println("Finally - db connected!")
	}
	currentTime := time.Now()
	db.Create(&visitors{Time: currentTime.Format("2006-01-02 15:04:05")})
	fmt.Println("row added!")
}

func addNewVisitor(db *gorm.DB) {
	currentTime := time.Now()
	db.Create(&visitors{Time: currentTime.Format("2006-01-02 15:04:05")})
}
