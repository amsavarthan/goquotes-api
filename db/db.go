package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func GetQuotesFromAsset() []Quote {

	ret := make([]Quote, 0, 0)
	file, err := os.Open("./assets/quotes.json")

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	data, _ := ioutil.ReadAll(file)
	json.Unmarshal(data, &ret)
	return ret

}

func GetAuthorsFromAsset() []Author {

	ret := make([]Author, 0, 0)

	file, err := os.Open("./assets/authors.json")

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	data, _ := ioutil.ReadAll(file)
	json.Unmarshal(data, &ret)

	return ret

}

func GetTagsFromAsset() []Tag {

	ret := make([]Tag, 0, 0)

	file, err := os.Open("./assets/tags.json")

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	data, _ := ioutil.ReadAll(file)
	json.Unmarshal(data, &ret)

	return ret

}
