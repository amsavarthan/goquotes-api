package handlers

import "quotes-rest-api/db"

//Response struct is used to hold the response
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Count   int         `json:"count"`
	Quotes  []db.Quote  `json:"quotes,omitempty"`
	Authors []db.Author `json:"authors,omitempty"`
	Tag     []db.Tag    `json:"tags,omitempty"`
}
