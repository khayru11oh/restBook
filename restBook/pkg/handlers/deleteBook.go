package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com.com/crudBook/pkg/mocks"
	"github.com/gorilla/mux"
)

func DeleteBook(wr http.ResponseWriter, req *http.Request) {
	// Read dynamic parameter

	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])

	mocks.Books = append(mocks.Books[:id-1], mocks.Books[id:]...)
	wr.WriteHeader(http.StatusOK)
	json.NewEncoder(wr).Encode("Deleted")

	// for _, book := range mocks.Books {
	// 	if book.Id == int64(id) {
	// 		mocks.Books = append(mocks.Books[:id-1], mocks.Books[id:]...)

	// 		wr.WriteHeader(http.StatusOK)
	// 		json.NewEncoder(wr).Encode("Deleted successfully")
	// 		break
	// 	}
	// }
}
