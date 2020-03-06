package main

import (
	"database/sql"
	"errors"
)

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *product) getProduct(db *sql.DB) error {
	return db.QueryRow("SELECT name, price FROM products WHERE id=$1",
		p.ID).Scan(&p.Name, &p.Price)
}

func (p *product) updateProduct(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE products SET name=$1, price=$2 WHERE id=$3",
			p.Name, p.Price, p.ID)

	return err
}

func (p *product) deleteProduct(db *sql.DB) error {
	return errors.New("not implemented")
}

func (p *product) createProduct(db *sql.DB) error {
	return errors.New("not implemented")
}

func (p *product) getProducts(db *sql.DB, start, count int) ([]product, error) {
	return nil, errors.New("not implemented")
}
