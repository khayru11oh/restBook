package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com.com/crudBook/pkg/mocks"
	"github.com/gorilla/mux"
)

func BookByID(wr http.ResponseWriter, req *http.Request) {
	// Read ID parameter
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	
	//Find it from booksData
	for _, book := range mocks.Books {
		if book.Id == int64(id) {
			wr.WriteHeader(http.StatusOK)
			wr.Header().Add("Content-Type", "application/json")
			json.NewEncoder(wr).Encode(book)
			break
		}
	}
}