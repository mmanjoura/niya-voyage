package golfs

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

// FindGolfs godoc
// @Summary Get all golfs with pagination
// @Description Get a list of all golfs with optional pagination
// @Tags golfs
// @Security ApiKeyAuth
// @Produce json
// @Param offset query int false "Offset for pagination" default(0)
// @Param limit query int false "Limit for pagination" default(10)
// @Success 200 {array} models.Golf "Successfully retrieved list of golfs"
// @Router /golfs [get]
func FindGolfs(c *gin.Context) {
	var golfs []models.Golf

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
	// cacheKey := "golfs_offset_" + offsetQuery + "_limit_" + limitQuery

	// Try fetching the data from Redis first
	// cachedGolfs, err := cache.Rdb.Get(cache.Ctx, cacheKey).Result()
	// if err == nil {
	// 	err := json.Unmarshal([]byte(cachedGolfs), &golfs)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal cached data"})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{"data": golfs})
	// 	return
	// }

	// If cache missed, fetch data from the database

	database.Database.DB.Offset(offset).Limit(limit).Raw(`SELECT ID,
			tag,
			title,
			price,
			location,
			reviews,
			ratings,
			animation,
			holes,
			duration,
			name,
			Created_At,
			Updated_At
		FROM Golfs`).Scan(&golfs)
	galleryImages := []models.GalleryImage{}
	slideImages := []models.SlideImage{}

	for i, v := range golfs {
		database.Database.DB.Find(&galleryImages, "golf_id = ?", v.ID)
		database.Database.DB.Find(&slideImages, "golf_id = ?", v.ID)
		golfs[i].GalleryImages = galleryImages
		golfs[i].SlideImages = slideImages
		for _, v := range slideImages {
			golfs[i].SlideImg = append(golfs[i].SlideImg, v.Img)
		}

	}

	// Serialize golfs object and store it in Redis
	//serializedGolfs, err := json.Marshal(golfs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
		return
	}
	//err = cache.Rdb.Set(cache.Ctx, cacheKey, serializedGolfs, time.Minute).Err() // Here TTL is set to one hour
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": golfs})
}

// CreateGolf godoc
// @Summary Create a new golf
// @Description Create a new golf with the given input data
// @Tags golfs
// @Security ApiKeyAuth
// @Security JwtAuth
// @Accept  json
// @Produce  json
// @Param   input     body   models.CreateGolf   true   "Create golf object"
// @Success 201 {object} models.Golf "Successfully created golf"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /golfs [post]
func CreateGolf(c *gin.Context) {
	var input models.CreateGolf

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	golf := models.Golf{

		Tag:       input.Tag,
		Title:     input.Title,
		Price:     input.Price,
		Location:  input.Location,
		Duration:  input.Duration,
		Reviews:   input.Reviews,
		Ratings:   input.Ratings,
		Holes:     input.Holes,
		Name:      input.Name,
		Animation: input.Animation,
	}

	database.Database.DB.Create(&golf)

	// Invalidate cache
	keysPattern := "golfs_offset_*"
	keys, err := cache.Rdb.Keys(cache.Ctx, keysPattern).Result()
	if err == nil {
		for _, key := range keys {
			cache.Rdb.Del(cache.Ctx, key)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": golf})
}

// FindGolf godoc
// @Summary Find a golf by ID
// @Description Get details of a golf by its ID
// @Tags golfs
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Golf ID"
// @Success 200 {object} models.Golf "Successfully retrieved golf"
// @Failure 404 {string} string "Golf not found"
// @Router /golfs/{id} [get]
func FindGolf(c *gin.Context) {
	var golf models.Golf

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&golf).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "golf not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": golf})
}

// UpdateGolf godoc
// @Summary Update a golf by ID
// @Description Update the golf details for the given ID
// @Tags golfs
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "Golf ID"
// @Param input body models.UpdateGolf true "Update golf object"
// @Success 200 {object} models.Golf "Successfully updated golf"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "golf not found"
// @Router /golfs/{id} [put]
func UpdateGolf(c *gin.Context) {
	var golf models.Golf
	var input models.UpdateGolf

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&golf).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "golf not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.DB.Model(&golf).Updates(models.Golf{
		Tag:       input.Tag,
		Title:     input.Title,
		Price:     input.Price,
		Location:  input.Location,
		Duration:  input.Duration,
		Reviews:   input.Reviews,
		Ratings:   input.Ratings,
		Holes:     input.Holes,
		Name:      input.Name,
		Animation: input.Animation,
	})

	c.JSON(http.StatusOK, gin.H{"data": golf})
}

// DeleteGolf godoc
// @Summary Delete a golf by ID
// @Description Delete the golf with the given ID
// @Tags golfs
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Golf ID"
// @Success 204 {string} string "Successfully deleted golf"
// @Failure 404 {string} string "golf not found"
// @Router /golfs/{id} [delete]
func DeleteGolf(c *gin.Context) {
	var golf models.Golf

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&golf).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "golf not found"})
		return
	}

	database.Database.DB.Delete(&golf)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
