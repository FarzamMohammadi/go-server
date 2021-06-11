package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Visits struct {
	gorm.Model
	ID   uint
	Time string
}

var DB *gorm.DB

func main() {

	fmt.Println("App Started")
	initConnection()
	http.HandleFunc("/hello", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.WithError(err).Fatal("Could Not Connect To Server!")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("World"))
	w.WriteHeader(http.StatusOK)

	if DB != nil {
		addNewVisitor()
	} else {
		log.Info("Problem Inserting Row - Row Was Not Added")
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
		log.WithError(err).Fatal("Database could not connect!")
	} else {
		log.Info("Database Connected!")
	}

	DB = db
	DB.AutoMigrate(&Visits{})
}

func addNewVisitor() {
	currentTime := time.Now()
	DB.Create(&Visits{Time: currentTime.Format("2006-01-02 15:04:05")})
	fmt.Println("Row Added!")
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s Environment Variable Not Set.\n", k)
	}
	return v
}
