package database

import (
	"github.com/caarlos0/go-web-api-example/model"
	"github.com/jmoiron/sqlx"
)

type beerstore struct {
	*sqlx.DB
}

func (db *beerstore) AllBeers() (beers []model.Beer, err error) {
	return beers, db.Select(&beers, "SELECT * FROM beers")
}

func (db *beerstore) GetBeer(id int64) (beer model.Beer, err error) {
	return beer, db.Get(&beer, "SELECT * FROM beers WHERE id = $1", id)
}

func (db *beerstore) CreateBeer(beer model.Beer) (err error) {
	_, err = db.Exec(
		"INSERT INTO beers(name, price) VALUES($1, $2)",
		beer.Name,
		beer.Price,
	)
	return
}

func (db *beerstore) DeleteBeer(id int64) (err error) {
	_, err = db.Exec("DELETE FROM beers WHERE id = $1", id)
	return
}
