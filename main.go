package main

import (
	"log"
	"net/http"
	"os"
	"quotes-rest-api/db"
	"quotes-rest-api/handlers"

	"github.com/gorilla/mux"
)

var quotes []db.Quote
var authors []db.Author

func main() {

	router := mux.NewRouter().StrictSlash(true)
	api := router.PathPrefix("/api/v1").Subrouter()
	router.HandleFunc("/", handlers.HandleHome)
	api.HandleFunc("/", handlers.HandleHome)
	api.HandleFunc("/rand", handlers.GetRandQuote).Methods(http.MethodGet)
	api.HandleFunc("/rand/{author/count}", handlers.GetRand).Methods(http.MethodGet)
	api.HandleFunc("/all/quotes", handlers.GetAllQuotes).Methods(http.MethodGet)
	api.HandleFunc("/all/authors", handlers.GetAuthors).Methods(http.MethodGet)
	api.HandleFunc("/all/quotes/{author}", handlers.GetAllQuotesOfAuthors).Methods(http.MethodGet)
	log.Println("Server running!")

	port := os.Getenv("PORT")
	//Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	log.Fatalln(http.ListenAndServe(":"+port, router))

}
