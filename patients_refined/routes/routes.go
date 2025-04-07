package routes

import (
	"net/http"
	"patient_refined/handlers"
	"patients_refined/handlers"
)

func RegisterPatientRoutes(){
	http.HandleFunc("/patients", func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodGet{
			handlers.GetAllPatients(w,r)
		}else if r.Method == http.MethodPost{
			handlers.AddPatient(w,r)
		}else{
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/patient/", handlers.GetPatientByID)
}