package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go-contact/controllers"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()
	//router.Use(app.JwtAuthentication) //attach JWT auth middleware

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8080" //localhost
	}



	fmt.Println(port)
	router.HandleFunc("/", controllers.GetContacts).Methods("GET")
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	err := http. ListenAndServe(":"+port,handlers.CORS(headersOk, originsOk, methodsOk) (router)) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}

}
