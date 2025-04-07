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


func handlePatients(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet{
		json.NewEncoder(w).Encode(patients)
		return
	}

	if r.Method == http.MethodPost{
		var p Patient
		err:= json.NewDecoder(r.Body).Decode(&p)
		if err!=nil{
			http.Error(w, "Invalid Input", http.StatusBadRequest)
		}
		p.ID = nextID
		nextID++
		patients = append(patients, p)
		json.NewEncoder(w).Encode(p)

		return

	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func handlePatientByID( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	idStr:= r.URL.Path[len("/patient/"):]
	id, err:= strconv.Atoi(idStr)

	if err!=nil {
		http.Error(w, "Invalid Patient ID",http.StatusBadRequest )
		return
	}

	for _, p := range patients {
		if p.ID == id{
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "Patient not Found", http.StatusNotFound)
}


