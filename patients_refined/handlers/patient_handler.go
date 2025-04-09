package handlers

import(
	"encoding/json"
	"net/http"
	"strconv"

	"patients_refined/models"
	"patients_refined/storage"
)

func GetAllPatients(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.Patients)
}

func GetPatientByID(w http.ResponseWriter, r *http.Request){
	idStr:= r.URL.Path[len("/patient/"):]
	id,err:= strconv.Atoi(idStr)

	if err!=nil{
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, p := range storage.Patients{
		if p.ID == id {
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "Patient not found", http.StatusNotFound)
}


func AddPatient(w http.ResponseWriter, r * http.Request){
	var p models.Patient
	err:=json.NewDecoder(r.Body).Decode(&p)

	if err != nil{
		http.Error(w, "Invalid data", http.StatusBadRequest)
	}

	query:= 'INSERT INTO patients (name, age, disease) VALUES ($1, $2, $3) '
}