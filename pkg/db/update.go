package db

import "FlowerHive/pkg/models"

func (p *postgres) Update(product models.Product) (err error) {
	_, err = p.db.Exec(
		"UPDATE products SET name = $1, description = $2, price = $3, quantity = $4, image_url = $5 WHERE id = $6",
		product.Name,
		product.Description,
		product.Price,
		product.Quantity,
		product.ImageURL,
		product.ID)

	if err != nil {
		return
	}

	return
}
