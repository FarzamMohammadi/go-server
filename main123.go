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

	if dbPool != nil {
		addNewVisitor(dbPool)
	}
}

func initConnection() {

	var (
		dbUser                 = mustGetenv("DB_USER")
		dbPwd                  = mustGetenv("DB_PASS")
		instanceConnectionName = "simple-go-app-315700:us-central1:go-app-postgres"
		dbName                 = mustGetenv("DB_NAME")
	)

	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		socketDir = "/cloudsql"
	}

	var dbURI string
	dbURI = fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", dbUser, dbPwd, dbName, socketDir, instanceConnectionName)

	dbPool, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		fmt.Println("Could not connect to database: sql.Open: %v", err)
	}
	dbPool.AutoMigrate(&visitors{}) //Database migration

}

func addNewVisitor(db *gorm.DB) {
	currentTime := time.Now()
	dbPool.Create(&visitors{Time: currentTime.Format("2006-01-02 15:04:05")})
}
