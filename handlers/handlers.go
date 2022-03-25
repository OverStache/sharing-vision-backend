package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleReq() {
	log.Println("Start development server localhost:5006")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", HomePage)
	myRouter.HandleFunc("/user", Create).Methods("OPTIONS", "POST")
	myRouter.HandleFunc("/users", GetUsers).Methods("OPTIONS", "GET")
	myRouter.HandleFunc("/user/{id}", GetUser).Methods("OPTIONS", "GET")

	log.Fatal(http.ListenAndServe(":5006", myRouter))
}
