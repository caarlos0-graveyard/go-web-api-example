package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/apex/httplog"
	"github.com/apex/log"
	"github.com/apex/log/handlers/logfmt"
	"github.com/caarlos0/go-web-api-example/config"
	"github.com/caarlos0/go-web-api-example/controller"
	"github.com/caarlos0/go-web-api-example/datastore/database"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log.SetHandler(logfmt.Default)
	log.SetLevel(log.InfoLevel)
	log.Info("starting up...")
	var cfg = config.Get()

	var db = database.Connect(cfg.DatabaseURL)
	defer func() {
		if err := db.Close(); err != nil {
			log.WithError(err).Error("failed to close database connections")
		}
	}()
	var ds = database.New(db)

	var mux = mux.NewRouter()
	mux.Path("/metrics").Handler(promhttp.Handler())
	mux.Path("/status").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	var beersMux = mux.PathPrefix("/beers").Subrouter()
	beersMux.Methods(http.MethodGet).HandlerFunc(controller.BeersIndex(ds))
	beersMux.Methods(http.MethodPost).HandlerFunc(controller.CreateBeer(ds))
	beersMux.Path("{id}").Methods(http.MethodGet).HandlerFunc(controller.GetBeer(ds))
	beersMux.Path("{id}").Methods(http.MethodDelete).HandlerFunc(controller.DeleteBeer(ds))

	var server = &http.Server{
		Handler:      httplog.New(mux),
		Addr:         ":" + cfg.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.WithField("addr", server.Addr).Info("started")
	if err := server.ListenAndServe(); err != nil {
		log.WithError(err).Fatal("failed to start up server")
	}
}
