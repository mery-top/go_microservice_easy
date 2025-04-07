package main

import(
	"fmt"
	"net/http"

	"patients_refined/routes"
	"patients_refined/storage"
)

func main(){
	storage.InitDB()
	routes.RegisterPatientRoutes()
	fmt.Println("Server Listening at PORT 8080")
	http.ListenAndServe(":8080", nil)
}