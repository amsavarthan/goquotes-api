package handlers

import "quotes-rest-api/db"

//Response struct is used to hold the response
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Count   int         `json:"count,omitempty"`
	Quotes  []db.Quote  `json:"quotes,omitempty"`
	Authors []db.Author `json:"authors,omitempty"`
	Tags    []db.Tag    `json:"tags,omitempty"`
}
