package payments

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

// FindPayments godoc
// @Summary Get all payments with pagination
// @Description Get a list of all payments with optional pagination
// @Tags payments
// @Security ApiKeyAuth
// @Produce json
// @Param offset query int false "Offset for pagination" default(0)
// @Param limit query int false "Limit for pagination" default(10)
// @Success 200 {array} models.Payment "Successfully retrieved list of payments"
// @Router /payments [get]
func FindPayments(c *gin.Context) {
	var payments []models.Payment

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
	// cacheKey := "payments_offset_" + offsetQuery + "_limit_" + limitQuery

	// Try fetching the data from Redis first
	// cachedPayments, err := cache.Rdb.Get(cache.Ctx, cacheKey).Result()
	// if err == nil {
	// 	err := json.Unmarshal([]byte(cachedPayments), &payments)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal cached data"})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{"data": payments})
	// 	return
	// }

	// If cache missed, fetch data from the database
	database.Database.DB.Offset(offset).Limit(limit).Find(&payments)

	// Serialize payments object and store it in Redis
	//serializedPayments, err := json.Marshal(payments)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
		return
	}
	//err = cache.Rdb.Set(cache.Ctx, cacheKey, serializedPayments, time.Minute).Err() // Here TTL is set to one hour
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": payments})
}

// CreatePayment godoc
// @Summary Create a new payment
// @Description Create a new payment with the given input data
// @Tags payments
// @Security ApiKeyAuth
// @Security JwtAuth
// @Accept  json
// @Produce  json
// @Param   input     body   models.CreatePayment   true   "Create payment object"
// @Success 201 {object} models.Payment "Successfully created payment"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /payments [post]
func CreatePayment(c *gin.Context) {
	var input models.CreatePayment

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment := models.Payment{Amount: input.Amount, Payment_Date: input.Payment_Date, Payment_Method: input.Payment_Method}

	database.Database.DB.Create(&payment)

	// Invalidate cache
	keysPattern := "payments_offset_*"
	keys, err := cache.Rdb.Keys(cache.Ctx, keysPattern).Result()
	if err == nil {
		for _, key := range keys {
			cache.Rdb.Del(cache.Ctx, key)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": payment})
}

// FindPayment godoc
// @Summary Find a payment by ID
// @Description Get details of a payment by its ID
// @Tags payments
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Payment ID"
// @Success 200 {object} models.Payment "Successfully retrieved payment"
// @Failure 404 {string} string "Payment not found"
// @Router /payments/{id} [get]
func FindPayment(c *gin.Context) {
	var payment models.Payment

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&payment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "payment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": payment})
}

// UpdatePayment godoc
// @Summary Update a payment by ID
// @Description Update the payment details for the given ID
// @Tags payments
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "Payment ID"
// @Param input body models.UpdatePayment true "Update payment object"
// @Success 200 {object} models.Payment "Successfully updated payment"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "payment not found"
// @Router /payments/{id} [put]
func UpdatePayment(c *gin.Context) {
	var payment models.Payment
	var input models.UpdatePayment

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&payment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "payment not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.DB.Model(&payment).Updates(models.Payment{Amount: input.Amount, Payment_Date: input.Payment_Date, Payment_Method: input.Payment_Method})

	c.JSON(http.StatusOK, gin.H{"data": payment})
}

// DeletePayment godoc
// @Summary Delete a payment by ID
// @Description Delete the payment with the given ID
// @Tags payments
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Payment ID"
// @Success 204 {string} string "Successfully deleted payment"
// @Failure 404 {string} string "payment not found"
// @Router /payments/{id} [delete]
func DeletePayment(c *gin.Context) {
	var payment models.Payment

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&payment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "payment not found"})
		return
	}

	database.Database.DB.Delete(&payment)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
