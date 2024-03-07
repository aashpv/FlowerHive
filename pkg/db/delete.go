package db

func (p *postgres) Delete(idProduct int) (err error) {
	_, err = p.db.Exec("DELETE FROM products WHERE id = $1", idProduct)
	if err != nil {
		return
	}

	return
}
