package main

import(
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Patient struct{
	ID int `json: "id"`
	Name string `json: "name"`
	Age int `json: "age"`
	Disease string `json: "disease"`
}

var patients []Patient
var nextID = 1

func main(){
	http.HandleFunc("/patients", handlePatients)
	http.HandleFunc("/patient/", handlePatientByID)

	fmt.Println("Serving running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

