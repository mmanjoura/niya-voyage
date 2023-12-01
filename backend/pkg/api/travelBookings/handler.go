package travelBookings

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

// FindTravelBookings godoc
// @Summary Get all travelBookings with pagination
// @Description Get a list of all travelBookings with optional pagination
// @Tags travelBookings
// @Security ApiKeyAuth
// @Produce json
// @Param offset query int false "Offset for pagination" default(0)
// @Param limit query int false "Limit for pagination" default(10)
// @Success 200 {array} models.TravelBooking "Successfully retrieved list of travelBookings"
// @Router /travelBookings [get]
func FindTravelBookings(c *gin.Context) {
	var travelBookings []models.TravelBooking

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
	// cacheKey := "travelBookings_offset_" + offsetQuery + "_limit_" + limitQuery

	// Try fetching the data from Redis first
	// cachedBooks, err := cache.Rdb.Get(cache.Ctx, cacheKey).Result()
	// if err == nil {
	// 	err := json.Unmarshal([]byte(cachedBooks), &travelBookings)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal cached data"})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{"data": travelBookings})
	// 	return
	// }

	// If cache missed, fetch data from the database
	database.Database.DB.Offset(offset).Limit(limit).Find(&travelBookings)

	// Serialize travelBookings object and store it in Redis
	//serializedBooks, err := json.Marshal(travelBookings)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
		return
	}
	//err = cache.Rdb.Set(cache.Ctx, cacheKey, serializedBooks, time.Minute).Err() // Here TTL is set to one hour
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": travelBookings})
}

// CreateTravelBookinggodoc
// @Summary Create a new travelBooking
// @Description Create a new travelBooking with the given input data
// @Tags travelBookings
// @Security ApiKeyAuth
// @Security JwtAuth
// @Accept  json
// @Produce  json
// @Param   input     body   models.CreateBook   true   "Create travelBooking object"
// @Success 201 {object} models.TravelBooking "Successfully created travelBooking"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /travelBookings [post]
func CreateTravelBooking(c *gin.Context) {
	var input models.CreateTravelBooking

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	travelBooking := models.TravelBooking{Travel_Date: input.Travel_Date, Status: input.Status}

	database.Database.DB.Create(&travelBooking)

	// Invalidate cache
	keysPattern := "travelBookings_offset_*"
	keys, err := cache.Rdb.Keys(cache.Ctx, keysPattern).Result()
	if err == nil {
		for _, key := range keys {
			cache.Rdb.Del(cache.Ctx, key)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": travelBooking})
}

// FindTravelBooking godoc
// @Summary Find a travelBooking by ID
// @Description Get details of a travelBooking by its ID
// @Tags travelBookings
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "TravelBooking ID"
// @Success 200 {object} models.TravelBooking "Successfully retrieved travelBooking"
// @Failure 404 {string} string "TravelBooking not found"
// @Router /travelBookings/{id} [get]
func FindTravelBooking(c *gin.Context) {
	var travelBooking models.TravelBooking

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&travelBooking).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "travelBooking not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": travelBooking})
}

// UpdateTravelBooking godoc
// @Summary Update a travelBooking by ID
// @Description Update the travelBooking details for the given ID
// @Tags travelBookings
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "TravelBooking ID"
// @Param input body models.UpdateBook true "Update travelBooking object"
// @Success 200 {object} models.TravelBooking "Successfully updated travelBooking"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "travelBooking not found"
// @Router /travelBookings/{id} [put]
func UpdateTravelBooking(c *gin.Context) {
	var travelBooking models.TravelBooking
	var input models.UpdateTravelBooking

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&travelBooking).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "travelBooking not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.DB.Model(&travelBooking).Updates(models.TravelBooking{Travel_Date: input.Travel_Date, Status: input.Status})

	c.JSON(http.StatusOK, gin.H{"data": travelBooking})
}

// DeleteTravelBooking godoc
// @Summary Delete a travelBooking by ID
// @Description Delete the travelBooking with the given ID
// @Tags travelBookings
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "TravelBooking ID"
// @Success 204 {string} string "Successfully deleted travelBooking"
// @Failure 404 {string} string "travelBooking not found"
// @Router /travelBookings/{id} [delete]
func DeleteTravelBooking(c *gin.Context) {
	var travelBooking models.TravelBooking

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&travelBooking).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "travelBooking not found"})
		return
	}

	database.Database.DB.Delete(&travelBooking)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
