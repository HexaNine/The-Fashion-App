package controller

import (
	"product_service/internal/config/database"
	"product_service/internal/entity"
	"product_service/internal/mapstruct"
	"product_service/internal/repository"
	"product_service/internal/service"
	"product_service/internal/transport/dto"

	"strconv"

	"github.com/gin-gonic/gin"
)

var productService service.ProductServiceInterface = service.NewProductService(
	repository.NewProductRepository(database.GetClient()))


func GetProducts(c *gin.Context) {

	filterQuery := c.QueryMap("filter")

	sortQuery := c.Query("sort")
	
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if err != nil {
		c.JSON(500, dto.Error[string]("Something went worng"))
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

		if err != nil {
		c.JSON(500, dto.Error[string]("Internal Server Error"))
		return
	}

	if page < 1{ page = 1 }

	if limit < 1 { limit = 10 }

	meta, products , err := productService.GetAllProducts(&dto.MetaData{
		Page: page,
		Limit: limit,
	}, &sortQuery, &filterQuery)

	if err != nil {
		c.JSON(500, dto.Error[string]("Failed to fetch products"))
		return
	}

	if len(products) < 1 {
		products = []entity.Product{}
	}

	c.JSON(200, dto.Ok("Product fetch successfully", &products, meta))
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(400, dto.BadRequest[string]("invalid ID"))
		return
	}

	product, err := productService.GetProductByID(id)
	if err != nil {
		c.JSON(500, dto.Error[string]("Failed to fetch product"))
		return
	}

	if product == nil {
		c.JSON(404, dto.NotFound("Product not found"))
		return
	}

	c.JSON(200, dto.Ok("Product fetched successfully", product, nil))
}

func CreateProduct(c *gin.Context) {
	product := entity.Product{}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(400, dto.BadRequest[string]("Failed to bind multipart form"))
		return
	}

	newProduct := make(map[string]interface{})
	for key, values := range form.Value {
		newProduct[key] = values[0]
	}

	err = mapstruct.MapToProduct(newProduct, &product)
	if err != nil {
		c.JSON(400, dto.BadRequest[string]("Invalid product data"))
		return
	}

	err = productService.CreateProduct(&product)
	if err != nil {
		c.JSON(500, dto.Error[string]("Failed to create product"))
		return
	}

	c.JSON(201, dto.Created("Product created successfully", &product))
}

func UpdateProduct(c *gin.Context) {

	product := entity.Product{}

	id := c.Param("id")

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(400, dto.BadRequest[string]("Failed to bind multipart form"))
		return
	}

	newProduct := make(map[string]interface{})
	for key, values := range form.Value {
		newProduct[key] = values[0]
	}

	err = mapstruct.MapToProduct(newProduct, &product)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid product data"})
		return
	}

	err = productService.UpdateProduct(id,&product)
	if err != nil {
		c.JSON(400, dto.NotFound("product not found : "))
		return
	}

	updatedProduct , err := productService.GetProductByID(id)

	if err != nil {
		c.JSON(500, dto.Error[string]("Failed to fetch updated product"))
		return
	}
	
	product = *updatedProduct

	
	c.JSON(200, dto.Ok("Product update successfully", &product, nil))
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(400, dto.BadRequest[string]("Invalid ID"))
		return
	}

	err := productService.DeleteProduct(id)
	
	if err != nil {
		c.JSON(404, dto.NotFound("Product not found"))
		return
	}

	c.JSON(200, dto.Ok[string]("Product deleted successfully", nil, nil))
}

func NoRoute(c *gin.Context) {
	c.JSON(404, dto.NotFound("Route not found"))
}