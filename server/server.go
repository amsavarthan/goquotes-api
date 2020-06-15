package server

import (
	"log"
	"net/http"
	"os"
	"quotes-rest-api/handlers"

	"github.com/gorilla/mux"
)

//InitializeServer is used to establish connection with server
func InitializeServer() error {

	router := mux.NewRouter().StrictSlash(true)
	api := router.PathPrefix("/api/v1").Subrouter()

	//Home page
	router.HandleFunc("/", handlers.HandleHome)
	api.HandleFunc("/", handlers.HandleHome)

	//Get random quotes
	api.Path("/random").Queries("count", "{count:[0-9]+}").HandlerFunc(handlers.GetRandQuotes).Methods(http.MethodGet)
	api.Path("/random/{count:[0-9]+}").Queries("type", "{type}").Queries("val", "{val}").HandlerFunc(handlers.HandleTypeQuery).Methods(http.MethodGet)

	//Get all datas
	api.HandleFunc("/all/quotes", handlers.GetAllQuotes).Methods(http.MethodGet)
	api.HandleFunc("/all/authors", handlers.GetAuthors).Methods(http.MethodGet)
	api.HandleFunc("/all/tags", handlers.GetTags).Methods(http.MethodGet)

	//Get all filtered datas
	api.Path("/all").Queries("type", "{type}").Queries("val", "{val}").HandlerFunc(handlers.HandleTypeQueryToGetAll).Methods(http.MethodGet)

	//Get port from .env file
	//we did not specify any port so this should return an empty string when tested locally
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	log.Println("Server running!")
	return http.ListenAndServe(":"+port, router)
}
