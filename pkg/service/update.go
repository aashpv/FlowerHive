package service

import (
	"FlowerHive/pkg/models"
	"encoding/json"
)

func (s *service) Update(body string) (err error) {
	var product models.Product

	err = json.Unmarshal([]byte(body), &product)
	if err != nil {
		return
	}

	err = s.pgs.Update(product)
	if err != nil { // if err use log for writing in file
		return
	}

	// если ошибок нет, то тоже надо логировать
	return
}
