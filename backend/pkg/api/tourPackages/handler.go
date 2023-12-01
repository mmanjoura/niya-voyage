package tourPackages

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

// FindTourPackages godoc
// @Summary Get all tourPackages with pagination
// @Description Get a list of all tourPackages with optional pagination
// @Tags tourPackages
// @Security ApiKeyAuth
// @Produce json
// @Param offset query int false "Offset for pagination" default(0)
// @Param limit query int false "Limit for pagination" default(10)
// @Success 200 {array} models.TourPackage "Successfully retrieved list of tourPackages"
// @Router /tourPackages [get]
func FindTourPackages(c *gin.Context) {
	var tourPackages []models.TourPackage

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
	// cacheKey := "tourPackages_offset_" + offsetQuery + "_limit_" + limitQuery

	// Try fetching the data from Redis first
	// cachedTourPackages, err := cache.Rdb.Get(cache.Ctx, cacheKey).Result()
	// if err == nil {
	// 	err := json.Unmarshal([]byte(cachedTourPackages), &tourPackages)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal cached data"})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{"data": tourPackages})
	// 	return
	// }

	// If cache missed, fetch data from the database
	database.Database.DB.Offset(offset).Limit(limit).Find(&tourPackages)

	// Serialize tourPackages object and store it in Redis
	//serializedTourPackages, err := json.Marshal(tourPackages)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
		return
	}
	//err = cache.Rdb.Set(cache.Ctx, cacheKey, serializedTourPackages, time.Minute).Err() // Here TTL is set to one hour
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tourPackages})
}

// CreateTourPackage godoc
// @Summary Create a new tourPackage
// @Description Create a new tourPackage with the given input data
// @Tags tourPackages
// @Security ApiKeyAuth
// @Security JwtAuth
// @Accept  json
// @Produce  json
// @Param   input     body   models.CreateTourPackage   true   "Create tourPackage object"
// @Success 201 {object} models.TourPackage "Successfully created tourPackage"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /tourPackages [post]
func CreateTourPackage(c *gin.Context) {
	var input models.CreateTourPackage

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tourPackage := models.TourPackage{Package_Name: input.Package_Name, Description: input.Description, Price: input.Price, Itinerary: input.Itinerary}

	database.Database.DB.Create(&tourPackage)

	// Invalidate cache
	keysPattern := "tourPackages_offset_*"
	keys, err := cache.Rdb.Keys(cache.Ctx, keysPattern).Result()
	if err == nil {
		for _, key := range keys {
			cache.Rdb.Del(cache.Ctx, key)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": tourPackage})
}

// FindTourPackage godoc
// @Summary Find a tourPackage by ID
// @Description Get details of a tourPackage by its ID
// @Tags tourPackages
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "TourPackage ID"
// @Success 200 {object} models.TourPackage "Successfully retrieved tourPackage"
// @Failure 404 {string} string "TourPackage not found"
// @Router /tourPackages/{id} [get]
func FindTourPackage(c *gin.Context) {
	var tourPackage models.TourPackage

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&tourPackage).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tourPackage not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tourPackage})
}

// UpdateTourPackage godoc
// @Summary Update a tourPackage by ID
// @Description Update the tourPackage details for the given ID
// @Tags tourPackages
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "TourPackage ID"
// @Param input body models.UpdateTourPackage true "Update tourPackage object"
// @Success 200 {object} models.TourPackage "Successfully updated tourPackage"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "tourPackage not found"
// @Router /tourPackages/{id} [put]
func UpdateTourPackage(c *gin.Context) {
	var tourPackage models.TourPackage
	var input models.UpdateTourPackage

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&tourPackage).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tourPackage not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.DB.Model(&tourPackage).Updates(models.TourPackage{Package_Name: input.Package_Name, Description: input.Description, Price: input.Price, Itinerary: input.Itinerary})

	c.JSON(http.StatusOK, gin.H{"data": tourPackage})
}

// DeleteTourPackage godoc
// @Summary Delete a tourPackage by ID
// @Description Delete the tourPackage with the given ID
// @Tags tourPackages
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "TourPackage ID"
// @Success 204 {string} string "Successfully deleted tourPackage"
// @Failure 404 {string} string "tourPackage not found"
// @Router /tourPackages/{id} [delete]
func DeleteTourPackage(c *gin.Context) {
	var tourPackage models.TourPackage

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&tourPackage).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tourPackage not found"})
		return
	}

	database.Database.DB.Delete(&tourPackage)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
