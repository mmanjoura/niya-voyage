package testimonials

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

// FindTestimonials godoc
// @Summary Get all testimonials with pagination
// @Description Get a list of all testimonials with optional pagination
// @Tags testimonials
// @Security ApiKeyAuth
// @Produce json
// @Param offset query int false "Offset for pagination" default(0)
// @Param limit query int false "Limit for pagination" default(10)
// @Success 200 {array} models.Testimonial "Successfully retrieved list of testimonials"
// @Router /testimonials [get]
func FindTestimonials(c *gin.Context) {
	var testimonials []models.Testimonial

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
	// cacheKey := "testimonials_offset_" + offsetQuery + "_limit_" + limitQuery

	// Try fetching the data from Redis first
	// cachedBooks, err := cache.Rdb.Get(cache.Ctx, cacheKey).Result()
	// if err == nil {
	// 	err := json.Unmarshal([]byte(cachedBooks), &testimonials)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal cached data"})
	// 		return
	// 	}cd
	// 	c.JSON(http.StatusOK, gin.H{"data": testimonials})
	// 	return
	// }

	// If cache missed, fetch data from the database
	database.Database.DB.Offset(offset).Limit(limit).Find(&testimonials)

	// Serialize testimonials object and store it in Redis
	//serializedBooks, err := json.Marshal(testimonials)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
		return
	}
	//err = cache.Rdb.Set(cache.Ctx, cacheKey, serializedBooks, time.Minute).Err() // Here TTL is set to one hour
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": testimonials})
}

// CreateTestimonialgodoc
// @Summary Create a new Testimonial
// @Description Create a new Testimonial with the given input data
// @Tags testimonials
// @Security ApiKeyAuth
// @Security JwtAuth
// @Accept  json
// @Produce  json
// @Param   input     body   models.CreateBook   true   "Create Testimonial object"
// @Success 201 {object} models.Testimonial "Successfully created Testimonial"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /testimonials [post]
func CreateTestimonial(c *gin.Context) {
	var input models.CreateTestimonial

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Testimonial := models.Testimonial{Hotel_ID: input.Hotel_ID,
		Meta:           input.Meta,
		Avatar:         input.Avatar,
		Name:           input.Name,
		Designation:    input.Designation,
		Text:           input.Text,
		DelayAnimation: input.DelayAnimation}

	database.Database.DB.Create(&Testimonial)

	// Invalidate cache
	keysPattern := "testimonials_offset_*"
	keys, err := cache.Rdb.Keys(cache.Ctx, keysPattern).Result()
	if err == nil {
		for _, key := range keys {
			cache.Rdb.Del(cache.Ctx, key)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": Testimonial})
}

// FindTestimonial godoc
// @Summary Find a Testimonial by ID
// @Description Get details of a Testimonial by its ID
// @Tags testimonials
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Testimonial ID"
// @Success 200 {object} models.Testimonial "Successfully retrieved Testimonial"
// @Failure 404 {string} string "Testimonial not found"
// @Router /testimonials/{id} [get]
func FindTestimonial(c *gin.Context) {
	var Testimonial models.Testimonial

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Testimonial).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Testimonial not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Testimonial})
}

// UpdateTestimonial godoc
// @Summary Update a Testimonial by ID
// @Description Update the Testimonial details for the given ID
// @Tags testimonials
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "Testimonial ID"
// @Param input body models.UpdateBook true "Update Testimonial object"
// @Success 200 {object} models.Testimonial "Successfully updated Testimonial"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Testimonial not found"
// @Router /testimonials/{id} [put]
func UpdateTestimonial(c *gin.Context) {
	var Testimonial models.Testimonial
	var input models.UpdateTestimonial

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Testimonial).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Testimonial not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.DB.Model(&Testimonial).Updates(models.Testimonial{Hotel_ID: input.Hotel_ID,
		Meta:           input.Meta,
		Avatar:         input.Avatar,
		Name:           input.Name,
		Designation:    input.Designation,
		Text:           input.Text,
		DelayAnimation: input.DelayAnimation})

	c.JSON(http.StatusOK, gin.H{"data": Testimonial})
}

// DeleteTestimonial godoc
// @Summary Delete a Testimonial by ID
// @Description Delete the Testimonial with the given ID
// @Tags testimonials
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Testimonial ID"
// @Success 204 {string} string "Successfully deleted Testimonial"
// @Failure 404 {string} string "Testimonial not found"
// @Router /testimonials/{id} [delete]
func DeleteTestimonial(c *gin.Context) {
	var Testimonial models.Testimonial

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Testimonial).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Testimonial not found"})
		return
	}

	database.Database.DB.Delete(&Testimonial)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
