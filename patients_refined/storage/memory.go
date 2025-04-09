package storage

import (
	"patients_refined/models"
	"database/sql"
	"fmt"
	"log"
	"github.com/lib/pq"
)

var DB *sql.DB

func InitDB(){
	connStr := "user=username password=password dbname=dbname sslmode=disable"
	var err error

	DB, err:= sql.Open("postgres", connStr)

	if err !=nil{
		log.Fatal("Failed to connect to DB", err)
	}

	err =DB.Ping()
	if err!= nil{
		log.Fatal("Database not Reachable:", err)
	}

	fmt.Println("Connected to PostgreSQL database")
	
}