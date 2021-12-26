package mneme

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/tidwall/gjson"
)
const googleApi = "https://www.googleapis.com/books/v1/volumes"

type Book struct {
	title       string
	author      []string
	previewLink string
}

type MnemeApi struct {
	bookTitles []string
	bookApiKey string
}


func New(bookTitles []string, bookApiKey string) MnemeApi {
  return MnemeApi{
	  bookTitles: bookTitles,
	  bookApiKey: bookApiKey,
  }
}

func (mneme MnemeApi) GetBooksMetadata() []Book {
	booksMetadata := make([]Book,0)
	
	for _,book := range mneme.bookTitles {
		booksMetadata = append(booksMetadata, fetchBookMatadata(book,mneme.bookApiKey))
	}
	return booksMetadata
}


// TODO to be moved to the service
func fetchBookMatadata(bookTitle string, apiKey string) Book {
	escapedTitle := url.QueryEscape(bookTitle)

	res, err := http.Get(googleApi + "?q=" + escapedTitle + "&key=" + apiKey + "&projection=lite&maxResults=1")

	if err != nil {
		fmt.Printf("error encountered while talking to Book Api %v", err)
	}
	defer res.Body.Close()

	responseBytes, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		fmt.Printf("Could not extract response body %v", readErr)
	}

	title := gjson.GetBytes(responseBytes, "items.#.volumeInfo.title").String()
	authorResults := gjson.GetBytes(responseBytes, "items.#.volumeInfo.authors|@flatten")
	bookThumb := gjson.GetBytes(responseBytes, "items.#.volumeInfo.imageLinks.smallThumbnail").String()

	authors := make([]string, 0)

	for _, author := range authorResults.Array() {
		authors = append(authors, author.String())
	}

	return Book{
		title:       title,
		author:      authors,
		previewLink: bookThumb,
	}

}
