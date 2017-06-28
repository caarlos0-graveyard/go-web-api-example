package datastore

import "github.com/caarlos0/go-web-api-example/model"

type BeersDatastore interface {
	AllBeers() (beers []model.Beer, err error)
	GetBeer(id int64) (beer model.Beer, err error)
	CreateBeer(beer model.Beer) (err error)
	DeleteBeer(id int64) (err error)
}
