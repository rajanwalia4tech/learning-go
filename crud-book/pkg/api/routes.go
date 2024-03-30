package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/books", GetAllBooks).Methods(http.MethodGet)
	r.HandleFunc("/books", CreateBook).Methods(http.MethodPost)
	r.HandleFunc("/books/{title}", UpdateBook).Methods(http.MethodPut)
	r.HandleFunc("/books/{title}", DeleteBook).Methods(http.MethodDelete)

	return r
}
