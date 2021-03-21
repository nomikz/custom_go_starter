package handlers

import (
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/nomikz/training/internal/platform/web"
	"github.com/nomikz/training/internal/product"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"time"
)

type Product struct {
	DB  *sqlx.DB
	Log *log.Logger
}

func (p *Product) Retrieve(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")

	prod, err := product.Retrieve(r.Context(), p.DB, id)
	if err != nil {
		switch err {
		case product.ErrNotFound:
			return web.NewRequestError(err, http.StatusNotFound)
		case product.ErrInvalidId:
			return web.NewRequestError(err, http.StatusBadRequest)
		default:
			return errors.Wrapf(err, "Fail on retrieving product id %q", id)
		}

	}

	if err := web.Respond(w, prod, http.StatusOK); err != nil {
		p.Log.Println(err)
	}

	return nil
}

func (p *Product) List(w http.ResponseWriter, r *http.Request) error {
	list, err := product.List(r.Context(), p.DB)

	if err != nil {
		return errors.Wrap(err, "Error querying db")
	}

	if err := web.Respond(w, list, http.StatusOK); err != nil {
		p.Log.Println("error writing", err)
	}

	return nil
}

func (p *Product) Create(w http.ResponseWriter, r *http.Request) error {
	var nb product.NewProduct
	if err := web.Decode(r, &nb); err != nil {
		return err
	}

	prod, err := product.Create(r.Context(), p.DB, nb, time.Now())
	if err != nil {
		return errors.Wrap(err, "error querying db")
	}

	if err := web.Respond(w, prod, http.StatusCreated); err != nil {
		p.Log.Println("error writing", err)
	}

	return nil
}
