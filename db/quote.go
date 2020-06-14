package db

//Quote struct is used to hold each quote
type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
	Tag    string `json:"tag"`
}
