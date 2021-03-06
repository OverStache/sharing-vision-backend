package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func HandleReq() {
	log.Println("Start development server localhost:8000")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", HomePage)
	myRouter.HandleFunc("/article", CreateArticle).Methods("OPTIONS", "POST")
	myRouter.HandleFunc("/article/{limit}/{offset}", GetArticles).Methods("OPTIONS", "GET")
	myRouter.HandleFunc("/article/{id}", GetArticle).Methods("OPTIONS", "GET")
	myRouter.HandleFunc("/article/{id}", UpdateArticle).Methods("OPTIONS", "PUT")
	myRouter.HandleFunc("/article/{id}", DeleteArticle).Methods("OPTIONS", "DELETE")

	handler := cors.AllowAll().Handler(myRouter)
	log.Fatal(http.ListenAndServe(":8000", handler))
}
