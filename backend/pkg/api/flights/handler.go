package flights

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

// FindFlights godoc
// @Summary Get all flights with pagination
// @Description Get a list of all flights with optional pagination
// @Tags flights
// @Security ApiKeyAuth
// @Produce json
// @Param offset query int false "Offset for pagination" default(0)
// @Param limit query int false "Limit for pagination" default(10)
// @Success 200 {array} models.Flight "Successfully retrieved list of flights"
// @Router /flights [get]
func FindFlights(c *gin.Context) {
	var flights []models.Flight

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
	// cacheKey := "flights_offset_" + offsetQuery + "_limit_" + limitQuery

	// Try fetching the data from Redis first
	// cachedBooks, err := cache.Rdb.Get(cache.Ctx, cacheKey).Result()
	// if err == nil {
	// 	err := json.Unmarshal([]byte(cachedBooks), &flights)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal cached data"})
	// 		return
	// 	}cd
	// 	c.JSON(http.StatusOK, gin.H{"data": flights})
	// 	return
	// }

	// If cache missed, fetch data from the database
	dB := database.Database.DB
	// dB.Joins("JOIN flight_images ON flight_id = flights.id").Offset(offset).Limit(limit).Find(&flights)
	// dB.Joins("JOIN flight_images").Offset(offset).Limit(limit).Find(&flights)
	// database.Database.DB.Offset(offset).Limit(limit).Find(&flights)
	// Raw SQL
	dB.Raw("Select * from Flight_List fl inner join flights f on fl.flight_id = f.id").Offset(offset).Limit(limit).Scan(&flights)

	// Serialize flights object and store it in Redis
	//serializedBooks, err := json.Marshal(flights)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
		return
	}
	//err = cache.Rdb.Set(cache.Ctx, cacheKey, serializedBooks, time.Minute).Err() // Here TTL is set to one hour
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": flights})
}

// CreateFlightgodoc
// @Summary Create a new Flight
// @Description Create a new Flight with the given input data
// @Tags flights
// @Security ApiKeyAuth
// @Security JwtAuth
// @Accept  json
// @Produce  json
// @Param   input     body   models.CreateBook   true   "Create Flight object"
// @Success 201 {object} models.Flight "Successfully created Flight"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /flights [post]
func CreateFlight(c *gin.Context) {
	var input models.CreateFlight

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Flight := models.Flight{
		Price:      input.Price,
		Deals:      input.Deals,
		Animation:  input.Animation,
		SelectId:   input.SelectId,
		FlightList: input.FlightList,
	}

	database.Database.DB.Create(&Flight)

	// Invalidate cache
	keysPattern := "flights_offset_*"
	keys, err := cache.Rdb.Keys(cache.Ctx, keysPattern).Result()
	if err == nil {
		for _, key := range keys {
			cache.Rdb.Del(cache.Ctx, key)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": Flight})
}

// FindFlight godoc
// @Summary Find a Flight by ID
// @Description Get details of a Flight by its ID
// @Tags flights
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Flight ID"
// @Success 200 {object} models.Flight "Successfully retrieved Flight"
// @Failure 404 {string} string "Flight not found"
// @Router /flights/{id} [get]
func FindFlight(c *gin.Context) {
	var Flight models.Flight

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Flight).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flight not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Flight})
}

// UpdateFlight godoc
// @Summary Update a Flight by ID
// @Description Update the Flight details for the given ID
// @Tags flights
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "Flight ID"
// @Param input body models.UpdateBook true "Update Flight object"
// @Success 200 {object} models.Flight "Successfully updated Flight"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Flight not found"
// @Router /flights/{id} [put]
func UpdateFlight(c *gin.Context) {
	var Flight models.Flight
	var input models.UpdateFlight

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Flight).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flight not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.DB.Model(&Flight).Updates(models.Flight{
		ID:         0,
		Price:      input.Price,
		Deals:      input.Deals,
		Animation:  input.Animation,
		SelectId:   input.SelectId,
		FlightList: input.FlightList,
	})

	c.JSON(http.StatusOK, gin.H{"data": Flight})
}

// DeleteFlight godoc
// @Summary Delete a Flight by ID
// @Description Delete the Flight with the given ID
// @Tags flights
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Flight ID"
// @Success 204 {string} string "Successfully deleted Flight"
// @Failure 404 {string} string "Flight not found"
// @Router /flights/{id} [delete]
func DeleteFlight(c *gin.Context) {
	var Flight models.Flight

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Flight).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flight not found"})
		return
	}

	database.Database.DB.Delete(&Flight)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
