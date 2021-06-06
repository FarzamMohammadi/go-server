package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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

var dbPool *gorm.DB
var err error

func main() {

	fmt.Println("App Started")
	initConnection()
	http.HandleFunc("/hello", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("World"))
	w.WriteHeader(http.StatusOK)

	if err != nil {
		addNewVisitor(dbPool)
	}
}

func initConnection() {

	var (
		dbUser    = mustGetenv("DB_USER")
		dbPwd     = mustGetenv("DB_PASS")
		dbTCPHost = "127.0.0.1"
		dbPort    = "5432"
		dbName    = mustGetenv("DB_NAME")
	)

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPwd, dbTCPHost, dbPort, dbName)
	dbPool, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	dbPool.AutoMigrate(&visitors{}) //Database migration

}

func addNewVisitor(db *gorm.DB) {
	currentTime := time.Now()
	dbPool.Create(&visitors{Time: currentTime.Format("2006-01-02 15:04:05")})
}
