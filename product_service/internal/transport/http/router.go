package http

import (
	c "product_service/internal/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	
	r := gin.Default()

	r.GET("/products", c.GetProducts)
    r.POST("/products", c.CreateProduct)
	r.PUT("/products/:id", c.UpdateProduct)
	r.DELETE("/products/:id", c.DeleteProduct)

	return r
}