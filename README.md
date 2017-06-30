# go-web-api-example

## Why?

The idea is to provide a sample architecture for web API's written in Go,
using PostgreSQL as database and Prometheus for metrics.

## What it depends on?

- gorilla/mux for better routing
- apex/log and apex/httplog for logging
- caarlos0/env for environment config parsing
- prometheus/client_golang for prometheus metrics
- jmoiron/sqlx for improved sql
- lib/pq as postgresql driver

## How is it organized?

```console
$ tree -d .
.
├── config
├── controller
├── datastore
│   └── database
├── migrations
└── model
```

- `config`: contains the config parsing logic;
- `controller`: contains the controllers, which usually take an instance
of the datasource as parameter;
- `datasource`: contains the interfaces that define the access to the database;
- `datasource/database`: contains the real implementation of the `datasource`
interfaces;
- `migrations`: simple SQL files to setup the database;
- `model`: contains the models, which are usually shared accross several
packages.

## Clonning, installing the dependencies, etc

Golang needs packages to be placed in the right folders inside `$GOPATH`
to work properly:

```console
$ cd $GOPATH
$ mkdir -p src/github.com/caarlos0
$ cd $_
$ git clone git@github.com:caarlos0/go-web-api-example.git
$ cd go-web-api-example
```

To install all dependencies, just run:

```console
$ make setup
```

There are several tasks available, try `make help`

## Running locally

#### 1. Create and migrate the database

```console
$ createdb beers
$ for sql in ./migrations/*; do psql ab -f $sql; done
```

#### 2. Start it up

```console
$ go run main.go
```

It should be alive at port 3000!
