package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"github.com/tidwall/gjson"
)

type Book struct {
	title string
	author []string
	previewLink string
}

const googleApi = "https://www.googleapis.com/books/v1/volumes"
// TODO to be moved to the service
func fetchBookMatadata(bookTitle string, apiKey string) Book {
	escapedTitle := url.QueryEscape(bookTitle)
    	
	res, err := http.Get(googleApi+"?q="+escapedTitle+"&key="+apiKey+"&projection=lite&maxResults=1")

	if err!= nil {
		fmt.Printf("error encountered while talking to Book Api %v",err)
	}
	defer res.Body.Close()

    responseBytes,readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		fmt.Printf("Could not extract response body %v",readErr)
	}	

	title := gjson.GetBytes(responseBytes,"items.#.volumeInfo.title").String()
	authorResults := gjson.GetBytes(responseBytes,"items.#.volumeInfo.authors|@flatten")
	bookThumb := gjson.GetBytes(responseBytes,"items.#.volumeInfo.imageLinks.smallThumbnail").String()
	
	authors := make([]string,0)

	for _, author := range authorResults.Array() {
		authors = append(authors, author.String())
	}

	return Book {
		title: title,
		author: authors,
		previewLink: bookThumb,
	}
	

}