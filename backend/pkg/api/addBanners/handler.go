package addBanners

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

// FindAddBanners godoc
// @Summary Get all addBanners with pagination
// @Description Get a list of all addBanners with optional pagination
// @Tags addBanners
// @Security ApiKeyAuth
// @Produce json
// @Param offset query int false "Offset for pagination" default(0)
// @Param limit query int false "Limit for pagination" default(10)
// @Success 200 {array} models.AddBanner "Successfully retrieved list of addBanners"
// @Router /addBanners [get]
func FindAddBanners(c *gin.Context) {
	var addBanners []models.AddBanner

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
	// cacheKey := "addBanners_offset_" + offsetQuery + "_limit_" + limitQuery

	// Try fetching the data from Redis first
	// cachedAddBanners, err := cache.Rdb.Get(cache.Ctx, cacheKey).Result()
	// if err == nil {
	// 	err := json.Unmarshal([]byte(cachedAddBanners), &addBanners)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal cached data"})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{"data": addBanners})
	// 	return
	// }

	// If cache missed, fetch data from the database
	database.Database.DB.Offset(offset).Limit(limit).Find(&addBanners)

	// Serialize addBanners object and store it in Redis
	//serializedAddBanners, err := json.Marshal(addBanners)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
		return
	}
	//err = cache.Rdb.Set(cache.Ctx, cacheKey, serializedAddBanners, time.Minute).Err() // Here TTL is set to one hour
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": addBanners})
}

// CreateAddBanner godoc
// @Summary Create a new addBanner
// @Description Create a new addBanner with the given input data
// @Tags addBanners
// @Security ApiKeyAuth
// @Security JwtAuth
// @Accept  json
// @Produce  json
// @Param   input     body   models.CreateAddBanner   true   "Create addBanner object"
// @Success 201 {object} models.AddBanner "Successfully created addBanner"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /addBanners [post]
func CreateAddBanner(c *gin.Context) {
	var input models.CreateAddBanner

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addBanner := models.AddBanner{

		Img:            input.Img,
		Title:          input.Title,
		Meta:           input.Meta,
		RouterPath:     input.RouterPath,
		DelayAnimation: input.DelayAnimation,
	}

	database.Database.DB.Create(&addBanner)

	// Invalidate cache
	keysPattern := "addBanners_offset_*"
	keys, err := cache.Rdb.Keys(cache.Ctx, keysPattern).Result()
	if err == nil {
		for _, key := range keys {
			cache.Rdb.Del(cache.Ctx, key)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": addBanner})
}

// FindAddBanner godoc
// @Summary Find a addBanner by ID
// @Description Get details of a addBanner by its ID
// @Tags addBanners
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "AddBanner ID"
// @Success 200 {object} models.AddBanner "Successfully retrieved addBanner"
// @Failure 404 {string} string "AddBanner not found"
// @Router /addBanners/{id} [get]
func FindAddBanner(c *gin.Context) {
	var addBanner models.AddBanner

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&addBanner).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "addBanner not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": addBanner})
}

// UpdateAddBanner godoc
// @Summary Update a addBanner by ID
// @Description Update the addBanner details for the given ID
// @Tags addBanners
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "AddBanner ID"
// @Param input body models.UpdateAddBanner true "Update addBanner object"
// @Success 200 {object} models.AddBanner "Successfully updated addBanner"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "addBanner not found"
// @Router /addBanners/{id} [put]
func UpdateAddBanner(c *gin.Context) {
	var addBanner models.AddBanner
	var input models.UpdateAddBanner

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&addBanner).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "addBanner not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.DB.Model(&addBanner).Updates(models.AddBanner{
		Img:            input.Img,
		Title:          input.Title,
		Meta:           input.Meta,
		RouterPath:     input.RouterPath,
		DelayAnimation: input.DelayAnimation,
	})

	c.JSON(http.StatusOK, gin.H{"data": addBanner})
}

// DeleteAddBanner godoc
// @Summary Delete a addBanner by ID
// @Description Delete the addBanner with the given ID
// @Tags addBanners
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "AddBanner ID"
// @Success 204 {string} string "Successfully deleted addBanner"
// @Failure 404 {string} string "addBanner not found"
// @Router /addBanners/{id} [delete]
func DeleteAddBanner(c *gin.Context) {
	var addBanner models.AddBanner

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&addBanner).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "addBanner not found"})
		return
	}

	database.Database.DB.Delete(&addBanner)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
