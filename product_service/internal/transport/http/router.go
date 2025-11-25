package http

import (
	c "product_service/internal/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	
	r := gin.Default()

   
	product := r.Group("/api/v1/products")

	product.GET("/", c.GetProducts)
	product.GET("/:id", c.GetProductByID)
	product.POST("/", c.CreateProduct)
	product.PUT("/:id", c.UpdateProduct)
	product.DELETE("/:id", c.DeleteProduct)

	r.NoRoute(c.NoRoute)

	return r
}