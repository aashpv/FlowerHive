package db

import (
	"FlowerHive/pkg/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DataBase interface {
	Create(product models.Product) (err error)
	GetAllProducts() (products []models.Product, err error)
	GetProductId(id int) (product models.Product, err error)
	Update(product models.Product) (err error)
	Delete(idProduct int) (err error)
}

type postgres struct {
	db *sqlx.DB
}

func New(dataSourceName string) (p DataBase, err error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &postgres{db: db}, nil
}
