package tours

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

// FindTours godoc
// @Summary Get all tours with pagination
// @Description Get a list of all tours with optional pagination
// @Tags tours
// @Security ApiKeyAuth
// @Produce json
// @Param offset query int false "Offset for pagination" default(0)
// @Param limit query int false "Limit for pagination" default(10)
// @Success 200 {array} models.Tour "Successfully retrieved list of tours"
// @Router /tours [get]
func FindTours(c *gin.Context) {
	var tours []models.Tour

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
	// cacheKey := "tours_offset_" + offsetQuery + "_limit_" + limitQuery

	// Try fetching the data from Redis first
	// cachedTours, err := cache.Rdb.Get(cache.Ctx, cacheKey).Result()
	// if err == nil {
	// 	err := json.Unmarshal([]byte(cachedTours), &tours)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal cached data"})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{"data": tours})
	// 	return
	// }

	// If cache missed, fetch data from the database
	database.Database.DB.Offset(offset).Limit(limit).Find(&tours)

	// Serialize tours object and store it in Redis
	//serializedTours, err := json.Marshal(tours)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
		return
	}
	//err = cache.Rdb.Set(cache.Ctx, cacheKey, serializedTours, time.Minute).Err() // Here TTL is set to one hour
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tours})
}

// CreateTour godoc
// @Summary Create a new tour
// @Description Create a new tour with the given input data
// @Tags tours
// @Security ApiKeyAuth
// @Security JwtAuth
// @Accept  json
// @Produce  json
// @Param   input     body   models.CreateTour   true   "Create tour object"
// @Success 201 {object} models.Tour "Successfully created tour"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /tours [post]
func CreateTour(c *gin.Context) {
	var input models.CreateTour

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tour := models.Tour{

		Tag:       input.Tag,
		Title:     input.Title,
		Price:     input.Price,
		Location:  input.Location,
		Duration:  input.Duration,
		Reviews:   input.Reviews,
		TourType:  input.TourType,
		Animation: input.Animation,
	}

	database.Database.DB.Create(&tour)

	// Invalidate cache
	keysPattern := "tours_offset_*"
	keys, err := cache.Rdb.Keys(cache.Ctx, keysPattern).Result()
	if err == nil {
		for _, key := range keys {
			cache.Rdb.Del(cache.Ctx, key)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": tour})
}

// FindTour godoc
// @Summary Find a tour by ID
// @Description Get details of a tour by its ID
// @Tags tours
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Tour ID"
// @Success 200 {object} models.Tour "Successfully retrieved tour"
// @Failure 404 {string} string "Tour not found"
// @Router /tours/{id} [get]
func FindTour(c *gin.Context) {
	var tour models.Tour

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&tour).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tour not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tour})
}

// UpdateTour godoc
// @Summary Update a tour by ID
// @Description Update the tour details for the given ID
// @Tags tours
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "Tour ID"
// @Param input body models.UpdateTour true "Update tour object"
// @Success 200 {object} models.Tour "Successfully updated tour"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "tour not found"
// @Router /tours/{id} [put]
func UpdateTour(c *gin.Context) {
	var tour models.Tour
	var input models.UpdateTour

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&tour).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tour not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.DB.Model(&tour).Updates(models.Tour{
		Tag:       input.Tag,
		Title:     input.Title,
		Price:     input.Price,
		Location:  input.Location,
		Duration:  input.Duration,
		Reviews:   input.Reviews,
		TourType:  input.TourType,
		Animation: input.Animation,
	})

	c.JSON(http.StatusOK, gin.H{"data": tour})
}

// DeleteTour godoc
// @Summary Delete a tour by ID
// @Description Delete the tour with the given ID
// @Tags tours
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Tour ID"
// @Success 204 {string} string "Successfully deleted tour"
// @Failure 404 {string} string "tour not found"
// @Router /tours/{id} [delete]
func DeleteTour(c *gin.Context) {
	var tour models.Tour

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&tour).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tour not found"})
		return
	}

	database.Database.DB.Delete(&tour)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
