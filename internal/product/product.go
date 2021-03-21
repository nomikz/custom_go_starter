package product

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrNotFound  = errors.New("product not found")
	ErrInvalidId = errors.New("id provided was not a valid UUI")
)

func List(ctx context.Context, db *sqlx.DB) ([]Product, error) {
	list := []Product{}

	const q = `select product_id, name, cost, quantity, date_updated, date_created  from products`
	if err := db.SelectContext(ctx, &list, q); err != nil {
		return nil, err
	}

	return list, nil
}

func Retrieve(ctx context.Context, db *sqlx.DB, id string) (*Product, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, ErrInvalidId
	}

	var p Product

	const q = `select product_id, name, cost, quantity, date_updated, date_created  from products where product_id = $1`
	if err := db.GetContext(ctx, &p, q, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &p, nil
}

func Create(ctx context.Context, db *sqlx.DB, np NewProduct, time time.Time) (*Product, error) {
	p := Product{
		ProductID:   uuid.New().String(),
		Name:        np.Name,
		Cost:        np.Cost,
		Quantity:    np.Quantity,
		DateCreated: time,
		DateUpdated: time,
	}

	const q = `insert into products 
(product_id, name, cost, quantity, date_updated, date_created)  
values ($1, $2, $3, $4, $5, $6) 
`
	if _, err := db.ExecContext(ctx, q, p.ProductID, p.Name, p.Cost, p.Quantity, p.DateUpdated, p.DateCreated); err != nil {
		return nil, errors.Wrapf(err, "inserting products %v", np)
	}

	return &p, nil
}
