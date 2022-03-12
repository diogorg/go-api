package controllers

import (
	adapters "api/adapters/users"
	providers "api/repositories/providers"
	support "api/support"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var repository = providers.GetUserRespository()

func Index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(repository.GetAll())
}

func Show(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	support.PanicError(err)
	user := repository.FindById(id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func Store(w http.ResponseWriter, r *http.Request) {
	user := adapters.Store(r)
	repository.Store(&user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	support.PanicError(err)
	repository.Delete(id)
	w.WriteHeader(http.StatusOK)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	support.PanicError(err)
	user := adapters.Update(r, id)
	repository.Update(&user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
