package product

import "github.com/jmoiron/sqlx"

func List(db *sqlx.DB) ([]Product, error) {
	list := []Product{}

	const q = `select product_id, name, cost, quantity, date_updated, date_created  from products`
	if err := db.Select(&list, q); err != nil {
		return nil, err
	}
	return list, nil
}
