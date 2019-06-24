package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rjohnt/SampleGoApp/app"
	"github.com/rjohnt/SampleGoApp/controllers"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/me/contacts", controllers.GetContacts).Methods("GET")
	router.HandleFunc("/api/me/contacts", controllers.CreateContact).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		fmt.Print(err)
	}
}