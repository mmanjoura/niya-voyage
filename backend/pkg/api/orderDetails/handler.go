package orderDetails

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

// FindOrderDetails godoc
// @Summary Get all orderDetails with pagination
// @Description Get a list of all orderDetails with optional pagination
// @Tags orderDetails
// @Security ApiKeyAuth
// @Produce json
// @Param offset query int false "Offset for pagination" default(0)
// @Param limit query int false "Limit for pagination" default(10)
// @Success 200 {array} models.OrderDetail "Successfully retrieved list of orderDetails"
// @Router /orderDetails [get]
func FindOrderDetails(c *gin.Context) {
	var orderDetails []models.OrderDetail

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
	// cacheKey := "orderDetails_offset_" + offsetQuery + "_limit_" + limitQuery

	// Try fetching the data from Redis first
	// cachedOrderDetails, err := cache.Rdb.Get(cache.Ctx, cacheKey).Result()
	// if err == nil {
	// 	err := json.Unmarshal([]byte(cachedOrderDetails), &orderDetails)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal cached data"})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{"data": orderDetails})
	// 	return
	// }

	// If cache missed, fetch data from the database
	database.Database.DB.Offset(offset).Limit(limit).Find(&orderDetails)

	// Serialize orderDetails object and store it in Redis
	//serializedOrderDetails, err := json.Marshal(orderDetails)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
		return
	}
	//err = cache.Rdb.Set(cache.Ctx, cacheKey, serializedOrderDetails, time.Minute).Err() // Here TTL is set to one hour
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orderDetails})
}

// CreateOrderDetail godoc
// @Summary Create a new orderDetail
// @Description Create a new orderDetail with the given input data
// @Tags orderDetails
// @Security ApiKeyAuth
// @Security JwtAuth
// @Accept  json
// @Produce  json
// @Param   input     body   models.CreateOrderDetail   true   "Create orderDetail object"
// @Success 201 {object} models.OrderDetail "Successfully created orderDetail"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /orderDetails [post]
func CreateOrderDetail(c *gin.Context) {
	var input models.CreateOrderDetail

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderDetail := models.OrderDetail{Quantity: input.Quantity, Subtotal: input.Subtotal}

	database.Database.DB.Create(&orderDetail)

	// Invalidate cache
	keysPattern := "orderDetails_offset_*"
	keys, err := cache.Rdb.Keys(cache.Ctx, keysPattern).Result()
	if err == nil {
		for _, key := range keys {
			cache.Rdb.Del(cache.Ctx, key)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": orderDetail})
}

// FindOrderDetail godoc
// @Summary Find a orderDetail by ID
// @Description Get details of a orderDetail by its ID
// @Tags orderDetails
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "OrderDetail ID"
// @Success 200 {object} models.OrderDetail "Successfully retrieved orderDetail"
// @Failure 404 {string} string "OrderDetail not found"
// @Router /orderDetails/{id} [get]
func FindOrderDetail(c *gin.Context) {
	var orderDetail models.OrderDetail

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&orderDetail).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "orderDetail not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orderDetail})
}

// UpdateOrderDetail godoc
// @Summary Update a orderDetail by ID
// @Description Update the orderDetail details for the given ID
// @Tags orderDetails
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "OrderDetail ID"
// @Param input body models.UpdateOrderDetail true "Update orderDetail object"
// @Success 200 {object} models.OrderDetail "Successfully updated orderDetail"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "orderDetail not found"
// @Router /orderDetails/{id} [put]
func UpdateOrderDetail(c *gin.Context) {
	var orderDetail models.OrderDetail
	var input models.UpdateOrderDetail

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&orderDetail).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "orderDetail not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.DB.Model(&orderDetail).Updates(models.OrderDetail{Quantity: input.Quantity, Subtotal: input.Subtotal})

	c.JSON(http.StatusOK, gin.H{"data": orderDetail})
}

// DeleteOrderDetail godoc
// @Summary Delete a orderDetail by ID
// @Description Delete the orderDetail with the given ID
// @Tags orderDetails
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "OrderDetail ID"
// @Success 204 {string} string "Successfully deleted orderDetail"
// @Failure 404 {string} string "orderDetail not found"
// @Router /orderDetails/{id} [delete]
func DeleteOrderDetail(c *gin.Context) {
	var orderDetail models.OrderDetail

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&orderDetail).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "orderDetail not found"})
		return
	}

	database.Database.DB.Delete(&orderDetail)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
