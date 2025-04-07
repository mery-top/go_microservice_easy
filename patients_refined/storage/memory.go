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
<<<<<<< HEAD
	connStr:= "user=postgres password=p dbname=p sslmode=disable"
=======
	connStr:= "user=p password=p dbname=p sslmode=disable"
>>>>>>> 1214024 (Add postgres)
	var err error

	DB, err:= sql.Open("postgres", connStr)

	if err !=nil{
		log.Fatal("Failed to connect to DB", err)
	}

	err =DB.ping()
	if err!= nil{
		log.Fatal("Database not Reachable:", err)
	}

	fmt.Println("Connected to PostgreSQL database")
	
}