package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/apex/httplog"
	"github.com/apex/log"
	"github.com/apex/log/handlers/logfmt"
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

	var db = database.Connect("postgres://localhost:5432/beers?sslmode=disable")
	defer func() {
		if err := db.Close(); err != nil {
			log.WithError(err).Error("failed to close connection with database")
		}
	}()
	var ds = database.New(db)

	var mux = mux.NewRouter()
	mux.Path("/metrics").Handler(promhttp.Handler())
	mux.Path("/status").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	mux.Path("/beers").Methods(http.MethodGet).HandlerFunc(controller.BeersIndex(ds))
	mux.Path("/beers").Methods(http.MethodPost).HandlerFunc(controller.CreateBeer(ds))

	var server = &http.Server{
		Handler:      httplog.New(mux),
		Addr:         ":3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.WithField("addr", server.Addr).Info("started")
	if err := server.ListenAndServe(); err != nil {
		log.WithError(err).Fatal("failed to start up server")
	}
}
