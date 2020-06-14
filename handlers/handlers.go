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
		log.Fatalln("Error parsing data")
	}
}

func HandleHome(w http.ResponseWriter, r *http.Request) {

	//setting the header
	w.Header().Set("Content-Type", "application/json")

	//Encoding quote entity into JSON
	res, _ := json.Marshal(Response{
		Status:  http.StatusOK,
		Message: "success",
	})

	//Set status as OK and show the JSON
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetAllQuotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if len(quotes) == 0 {

		w.WriteHeader(http.StatusNoContent)
		res, _ := json.Marshal(Response{
			Status:  http.StatusNoContent,
			Message: "no content",
			Count:   len(quotes),
			Quotes:  quotes,
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

func GetQuote(w http.ResponseWriter, r *http.Request, count int) {

	w.Header().Set("Content-Type", "application/json")

	if count > len(quotes) {

		//If the entered count is exceeding the total count of quotes available
		w.WriteHeader(http.StatusNotAcceptable)
		res, _ := json.Marshal(Response{
			Status:  http.StatusNotAcceptable,
			Message: "max count exceeded",
			Count:   0,
			Quotes:  make([]db.Quote, 0, 0),
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

func GetRandQuote(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	resQuotes := make([]db.Quote, 0, 0)

	//Used to generate a random number from 0- @code{len(quotes)}
	rand.Seed(time.Now().Unix())
	resQuotes = append(resQuotes, quotes[rand.Intn(len(quotes))])

	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(Response{
		Status:  http.StatusOK,
		Message: "success",
		Count:   len(resQuotes),
		Quotes:  resQuotes,
	})
	w.Write(res)

}

func GetAuthors(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if len(authors) == 0 {

		w.WriteHeader(http.StatusNoContent)
		res, _ := json.Marshal(Response{
			Status:  http.StatusNoContent,
			Message: "no content",
			Count:   len(authors),
			Authors: authors,
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

func GetRandAuthorQuote(w http.ResponseWriter, r *http.Request, author string) {

	w.Header().Set("Content-Type", "application/json")
	//used to store the filtered quotes
	authorQuotes := make([]db.Quote, 0, 0)
	//used to get random quote from filtered author slice
	resQuotes := make([]db.Quote, 0, 0)

	//filtering quotes based on author query
	for _, v := range quotes {
		if v.Author == author {
			authorQuotes = append(authorQuotes, v)
		}
	}

	//Used to generate a random number from 0- @code{len(authorQuotes)}
	rand.Seed(time.Now().Unix())
	resQuotes = append(resQuotes, authorQuotes[rand.Intn(len(authorQuotes))])

	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(Response{
		Status:  http.StatusOK,
		Message: "success",
		Count:   len(resQuotes),
		Quotes:  resQuotes,
	})
	w.Write(res)

}

func GetRand(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	//If there is any parsing error then it should be string so sending it to
	//GetRandAuthorQuote func as Author name
	if count, err := strconv.Atoi(params["author/count"]); err != nil {
		GetRandAuthorQuote(w, r, params["author/count"])
	} else {
		GetQuote(w, r, count)
	}

}

func GetAllQuotesOfAuthors(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	author := params["author"]

	w.Header().Set("Content-Type", "application/json")
	//used to store the filtered quotes
	resQuotes := make([]db.Quote, 0, 0)

	//filtering quotes based on author query
	for _, v := range quotes {
		if v.Author == author {
			resQuotes = append(resQuotes, v)
		}
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
