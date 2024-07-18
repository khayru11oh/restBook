package handlers

import (
	"encoding/json"
	"net/http"

	"github.com.com/crudBook/pkg/mocks"
)

func GetBooks(wr http.ResponseWriter, req *http.Request) {

	wr.Header().Add("Content-Type", "application/json")
	wr.WriteHeader(http.StatusOK)
	json.NewEncoder(wr).Encode(mocks.Books)
}
