package db

//Tag struct is used to hold authors data
type Tag struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}
