package handlers

import (
	"github.com/jmoiron/sqlx"
	"github.com/nomikz/training/internal/platform/web"
	"log"
	"net/http"
)

func API(logger *log.Logger, db *sqlx.DB) http.Handler {
	app := web.NewApp(logger)

	p := Product{
		DB:  db,
		Log: logger,
	}

	app.Handle(http.MethodGet, "/v1/products", p.List)
	app.Handle(http.MethodPost, "/v1/products", p.Create)
	app.Handle(http.MethodGet, "/v1/products/{id}", p.Retrieve)

	return app
}
