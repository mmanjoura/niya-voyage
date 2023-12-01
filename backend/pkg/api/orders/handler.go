package orders

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

// FindOrders godoc
// @Summary Get all orders with pagination
// @Description Get a list of all orders with optional pagination
// @Tags orders
// @Security ApiKeyAuth
// @Produce json
// @Param offset query int false "Offset for pagination" default(0)
// @Param limit query int false "Limit for pagination" default(10)
// @Success 200 {array} models.Order "Successfully retrieved list of orders"
// @Router /orders [get]
func FindOrders(c *gin.Context) {
	var orders []models.Order

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
	// cacheKey := "orders_offset_" + offsetQuery + "_limit_" + limitQuery

	// Try fetching the data from Redis first
	// cachedOrders, err := cache.Rdb.Get(cache.Ctx, cacheKey).Result()
	// if err == nil {
	// 	err := json.Unmarshal([]byte(cachedOrders), &orders)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal cached data"})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{"data": orders})
	// 	return
	// }

	// If cache missed, fetch data from the database
	database.Database.DB.Offset(offset).Limit(limit).Find(&orders)

	// Serialize orders object and store it in Redis
	//serializedOrders, err := json.Marshal(orders)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
		return
	}
	//err = cache.Rdb.Set(cache.Ctx, cacheKey, serializedOrders, time.Minute).Err() // Here TTL is set to one hour
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with the given input data
// @Tags orders
// @Security ApiKeyAuth
// @Security JwtAuth
// @Accept  json
// @Produce  json
// @Param   input     body   models.CreateOrder   true   "Create order object"
// @Success 201 {object} models.Order "Successfully created order"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /orders [post]
func CreateOrder(c *gin.Context) {
	var input models.CreateOrder

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := models.Order{Total_Amount: input.Total_Amount, Status: input.Status}

	database.Database.DB.Create(&order)

	// Invalidate cache
	keysPattern := "orders_offset_*"
	keys, err := cache.Rdb.Keys(cache.Ctx, keysPattern).Result()
	if err == nil {
		for _, key := range keys {
			cache.Rdb.Del(cache.Ctx, key)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": order})
}

// FindOrder godoc
// @Summary Find a order by ID
// @Description Get details of a order by its ID
// @Tags orders
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} models.Order "Successfully retrieved order"
// @Failure 404 {string} string "Order not found"
// @Router /orders/{id} [get]
func FindOrder(c *gin.Context) {
	var order models.Order

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// UpdateOrder godoc
// @Summary Update a order by ID
// @Description Update the order details for the given ID
// @Tags orders
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "Order ID"
// @Param input body models.UpdateOrder true "Update order object"
// @Success 200 {object} models.Order "Successfully updated order"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "order not found"
// @Router /orders/{id} [put]
func UpdateOrder(c *gin.Context) {
	var order models.Order
	var input models.UpdateOrder

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.DB.Model(&order).Updates(models.Order{Total_Amount: input.Total_Amount, Status: input.Status})

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// DeleteOrder godoc
// @Summary Delete a order by ID
// @Description Delete the order with the given ID
// @Tags orders
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Order ID"
// @Success 204 {string} string "Successfully deleted order"
// @Failure 404 {string} string "order not found"
// @Router /orders/{id} [delete]
func DeleteOrder(c *gin.Context) {
	var order models.Order

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	database.Database.DB.Delete(&order)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
