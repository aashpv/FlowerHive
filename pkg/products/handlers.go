package products

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetupRouter настраивает маршруты и обработчики для микросервиса управления товарами
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Маршрут для получения всех товаров
	r.GET("/products", getAllProductsHandler)

	// Маршрут для получения информации о конкретном товаре
	r.GET("/products/:id", getProductHandler)

	// Маршрут для создания нового товара
	r.POST("/products", createProductHandler)

	// Маршрут для обновления информации о товаре
	r.PUT("/products/:id", updateProductHandler)

	// Маршрут для удаления товара
	r.DELETE("/products/:id", deleteProductHandler)

	return r
}

func getAllProductsHandler(c *gin.Context) {
	// Выполнение запроса к базе данных для получения всех товаров
	var products []Product
	err := DB.Select(&products, "SELECT * FROM products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Отправка товаров в формате JSON в ответ на запрос
	c.JSON(http.StatusOK, products)
}

func getProductHandler(c *gin.Context) {
	// Получение ID товара из параметра запроса
	productID := c.Param("id")

	// Выполнение запроса к базе данных для получения информации о товаре по его ID
	var product Product
	err := DB.Get(&product, "SELECT * FROM products WHERE id = $1", productID)
	if err != nil {
		// Если произошла ошибка, отправляем ответ с кодом ошибки и сообщением об ошибке
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Отправка информации о товаре в формате JSON в ответ на запрос
	c.JSON(http.StatusOK, product)
}

func createProductHandler(c *gin.Context) {
	// Создание переменной для хранения данных о товаре
	var newProduct Product

	// Попытка привязать данные из тела запроса к структуре нового товара
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		// Если не удалось привязать данные, отправляем ответ с кодом ошибки и сообщением об ошибке
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Выполнение запроса к базе данных для добавления нового товара
	_, err := DB.Exec("INSERT INTO products (name, description, price, quantity, image_url) VALUES ($1, $2, $3, $4, $5)",
		newProduct.Name, newProduct.Description, newProduct.Price, newProduct.Quantity, newProduct.ImageURL)
	if err != nil {
		// Если произошла ошибка при выполнении запроса, отправляем ответ с кодом ошибки и сообщением об ошибке
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Отправка ответа об успешном создании товара
	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}

func updateProductHandler(c *gin.Context) {
	// Получение ID товара из параметра запроса
	productID := c.Param("id")

	// Создание переменной для хранения данных об обновляемом товаре
	var updatedProduct Product

	// Попытка привязать данные из тела запроса к структуре обновляемого товара
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		// Если не удалось привязать данные, отправляем ответ с кодом ошибки и сообщением об ошибке
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Выполнение запроса к базе данных для обновления информации о товаре
	_, err := DB.Exec("UPDATE products SET name = $1, description = $2, price = $3, quantity = $4, image_url = $5 WHERE id = $6",
		updatedProduct.Name, updatedProduct.Description, updatedProduct.Price, updatedProduct.Quantity, updatedProduct.ImageURL, productID)
	if err != nil {
		// Если произошла ошибка при выполнении запроса, отправляем ответ с кодом ошибки и сообщением об ошибке
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Отправка ответа об успешном обновлении информации о товаре
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func deleteProductHandler(c *gin.Context) {
	// Получение ID товара из параметра запроса
	productID := c.Param("id")

	// Выполнение запроса к базе данных для удаления товара по его ID
	_, err := DB.Exec("DELETE FROM products WHERE id = $1", productID)
	if err != nil {
		// Если произошла ошибка при выполнении запроса, отправляем ответ с кодом ошибки и сообщением об ошибке
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Отправка ответа об успешном удалении товара
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
