package controller

import (
	"fmt"
	"net/http"


)

func BeersIndex(ds datastore.Datastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ds.
	}
}
