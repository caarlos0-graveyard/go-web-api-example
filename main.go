package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/apex/httplog"
	"github.com/apex/log"
	"github.com/apex/log/handlers/logfmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log.SetHandler(logfmt.Default)
	log.SetLevel(log.InfoLevel)
	log.Info("starting up...")

	var mux = mux.NewRouter()
	mux.Path("/metrics").Handler(promhttp.Handler())
	mux.Path("/status").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})
	// var beers = mux.PathPrefix("/beers")

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
