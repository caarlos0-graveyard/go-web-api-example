package controller

import (
	"encoding/json"
	"net/http"

	"github.com/caarlos0/go-web-api-example/datastore"
	"github.com/caarlos0/go-web-api-example/model"
)

func BeersIndex(ds datastore.Datastore) http.HandlerFunc {
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

func CreateBeer(ds datastore.Datastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var beer model.Beer
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&beer); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := ds.CreateBeer(beer); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}
