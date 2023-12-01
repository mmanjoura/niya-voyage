package reviewRatings

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

// FindReviewRatings godoc
// @Summary Get all reviewRatings with pagination
// @Description Get a list of all reviewRatings with optional pagination
// @Tags reviewRatings
// @Security ApiKeyAuth
// @Produce json
// @Param offset query int false "Offset for pagination" default(0)
// @Param limit query int false "Limit for pagination" default(10)
// @Success 200 {array} models.ReviewRating "Successfully retrieved list of reviewRatings"
// @Router /reviewRatings [get]
func FindReviewRatings(c *gin.Context) {
	var reviewRatings []models.ReviewRating

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
	// cacheKey := "reviewRatings_offset_" + offsetQuery + "_limit_" + limitQuery

	// Try fetching the data from Redis first
	// cachedReviewRatings, err := cache.Rdb.Get(cache.Ctx, cacheKey).Result()
	// if err == nil {
	// 	err := json.Unmarshal([]byte(cachedReviewRatings), &reviewRatings)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal cached data"})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{"data": reviewRatings})
	// 	return
	// }

	// If cache missed, fetch data from the database
	database.Database.DB.Offset(offset).Limit(limit).Find(&reviewRatings)

	// Serialize reviewRatings object and store it in Redis
	//serializedReviewRatings, err := json.Marshal(reviewRatings)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
		return
	}
	//err = cache.Rdb.Set(cache.Ctx, cacheKey, serializedReviewRatings, time.Minute).Err() // Here TTL is set to one hour
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reviewRatings})
}

// CreateReviewRating godoc
// @Summary Create a new reviewRating
// @Description Create a new reviewRating with the given input data
// @Tags reviewRatings
// @Security ApiKeyAuth
// @Security JwtAuth
// @Accept  json
// @Produce  json
// @Param   input     body   models.CreateReviewRating   true   "Create reviewRating object"
// @Success 201 {object} models.ReviewRating "Successfully created reviewRating"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /reviewRatings [post]
func CreateReviewRating(c *gin.Context) {
	var input models.CreateReviewRating

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reviewRating := models.ReviewRating{Rating: input.Rating, Review_Text: input.Review_Text}

	database.Database.DB.Create(&reviewRating)

	// Invalidate cache
	keysPattern := "reviewRatings_offset_*"
	keys, err := cache.Rdb.Keys(cache.Ctx, keysPattern).Result()
	if err == nil {
		for _, key := range keys {
			cache.Rdb.Del(cache.Ctx, key)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": reviewRating})
}

// FindReviewRating godoc
// @Summary Find a reviewRating by ID
// @Description Get details of a reviewRating by its ID
// @Tags reviewRatings
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "ReviewRating ID"
// @Success 200 {object} models.ReviewRating "Successfully retrieved reviewRating"
// @Failure 404 {string} string "ReviewRating not found"
// @Router /reviewRatings/{id} [get]
func FindReviewRating(c *gin.Context) {
	var reviewRating models.ReviewRating

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&reviewRating).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "reviewRating not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reviewRating})
}

// UpdateReviewRating godoc
// @Summary Update a reviewRating by ID
// @Description Update the reviewRating details for the given ID
// @Tags reviewRatings
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "ReviewRating ID"
// @Param input body models.UpdateReviewRating true "Update reviewRating object"
// @Success 200 {object} models.ReviewRating "Successfully updated reviewRating"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "reviewRating not found"
// @Router /reviewRatings/{id} [put]
func UpdateReviewRating(c *gin.Context) {
	var reviewRating models.ReviewRating
	var input models.UpdateReviewRating

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&reviewRating).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "reviewRating not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.DB.Model(&reviewRating).Updates(models.ReviewRating{Rating: input.Rating, Review_Text: input.Review_Text})

	c.JSON(http.StatusOK, gin.H{"data": reviewRating})
}

// DeleteReviewRating godoc
// @Summary Delete a reviewRating by ID
// @Description Delete the reviewRating with the given ID
// @Tags reviewRatings
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "ReviewRating ID"
// @Success 204 {string} string "Successfully deleted reviewRating"
// @Failure 404 {string} string "reviewRating not found"
// @Router /reviewRatings/{id} [delete]
func DeleteReviewRating(c *gin.Context) {
	var reviewRating models.ReviewRating

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&reviewRating).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "reviewRating not found"})
		return
	}

	database.Database.DB.Delete(&reviewRating)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
