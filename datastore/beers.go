package datastore

import "github.com/caarlos0/go-web-api-example/model"

type BeersDatastore interface {
	AllBeers() (beers []model.Beer, err error)
	BeerByID(id int64) (beer model.Beer, err error)
	CreateBeer(beer model.Beer) (err error)
	DeleteBeer(beer model.Beer) (err error)
}
