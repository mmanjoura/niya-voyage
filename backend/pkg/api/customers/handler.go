package customers

import (
	"encoding/json"
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

// FindCustomers godoc
// @Summary Get all Customers with pagination
// @Description Get a list of all Customers with optional pagination
// @Tags Customers
// @Security ApiKeyAuth
// @Produce json
// @Param offset query int false "Offset for pagination" default(0)
// @Param limit query int false "Limit for pagination" default(10)
// @Success 200 {array} models.Customer "Successfully retrieved list of Customers"
// @Router /Customers [get]
func FindCustomers(c *gin.Context) {
	var Customers []models.Customer

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
	cacheKey := "Customers_offset_" + offsetQuery + "_limit_" + limitQuery

	// Try fetching the data from Redis first
	cachedCustomers, err := cache.Rdb.Get(cache.Ctx, cacheKey).Result()
	if err == nil {
		err := json.Unmarshal([]byte(cachedCustomers), &Customers)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal cached data"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": Customers})
		return
	}

	// If cache missed, fetch data from the database
	database.Database.DB.Offset(offset).Limit(limit).Find(&Customers)

	// Serialize Customers object and store it in Redis
	// serializedCustomers, err := json.Marshal(Customers)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
	// 	return
	// }
	// err = cache.Rdb.Set(cache.Ctx, cacheKey, serializedCustomers, time.Minute).Err() // Here TTL is set to one hour
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"data": Customers})
}

// CreateCustomer godoc
// @Summary Create a new Customer
// @Description Create a new Customer with the given input data
// @Tags Customers
// @Security ApiKeyAuth
// @Security JwtAuth
// @Accept  json
// @Produce  json
// @Param   input     body   models.CreateCustomer   true   "Create Customer object"
// @Success 201 {object} models.Customer "Successfully created Customer"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /Customers [post]
func CreateCustomer(c *gin.Context) {
	var input models.CreateCustomer

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Customer := models.Customer{FirstName: input.FirstName, LastName: input.FirstName, Email: input.Email, Phone: input.Phone, Address: input.Address}

	database.Database.DB.Create(&Customer)

	// Invalidate cache
	keysPattern := "Customers_offset_*"
	keys, err := cache.Rdb.Keys(cache.Ctx, keysPattern).Result()
	if err == nil {
		for _, key := range keys {
			cache.Rdb.Del(cache.Ctx, key)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": Customer})
}

// FindCustomer godoc
// @Summary Find a Customer by ID
// @Description Get details of a Customer by its ID
// @Tags Customers
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Customer ID"
// @Success 200 {object} models.Customer "Successfully retrieved Customer"
// @Failure 404 {string} string "Customer not found"
// @Router /Customers/{id} [get]
func FindCustomer(c *gin.Context) {
	var Customer models.Customer

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Customer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Customer})
}

// UpdateCustomer godoc
// @Summary Update a Customer by ID
// @Description Update the Customer details for the given ID
// @Tags Customers
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "Customer ID"
// @Param input body models.UpdateCustomer true "Update Customer object"
// @Success 200 {object} models.Customer "Successfully updated Customer"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Customer not found"
// @Router /Customers/{id} [put]
func UpdateCustomer(c *gin.Context) {
	var Customer models.Customer
	var input models.UpdateCustomer

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Customer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.DB.Model(&Customer).Updates(models.Customer{FirstName: input.FirstName, LastName: input.FirstName, Email: input.Email, Phone: input.Phone, Address: input.Address})

	c.JSON(http.StatusOK, gin.H{"data": Customer})
}

// DeleteCustomer godoc
// @Summary Delete a Customer by ID
// @Description Delete the Customer with the given ID
// @Tags Customers
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Customer ID"
// @Success 204 {string} string "Successfully deleted Customer"
// @Failure 404 {string} string "Customer not found"
// @Router /Customers/{id} [delete]
func DeleteCustomer(c *gin.Context) {
	var Customer models.Customer

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Customer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	database.Database.DB.Delete(&Customer)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
