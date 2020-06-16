package handlers

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"quotes-rest-api/db"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var quotes []db.Quote
var authors []db.Author
var tags []db.Tag

func init() {
	quotes = db.GetQuotesFromAsset()
	authors = db.GetAuthorsFromAsset()
	tags = db.GetTagsFromAsset()

	if len(quotes) <= 0 || len(authors) <= 0 || len(tags) <= 0 {
		log.Fatalln("Error getting data!")
	}
}

//HandleHome is used to handle home
func HandleHome(w http.ResponseWriter, r *http.Request) {

	//setting the header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//Encoding quote entity into JSON
	res, _ := json.Marshal(Response{
		Status:  http.StatusOK,
		Message: "success",
	})

	//Set status as OK and show the JSON
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// GetAllQuotes returns all quotes from original data
func GetAllQuotes(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if len(quotes) == 0 {

		w.WriteHeader(http.StatusNoContent)
		res, _ := json.Marshal(Response{
			Status:  http.StatusNoContent,
			Message: "no content",
		})
		w.Write(res)

	} else {

		w.WriteHeader(http.StatusOK)
		res, _ := json.Marshal(Response{
			Status:  http.StatusOK,
			Message: "success",
			Count:   len(quotes),
			Quotes:  quotes,
		})
		w.Write(res)

	}

}

// GetAuthors returns all authors from original data
func GetAuthors(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if len(authors) == 0 {

		w.WriteHeader(http.StatusNoContent)
		res, _ := json.Marshal(Response{
			Status:  http.StatusNoContent,
			Message: "no content",
		})
		w.Write(res)

	} else {

		w.WriteHeader(http.StatusOK)
		res, _ := json.Marshal(Response{
			Status:  http.StatusOK,
			Message: "success",
			Count:   len(authors),
			Authors: authors,
		})
		w.Write(res)

	}

}

// GetTags returns all tags from original data
func GetTags(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if len(tags) == 0 {

		w.WriteHeader(http.StatusNoContent)
		res, _ := json.Marshal(Response{
			Status:  http.StatusNoContent,
			Message: "no content",
		})
		w.Write(res)

	} else {

		w.WriteHeader(http.StatusOK)
		res, _ := json.Marshal(Response{
			Status:  http.StatusOK,
			Message: "success",
			Count:   len(tags),
			Tags:    tags,
		})
		w.Write(res)

	}

}

//GetRandQuotes returns random quote from original data
func GetRandQuotes(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	count, _ := strconv.Atoi(mux.Vars(r)["count"])

	if count > len(quotes) {

		//If the entered count is exceeding the total count of quotes available
		w.WriteHeader(http.StatusNotAcceptable)
		res, _ := json.Marshal(Response{
			Status:  http.StatusNotAcceptable,
			Message: "max count exceeded",
		})
		w.Write(res)

	} else {

		resQuotes := make([]db.Quote, 0, 0)

		rand.Seed(time.Now().Unix())
		for i := 0; i < count; i++ {
			//gets random quote and appends to resQuotes slice
			resQuotes = append(resQuotes, quotes[rand.Intn(len(quotes))])
		}

		w.WriteHeader(http.StatusOK)
		res, _ := json.Marshal(Response{
			Status:  http.StatusOK,
			Message: "success",
			Count:   len(resQuotes),
			Quotes:  resQuotes,
		})
		w.Write(res)

	}
}

/*
GetAllQuotesOfAuthors is used to get all quotes where author equals to passed author param
First we filter the original quotes data and stored the filtered data in resQuotes slice then display all values
*/
func GetAllQuotesOfAuthors(w http.ResponseWriter, r *http.Request, author string) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//used to store the filtered quotes
	resQuotes := make([]db.Quote, 0, 0)

	//filtering quotes based on author query
	for _, v := range quotes {
		if v.Author == author {
			resQuotes = append(resQuotes, v)
		}
	}

	if len(resQuotes) == 0 {

		w.WriteHeader(http.StatusNoContent)
		res, _ := json.Marshal(Response{
			Status:  http.StatusNoContent,
			Message: "no content",
		})
		w.Write(res)

	} else {

		w.WriteHeader(http.StatusOK)
		res, _ := json.Marshal(Response{
			Status:  http.StatusOK,
			Message: "success",
			Count:   len(resQuotes),
			Quotes:  resQuotes,
		})
		w.Write(res)

	}

}

/*
GetAllQuotesOfTag is used to get all quotes where tag equals to passed tag param

First we filter the original quotes data and stored the filtered data in resQuotes slice then display all values
*/
func GetAllQuotesOfTag(w http.ResponseWriter, r *http.Request, tag string) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//used to store the filtered quotes
	resQuotes := make([]db.Quote, 0, 0)

	//filtering quotes based on author query
	for _, v := range quotes {
		if v.Tag == tag {
			resQuotes = append(resQuotes, v)
		}
	}

	if len(resQuotes) == 0 {

		w.WriteHeader(http.StatusNoContent)
		res, _ := json.Marshal(Response{
			Status:  http.StatusNoContent,
			Message: "no content",
		})
		w.Write(res)

	} else {

		w.WriteHeader(http.StatusOK)
		res, _ := json.Marshal(Response{
			Status:  http.StatusOK,
			Message: "success",
			Count:   len(resQuotes),
			Quotes:  resQuotes,
		})
		w.Write(res)

	}

}

/*
GetRandQuotesByAuthor is used to get random quotes where author equals to passed author param

First we filter the original quotes data and stored the filtered data in resAuthorQuotes slice.
based on pCount param, pCount random numbers are generated and passed as index values to resAuthorQuotes
to get the value at that index.
*/
func GetRandQuotesByAuthor(w http.ResponseWriter, r *http.Request, author string, count int) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//used to store the filtered quotes
	resAuthorQuotes := make([]db.Quote, 0, 0)

	//filtering quotes based on author query
	for _, v := range quotes {
		if v.Author == author {
			resAuthorQuotes = append(resAuthorQuotes, v)
		}
	}

	if len(resAuthorQuotes) == 0 {

		w.WriteHeader(http.StatusNoContent)
		res, _ := json.Marshal(Response{
			Status:  http.StatusNoContent,
			Message: "no content",
		})
		w.Write(res)

	} else {

		if count > len(resAuthorQuotes) {

			//If the entered count is exceeding the total count of quotes available
			w.WriteHeader(http.StatusNotAcceptable)
			res, _ := json.Marshal(Response{
				Status:  http.StatusNotAcceptable,
				Message: "max count exceeded",
			})
			w.Write(res)

		} else {

			resQuotes := make([]db.Quote, 0, 0)

			rand.Seed(time.Now().Unix())
			for i := 0; i < count; i++ {
				//gets random quote and appends to resQuotes slice
				resQuotes = append(resQuotes, resAuthorQuotes[rand.Intn(len(resAuthorQuotes))])
			}

			w.WriteHeader(http.StatusOK)
			res, _ := json.Marshal(Response{
				Status:  http.StatusOK,
				Message: "success",
				Count:   len(resQuotes),
				Quotes:  resQuotes,
			})
			w.Write(res)

		}

	}

}

/*
GetRandQuotesByTag is used to get random quotes where tag equals to passed tag param

First we filter the original quotes data and stored the filtered data in resTagQuotes slice.
based on pCount param, pCount random numbers are generated and passed as index values to resTagQuotes
to get the value at that index.
*/
func GetRandQuotesByTag(w http.ResponseWriter, r *http.Request, tag string, count int) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//used to store the filtered quotes
	resTagQuotes := make([]db.Quote, 0, 0)

	//filtering quotes based on author query
	for _, v := range quotes {
		if v.Tag == tag {
			resTagQuotes = append(resTagQuotes, v)
		}
	}

	if len(resTagQuotes) == 0 {

		w.WriteHeader(http.StatusNoContent)
		res, _ := json.Marshal(Response{
			Status:  http.StatusNoContent,
			Message: "no content",
		})
		w.Write(res)

	} else {

		if count > len(resTagQuotes) {

			//If the entered count is exceeding the total count of quotes available
			w.WriteHeader(http.StatusNotAcceptable)
			res, _ := json.Marshal(Response{
				Status:  http.StatusNotAcceptable,
				Message: "max count exceeded",
			})
			w.Write(res)

		} else {

			resQuotes := make([]db.Quote, 0, 0)

			rand.Seed(time.Now().Unix())
			for i := 0; i < count; i++ {
				//gets random quote and appends to resQuotes slice
				resQuotes = append(resQuotes, resTagQuotes[rand.Intn(len(resTagQuotes))])
			}

			w.WriteHeader(http.StatusOK)
			res, _ := json.Marshal(Response{
				Status:  http.StatusOK,
				Message: "success",
				Count:   len(resQuotes),
				Quotes:  resQuotes,
			})
			w.Write(res)

		}

	}

}

/*
HandleTypeQuery is used to get the params passed and perform query request
based on the values passed.

type can be of values author or tag
val should be a valid text
count should be a valid number, if not an error is returned

based on the type the requested query filter function is passed
*/
func HandleTypeQuery(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	typeQuery := params["type"]
	val := params["val"]
	count, _ := strconv.Atoi(params["count"])

	switch typeQuery {
	case "author":
		GetRandQuotesByAuthor(w, r, val, count)
	case "tag":
		GetRandQuotesByTag(w, r, val, count)
	default:
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(Response{
			Status:  http.StatusBadRequest,
			Message: "invalid request",
		})
		w.Write(res)
	}
}

/*
HandleTypeQueryToGetAll is used to get the params passed and perform query request
based on the values passed.

type can be of values author or tag
val should be a valid text

based on the type the requested query filter function is passed
*/
func HandleTypeQueryToGetAll(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	typeQuery := params["type"]
	val := params["val"]

	switch typeQuery {
	case "author":
		GetAllQuotesOfAuthors(w, r, val)
	case "tag":
		GetAllQuotesOfTag(w, r, val)
	default:
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(Response{
			Status:  http.StatusBadRequest,
			Message: "invalid request",
		})
		w.Write(res)
	}
}
