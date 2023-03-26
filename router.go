
package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!\n"))
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Todo Index!\n"))
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	w.Write([]byte("Todo show: " + todoId + "\n"))
}
