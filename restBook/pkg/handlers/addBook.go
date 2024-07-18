package handlers

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"

	"github.com.com/crudBook/pkg/mocks"
	"github.com.com/crudBook/pkg/models"
)

func AddBook(wr http.ResponseWriter, req *http.Request) {
	/*
		1. Read request body
		2. Append item to mock
		3. Send 201 created responce
	*/

	// 1.

	body, err := io.ReadAll(req.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	var book models.Book
	json.Unmarshal(body, &book)

	// 2.

	book.Id = int64(rand.Intn(100))
	mocks.Books = append(mocks.Books, book)

	// 3.

	wr.WriteHeader(http.StatusCreated)
	wr.Header().Add("Content-Type", "application/json")
	json.NewEncoder(wr).Encode("Created")

}
