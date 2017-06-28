package database

import (
	"github.com/caarlos0/go-web-api-example/model"
	"github.com/jmoiron/sqlx"
)

type beerstore struct {
	*sqlx.DB
}

func (db beerstore) AllBeers() (beers []model.Beer, err error) {
	return beers, db.Select(&beers, "SELECT * from beers")
}
func (db beerstore) BeerByID(id int64) (beer model.Beer, err error) {
	return
}
func (db beerstore) CreateBeer(beer model.Beer) (err error) {
	return
}
func (db beerstore) DeleteBeer(beer model.Beer) (err error) {
	return
}
