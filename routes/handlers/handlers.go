package handlers

import (
	"api/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

var R = createRoute()

func createRoute() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	return r
}

func Get(pattern string, handlers func(w http.ResponseWriter, r *http.Request)) {
	R.HandleFunc(pattern, handlers).Methods("GET")
}

func Post(pattern string, handlers func(w http.ResponseWriter, r *http.Request)) {
	R.HandleFunc(pattern, handlers).Methods("POST")
}

func Delete(pattern string, handlers func(w http.ResponseWriter, r *http.Request)) {
	R.HandleFunc(pattern, handlers).Methods("DELETE")
}

func Patch(pattern string, handlers func(w http.ResponseWriter, r *http.Request)) {
	R.HandleFunc(pattern, handlers).Methods("PATCH")
}

func Put(pattern string, handlers func(w http.ResponseWriter, r *http.Request)) {
	R.HandleFunc(pattern, handlers).Methods("Put")
}
