package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sql"
	"time"

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

var dbPool *sql.DB
var err error

func main() {

	initConnection()
	fmt.Println("App Started")
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
		dbPort    = mustGetenv("DB_PORT")
		dbName    = mustGetenv("DB_NAME")
	)

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPwd, dbTCPHost, dbPort, dbName)
	dbPool, err := sql.Open("postgres", dbURI)

	if err != nil {
		fmt.Errorf("sql.Open: %v", err)
	}

	dbPool.AutoMigrate(&visitors{}) //Database migration

}

func addNewVisitor(db *gorm.DB) {
	currentTime := time.Now()
	dbPool.Create(&visitors{Time: currentTime.Format("2006-01-02 15:04:05")})
}
