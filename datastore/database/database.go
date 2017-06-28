package database

import (
	"database/sql"

	"github.com/apex/log"
	"github.com/caarlos0/go-web-api-example/datastore"
	"github.com/jmoiron/sqlx"
)

func Connect(url string) *sql.DB {
	var log = log.WithField("url", url)
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.WithError(err).Fatal("failed to open connection to database")
	}
	if err := db.Ping(); err != nil {
		log.WithError(err).Fatal("failed to ping database")
	}
	return db
}

func New(db *sql.DB) datastore.Datastore {
	var dbx = sqlx.NewDb(db, "postgres")
	return struct {
		*beerstore
	}{
		&beerstore{dbx},
	}
}
