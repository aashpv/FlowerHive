package service

import (
	"FlowerHive/pkg/db"
	"FlowerHive/pkg/models"
)

type Service interface {
	Create(body string) (err error)
	Update(body string) (err error)
	GetProduct(id int) (product models.Product, err error)
	GetAllProduct() (products []models.Product, err error)
	Delete(id int) (err error)
}

type service struct {
	pgs db.DataBase
}

func New(postgres db.DataBase) Service {
	return &service{pgs: postgres}
}
