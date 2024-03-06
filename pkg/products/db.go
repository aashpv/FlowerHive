package products

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// DB представляет подключение к базе данных
var DB *sqlx.DB

// InitializeDB инициализирует подключение к базе данных
func InitializeDB(dataSourceName string) error {
	var err error
	DB, err = sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	return nil
}
