package locations

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

// FindLocations godoc
// @Summary Get all Locations with pagination
// @Description Get a list of all Locations with optional pagination
// @Tags Locations
// @Security ApiKeyAuth
// @Produce json
// @Param offset query int false "Offset for pagination" default(0)
// @Param limit query int false "Limit for pagination" default(10)
// @Success 200 {array} models.Location "Successfully retrieved list of Locations"
// @Router /Locations [get]
func FindLocations(c *gin.Context) {
	var Locations []models.Location

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
	cacheKey := "Locations_offset_" + offsetQuery + "_limit_" + limitQuery

	// Try fetching the data from Redis first
	cachedLocations, err := cache.Rdb.Get(cache.Ctx, cacheKey).Result()
	if err == nil {
		err := json.Unmarshal([]byte(cachedLocations), &Locations)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal cached data"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": Locations})
		return
	}

	// If cache missed, fetch data from the database
	database.Database.DB.Offset(offset).Limit(limit).Find(&Locations)

	// Serialize Locations object and store it in Redis
	// serializedLocations, err := json.Marshal(Locations)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
	// 	return
	// }
	// err = cache.Rdb.Set(cache.Ctx, cacheKey, serializedLocations, time.Minute).Err() // Here TTL is set to one hour
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"data": Locations})
}

// CreateLocation godoc
// @Summary Create a new Location
// @Description Create a new Location with the given input data
// @Tags Locations
// @Security ApiKeyAuth
// @Security JwtAuth
// @Accept  json
// @Produce  json
// @Param   input     body   models.CreateLocation   true   "Create Location object"
// @Success 201 {object} models.Location "Successfully created Location"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /Locations [post]
func CreateLocation(c *gin.Context) {
	var input models.CreateLocation

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Location := models.Location{
		MerchantID:   input.MerchantID,
		AddressLine1: input.AddressLine1,
		AddressLine2: input.AddressLine2,
		City:         input.City,
		State:        input.State,
		Country:      input.Country,
		ZipCode:      input.ZipCode,
	}

	database.Database.DB.Create(&Location)

	// Invalidate cache
	keysPattern := "Locations_offset_*"
	keys, err := cache.Rdb.Keys(cache.Ctx, keysPattern).Result()
	if err == nil {
		for _, key := range keys {
			cache.Rdb.Del(cache.Ctx, key)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": Location})
}

// FindLocation godoc
// @Summary Find a Location by ID
// @Description Get details of a Location by its ID
// @Tags Locations
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Location ID"
// @Success 200 {object} models.Location "Successfully retrieved Location"
// @Failure 404 {string} string "Location not found"
// @Router /Locations/{id} [get]
func FindLocation(c *gin.Context) {
	var Location models.Location

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Location).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Location})
}

// UpdateLocation godoc
// @Summary Update a Location by ID
// @Description Update the Location details for the given ID
// @Tags Locations
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "Location ID"
// @Param input body models.UpdateLocation true "Update Location object"
// @Success 200 {object} models.Location "Successfully updated Location"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Location not found"
// @Router /Locations/{id} [put]
func UpdateLocation(c *gin.Context) {
	var Location models.Location
	var input models.UpdateLocation

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Location).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.DB.Model(&Location).Updates(models.Location{
		MerchantID:   input.MerchantID,
		AddressLine1: input.AddressLine1,
		AddressLine2: input.AddressLine2,
		City:         input.City,
		State:        input.State,
		Country:      input.Country,
		ZipCode:      input.ZipCode,
	})

	c.JSON(http.StatusOK, gin.H{"data": Location})
}

// DeleteLocation godoc
// @Summary Delete a Location by ID
// @Description Delete the Location with the given ID
// @Tags Locations
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Location ID"
// @Success 204 {string} string "Successfully deleted Location"
// @Failure 404 {string} string "Location not found"
// @Router /Locations/{id} [delete]
func DeleteLocation(c *gin.Context) {
	var Location models.Location

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Location).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
		return
	}

	database.Database.DB.Delete(&Location)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
