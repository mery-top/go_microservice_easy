package main

import(
	"fmt"
	"net/http"

	"patients_refined/routes"
)

func main(){
	routes.RegisterPatientRoutes()
	fmt.Println("Server Listening at PORT 8080")
	http.ListenAndServe(":8080", nil)
}