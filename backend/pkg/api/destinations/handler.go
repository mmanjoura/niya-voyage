package destinations

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

// FindDestinations godoc
// @Summary Get all destinations with pagination
// @Description Get a list of all destinations with optional pagination
// @Tags destinations
// @Security ApiKeyAuth
// @Produce json
// @Param offset query int false "Offset for pagination" default(0)
// @Param limit query int false "Limit for pagination" default(10)
// @Success 200 {array} models.Destination "Successfully retrieved list of destinations"
// @Router /destinations [get]
func FindDestinations(c *gin.Context) {
	var destinations []models.Destination

	// Get query params
	offsetQuery := c.DefaultQuery("offset", "0")
	limitQuery := c.DefaultQuery("limit", "3")

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
	// cacheKey := "destinations_offset_" + offsetQuery + "_limit_" + limitQuery

	// Try fetching the data from Redis first
	// cachedBooks, err := cache.Rdb.Get(cache.Ctx, cacheKey).Result()
	// if err == nil {
	// 	err := json.Unmarshal([]byte(cachedBooks), &destinations)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal cached data"})
	// 		return
	// 	}cd
	// 	c.JSON(http.StatusOK, gin.H{"data": destinations})
	// 	return
	// }

	// If cache missed, fetch data from the database
	database.Database.DB.Offset(offset).Limit(limit).Find(&destinations)

	// Serialize destinations object and store it in Redis
	//serializedBooks, err := json.Marshal(destinations)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
		return
	}
	//err = cache.Rdb.Set(cache.Ctx, cacheKey, serializedBooks, time.Minute).Err() // Here TTL is set to one hour
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": destinations})
}

// CreateDestinationgodoc
// @Summary Create a new Destination
// @Description Create a new Destination with the given input data
// @Tags destinations
// @Security ApiKeyAuth
// @Security JwtAuth
// @Accept  json
// @Produce  json
// @Param   input     body   models.CreateBook   true   "Create Destination object"
// @Success 201 {object} models.Destination "Successfully created Destination"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /destinations [post]
func CreateDestination(c *gin.Context) {
	var input models.CreateDestination

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Destination := models.Destination{Class: input.Class,
		Title:      input.Title,
		Location:   input.Location,
		Travellers: input.Travellers,
		Hover:      input.Hover,
		Img:        input.Img,
		City:       input.City,
		Properties: input.Properties,
		Region:     input.Region,
		Animation:  input.Animation,
		Name:       input.Name}

	database.Database.DB.Create(&Destination)

	// Invalidate cache
	keysPattern := "destinations_offset_*"
	keys, err := cache.Rdb.Keys(cache.Ctx, keysPattern).Result()
	if err == nil {
		for _, key := range keys {
			cache.Rdb.Del(cache.Ctx, key)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": Destination})
}

// FindDestination godoc
// @Summary Find a Destination by ID
// @Description Get details of a Destination by its ID
// @Tags destinations
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Destination ID"
// @Success 200 {object} models.Destination "Successfully retrieved Destination"
// @Failure 404 {string} string "Destination not found"
// @Router /destinations/{id} [get]
func FindDestination(c *gin.Context) {
	var Destination models.Destination

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Destination).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Destination not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Destination})
}

// UpdateDestination godoc
// @Summary Update a Destination by ID
// @Description Update the Destination details for the given ID
// @Tags destinations
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "Destination ID"
// @Param input body models.UpdateBook true "Update Destination object"
// @Success 200 {object} models.Destination "Successfully updated Destination"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Destination not found"
// @Router /destinations/{id} [put]
func UpdateDestination(c *gin.Context) {
	var Destination models.Destination
	var input models.UpdateDestination

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Destination).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Destination not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.DB.Model(&Destination).Updates(models.Destination{Class: input.Class,
		Title:      input.Title,
		Location:   input.Location,
		Travellers: input.Travellers,
		Hover:      input.Hover,
		Img:        input.Img,
		City:       input.City,
		Properties: input.Properties,
		Region:     input.Region,
		Animation:  input.Animation,
		Name:       input.Name})

	c.JSON(http.StatusOK, gin.H{"data": Destination})
}

// DeleteDestination godoc
// @Summary Delete a Destination by ID
// @Description Delete the Destination with the given ID
// @Tags destinations
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Destination ID"
// @Success 204 {string} string "Successfully deleted Destination"
// @Failure 404 {string} string "Destination not found"
// @Router /destinations/{id} [delete]
func DeleteDestination(c *gin.Context) {
	var Destination models.Destination

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Destination).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Destination not found"})
		return
	}

	database.Database.DB.Delete(&Destination)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
