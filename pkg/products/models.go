package products

import "time"

// Product представляет информацию о товаре (цветке)
type Product struct {
	ID          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	Quantity    int     `json:"quantity" db:"quantity"`
	CategoryID  int     `json:"id_category" db:"id_category"`
}

type Category struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type User struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Role     string `json:"role" db:"role"`
	Number   string `json:"number" db:"number"`
	Email    string `json:"email" db:"email"`
}

type Order struct {
	ID          int       `json:"id" db:"id"`
	Date        time.Time `json:"date" db:"date"`
	Status      string    `json:"status" db:"status"`
	ProductID   int       `json:"id_product" db:"id_product"`
	UserID      int       `json:"id_user" db:"id_user"`
	Address     string    `json:"address" db:"address"`
	OtherNumber string    `json:"other_number" db:"other_number"`
	OtherName   string    `json:"other_name" db:"other_name"`
}
