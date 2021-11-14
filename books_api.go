package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)


const googleApi = "https://www.googleapis.com/books/v1/volumes"

func fetchBookMatadata(bookTitle string, apiKey string) {
	escapedTitle := url.QueryEscape(bookTitle)
    	
	res, err := http.Get(googleApi+"?q="+escapedTitle+"&key="+apiKey+"&projection=lite&maxResults=1")

	if err!= nil {
		fmt.Printf("error encountered while talking to Book Api %v",err)
	}
	defer res.Body.Close()

	responseBytes,err2 := ioutil.ReadAll(res.Body)

	if err2!=nil {
		fmt.Printf("Could not extract response body %v",err2)
	}

	fmt.Println(string(responseBytes))

}