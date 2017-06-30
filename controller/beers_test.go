package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"encoding/json"

	"bytes"

	"github.com/caarlos0/go-web-api-example/model"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBeersIndexSuccess(t *testing.T) {
	var assert = assert.New(t)
	var mockedBeers = []model.Beer{
		{
			ID:        1,
			Name:      "Opa IPA",
			Price:     15.1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	var ds = new(mockBeersDatastore)
	ds.On("AllBeers").Return(mockedBeers, nil)
	var ts = httptest.NewServer(BeersIndex(ds))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	var beers []model.Beer
	json.NewDecoder(res.Body).Decode(&beers)
	defer res.Body.Close()
	assert.Equal(mockedBeers, beers)
}

func TestBeersIndexDatabaseError(t *testing.T) {
	var assert = assert.New(t)
	var ds = new(mockBeersDatastore)
	ds.On("AllBeers").Return([]model.Beer{}, errors.New("random failure"))
	var ts = httptest.NewServer(BeersIndex(ds))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusInternalServerError, res.StatusCode)
}

func TestCreateBeer(t *testing.T) {
	var assert = assert.New(t)
	var beer = model.Beer{
		Name:  "brahma",
		Price: 2.5,
	}
	var ds = new(mockBeersDatastore)
	ds.On("CreateBeer", beer).Return(nil)
	var ts = httptest.NewServer(CreateBeer(ds))
	defer ts.Close()
	body, _ := json.Marshal(&beer)
	res, err := http.Post(ts.URL, "application/json", bytes.NewReader(body))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)
}

func TestCreateBeerDSError(t *testing.T) {
	var assert = assert.New(t)
	var beer = model.Beer{
		Name:  "skoll",
		Price: 2.1,
	}
	var ds = new(mockBeersDatastore)
	ds.On("CreateBeer", beer).Return(errors.New("error"))
	var ts = httptest.NewServer(CreateBeer(ds))
	defer ts.Close()
	body, _ := json.Marshal(&beer)
	res, err := http.Post(ts.URL, "application/json", bytes.NewReader(body))
	assert.NoError(err)
	assert.Equal(http.StatusBadRequest, res.StatusCode)
}

func TestGetBeer(t *testing.T) {
	var assert = assert.New(t)
	var beer = model.Beer{
		ID:        1,
		Name:      "Coruja",
		Price:     52.3,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	var ds = new(mockBeersDatastore)
	ds.On("GetBeer", int64(1)).Return(beer, nil)

	var mux = mux.NewRouter()
	mux.HandleFunc("/{id}", GetBeer(ds))
	var ts = httptest.NewServer(mux)
	defer ts.Close()

	var respBeer model.Beer
	res, err := http.Get(ts.URL + "/1")
	assert.NoError(err)
	assert.NoError(json.NewDecoder(res.Body).Decode(&respBeer))
	assert.Equal(beer, respBeer)
	assert.Equal(http.StatusOK, res.StatusCode)
}

func TestGetBeerDSError(t *testing.T) {
	var assert = assert.New(t)
	var ds = new(mockBeersDatastore)
	ds.On("GetBeer", int64(1)).Return(model.Beer{}, errors.New("fake err"))

	var mux = mux.NewRouter()
	mux.HandleFunc("/{id}", GetBeer(ds))
	var ts = httptest.NewServer(mux)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/1")
	assert.NoError(err)
	assert.Equal(http.StatusBadRequest, res.StatusCode)
}

func TestGetBeerInvalidURL(t *testing.T) {
	var assert = assert.New(t)
	var ds = new(mockBeersDatastore)

	var mux = mux.NewRouter()
	mux.HandleFunc("/{id}", GetBeer(ds))
	var ts = httptest.NewServer(mux)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/asdasd")
	assert.NoError(err)
	assert.Equal(http.StatusBadRequest, res.StatusCode)
}

func TestDeleteBeer(t *testing.T) {
	var assert = assert.New(t)
	var ds = new(mockBeersDatastore)
	ds.On("DeleteBeer", int64(1)).Return(nil)

	var mux = mux.NewRouter()
	mux.HandleFunc("/{id}", DeleteBeer(ds))
	var ts = httptest.NewServer(mux)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/1")
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
}

func TestDeleteBeerDSError(t *testing.T) {
	var assert = assert.New(t)
	var ds = new(mockBeersDatastore)
	ds.On("DeleteBeer", int64(1)).Return(errors.New("fake"))

	var mux = mux.NewRouter()
	mux.HandleFunc("/{id}", DeleteBeer(ds))
	var ts = httptest.NewServer(mux)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/1")
	assert.NoError(err)
	assert.Equal(http.StatusBadRequest, res.StatusCode)
}

type mockBeersDatastore struct {
	mock.Mock
}

func (m mockBeersDatastore) AllBeers() (beers []model.Beer, err error) {
	var args = m.Called()
	return args.Get(0).([]model.Beer), args.Error(1)
}

func (m mockBeersDatastore) GetBeer(id int64) (beer model.Beer, err error) {
	var args = m.Called(id)
	return args.Get(0).(model.Beer), args.Error(1)
}

func (m mockBeersDatastore) CreateBeer(beer model.Beer) (err error) {
	var args = m.Called(beer)
	return args.Error(0)
}

func (m mockBeersDatastore) DeleteBeer(id int64) (err error) {
	var args = m.Called(id)
	return args.Error(0)
}
