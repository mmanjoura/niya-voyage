package activities

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

// FindActivities godoc
// @Summary Get all activities with pagination
// @Description Get a list of all activities with optional pagination
// @Tags activities
// @Security ApiKeyAuth
// @Produce json
// @Param offset query int false "Offset for pagination" default(0)
// @Param limit query int false "Limit for pagination" default(10)
// @Success 200 {array} models.Activity "Successfully retrieved list of activities"
// @Router /activities [get]
func FindActivities(c *gin.Context) {
	var activities []models.Activity

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
	// cacheKey := "activities_offset_" + offsetQuery + "_limit_" + limitQuery

	// Try fetching the data from Redis first
	// cachedActivities, err := cache.Rdb.Get(cache.Ctx, cacheKey).Result()
	// if err == nil {
	// 	err := json.Unmarshal([]byte(cachedActivities), &activities)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal cached data"})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{"data": activities})
	// 	return
	// }

	// If cache missed, fetch data from the database
	database.Database.DB.Offset(offset).Limit(limit).Find(&activities)

	// Serialize activities object and store it in Redis
	//serializedActivities, err := json.Marshal(activities)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
		return
	}
	//err = cache.Rdb.Set(cache.Ctx, cacheKey, serializedActivities, time.Minute).Err() // Here TTL is set to one hour
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": activities})
}

// CreateActivity godoc
// @Summary Create a new tour
// @Description Create a new tour with the given input data
// @Tags activities
// @Security ApiKeyAuth
// @Security JwtAuth
// @Accept  json
// @Produce  json
// @Param   input     body   models.CreateActivity   true   "Create tour object"
// @Success 201 {object} models.Activity "Successfully created tour"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /activities [post]
func CreateActivity(c *gin.Context) {
	var input models.CreateActivity

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tour := models.Activity{

		Tag:       input.Tag,
		Title:     input.Title,
		Price:     input.Price,
		Location:  input.Location,
		Duration:  input.Duration,
		Reviews:   input.Reviews,
		Ratings:   input.Ratings,
		Animation: input.Animation,
	}

	database.Database.DB.Create(&tour)

	// Invalidate cache
	keysPattern := "activities_offset_*"
	keys, err := cache.Rdb.Keys(cache.Ctx, keysPattern).Result()
	if err == nil {
		for _, key := range keys {
			cache.Rdb.Del(cache.Ctx, key)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": tour})
}

// FindActivity godoc
// @Summary Find a tour by ID
// @Description Get details of a tour by its ID
// @Tags activities
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Activity ID"
// @Success 200 {object} models.Activity "Successfully retrieved tour"
// @Failure 404 {string} string "Activity not found"
// @Router /activities/{id} [get]
func FindActivity(c *gin.Context) {
	var tour models.Activity

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&tour).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tour not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tour})
}

// UpdateActivity godoc
// @Summary Update a tour by ID
// @Description Update the tour details for the given ID
// @Tags activities
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "Activity ID"
// @Param input body models.UpdateActivity true "Update tour object"
// @Success 200 {object} models.Activity "Successfully updated tour"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "tour not found"
// @Router /activities/{id} [put]
func UpdateActivity(c *gin.Context) {
	var tour models.Activity
	var input models.UpdateActivity

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&tour).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tour not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.DB.Model(&tour).Updates(models.Activity{
		Tag:       input.Tag,
		Title:     input.Title,
		Price:     input.Price,
		Location:  input.Location,
		Duration:  input.Duration,
		Reviews:   input.Reviews,
		Ratings:   input.Ratings,
		Animation: input.Animation,
	})

	c.JSON(http.StatusOK, gin.H{"data": tour})
}

// DeleteActivity godoc
// @Summary Delete a tour by ID
// @Description Delete the tour with the given ID
// @Tags activities
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Activity ID"
// @Success 204 {string} string "Successfully deleted tour"
// @Failure 404 {string} string "tour not found"
// @Router /activities/{id} [delete]
func DeleteActivity(c *gin.Context) {
	var tour models.Activity

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&tour).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tour not found"})
		return
	}

	database.Database.DB.Delete(&tour)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
