package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/caarlos0/go-web-api-example/datastore"
	"github.com/caarlos0/go-web-api-example/model"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func BeersIndex(ds datastore.BeersDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		beers, err := ds.AllBeers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(beers); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func CreateBeer(ds datastore.BeersDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var beer model.Beer
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&beer); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := ds.CreateBeer(beer); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func DeleteBeer(ds datastore.BeersDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := getIdFromPath(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := ds.DeleteBeer(id); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func GetBeer(ds datastore.BeersDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := getIdFromPath(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		beer, err := ds.GetBeer(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := json.NewEncoder(w).Encode(beer); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func getIdFromPath(r *http.Request) (id int64, err error) {
	id, err = strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		return id, errors.Wrap(err, "failed to parse id")
	}
	return
}
