package products

import (
	"net/http"
	"niya-voyage/backend/pkg/cache"
	"niya-voyage/backend/pkg/database"
	"niya-voyage/backend/pkg/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// Healthcheck godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router / [get]
func Healthcheck(g *gin.Context) {
	g.JSON(http.StatusOK, "ok")
}

// FindProducts godoc
// @Summary Get all products with pagination
// @Description Get a list of all products with optional pagination
// @Tags products
// @Security ApiKeyAuth
// @Produce json
// @Param offset query int false "Offset for pagination" default(0)
// @Param limit query int false "Limit for pagination" default(10)
// @Success 200 {array} models.Product "Successfully retrieved list of products"
// @Router /products [get]
func FindProducts(c *gin.Context) {
	var products []models.Product

	// Get query params
	offsetQuery := c.DefaultQuery("offset", "0")
	limitQuery := c.DefaultQuery("limit", "10")

	// Convert query params to integers
	offset, err := strconv.Atoi(offsetQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset format"})
		return
	}

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit format"})
		return
	}

	// Create a cache key based on query params
	// cacheKey := "products_offset_" + offsetQuery + "_limit_" + limitQuery

	// Try fetching the data from Redis first
	// cachedProducts, err := cache.Rdb.Get(cache.Ctx, cacheKey).Result()
	// if err == nil {
	// 	err := json.Unmarshal([]byte(cachedProducts), &products)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal cached data"})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{"data": products})
	// 	return
	// }

	// If cache missed, fetch data from the database
	database.Database.DB.Offset(offset).Limit(limit).Find(&products)

	// Serialize products object and store it in Redis
	//serializedProducts, err := json.Marshal(products)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
		return
	}
	// err = cache.Rdb.Set(cache.Ctx, cacheKey, serializedProducts, time.Minute).Err() // Here TTL is set to one hour
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with the given input data
// @Tags products
// @Security ApiKeyAuth
// @Security JwtAuth
// @Accept  json
// @Produce  json
// @Param   input     body   models.CreateProduct   true   "Create product object"
// @Success 201 {object} models.Product "Successfully created product"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	var input models.CreateProduct

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{Product_Name: input.Product_Name, Description: input.Description, Price: input.Price, Category: input.Category}

	database.Database.DB.Create(&product)

	// Invalidate cache
	keysPattern := "products_offset_*"
	keys, err := cache.Rdb.Keys(cache.Ctx, keysPattern).Result()
	if err == nil {
		for _, key := range keys {
			cache.Rdb.Del(cache.Ctx, key)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": product})
}

// FindProduct godoc
// @Summary Find a product by ID
// @Description Get details of a product by its ID
// @Tags products
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.Product "Successfully retrieved product"
// @Failure 404 {string} string "Product not found"
// @Router /products/{id} [get]
func FindProduct(c *gin.Context) {
	var product models.Product

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// UpdateProduct godoc
// @Summary Update a product by ID
// @Description Update the product details for the given ID
// @Tags products
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID"
// @Param input body models.UpdateProduct true "Update product object"
// @Success 200 {object} models.Product "Successfully updated product"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "product not found"
// @Router /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	var product models.Product
	var input models.UpdateProduct

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.DB.Model(&product).Updates(models.Product{Product_Name: input.Product_Name, Description: input.Description, Price: input.Price, Category: input.Category})

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// DeleteProduct godoc
// @Summary Delete a product by ID
// @Description Delete the product with the given ID
// @Tags products
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Product ID"
// @Success 204 {string} string "Successfully deleted product"
// @Failure 404 {string} string "product not found"
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	var product models.Product

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	database.Database.DB.Delete(&product)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
