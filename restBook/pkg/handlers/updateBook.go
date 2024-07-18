package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com.com/crudBook/pkg/mocks"
	"github.com.com/crudBook/pkg/models"
	"github.com/gorilla/mux"
)

func UpdateBook(wr http.ResponseWriter, req *http.Request) {

	// Read dynamic parameters
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])

	body, err := io.ReadAll(req.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer req.Body.Close()

	var updateBook models.Book

	json.Unmarshal(body, &updateBook)

	for index, book := range mocks.Books {
		if book.Id == int64(id) {
			book.Title = updateBook.Title
			book.Author = updateBook.Author
			book.Desc = updateBook.Desc
			mocks.Books[index] = book

			wr.WriteHeader(http.StatusOK)
			wr.Header().Add("Content-Type", "application/json")
			json.NewEncoder(wr).Encode("Updated")
			break
		}
	}
}