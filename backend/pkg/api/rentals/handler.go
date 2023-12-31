package rentals

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

// FindRentals godoc
// @Summary Get all rentals with pagination
// @Description Get a list of all rentals with optional pagination
// @Tags rentals
// @Security ApiKeyAuth
// @Produce json
// @Param offset query int false "Offset for pagination" default(0)
// @Param limit query int false "Limit for pagination" default(10)
// @Success 200 {array} models.Rental "Successfully retrieved list of rentals"
// @Router /rentals [get]
func FindRentals(c *gin.Context) {
	var rentals []models.Rental

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
	// cacheKey := "rentals_offset_" + offsetQuery + "_limit_" + limitQuery

	// Try fetching the data from Redis first
	// cachedRentals, err := cache.Rdb.Get(cache.Ctx, cacheKey).Result()
	// if err == nil {
	// 	err := json.Unmarshal([]byte(cachedRentals), &rentals)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal cached data"})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{"data": rentals})
	// 	return
	// }

	// If cache missed, fetch data from the database
	// dB := database.Database.DB
	// dB.Model(&models.Rental{}).Preload("SlideImages").Preload("GalleryImages").Offset(offset).Limit(limit).Find(&rentals)

	database.Database.DB.Offset(offset).Limit(limit).Raw(`SELECT ID,
															tag,
															title,
															price,
															location,
															duration,
															reviews,
															ratings,
															animation,
															guest,
															bedroom,
															bed,
															Created_At,
															Updated_At
														FROM Rentals`).Scan(&rentals)

	galleryImages := []models.GalleryImage{}
	slideImages := []models.SlideImage{}

	for i, v := range rentals {
		database.Database.DB.Find(&galleryImages, "rental_id = ?", v.ID)
		database.Database.DB.Find(&slideImages, "rental_id = ?", v.ID)
		rentals[i].GalleryImages = galleryImages
		rentals[i].SlideImages = slideImages
		for _, v := range slideImages {
			rentals[i].SlideImg = append(rentals[i].SlideImg, v.Img)
		}

	}

	// Serialize rentals object and store it in Redis
	//serializedRentals, err := json.Marshal(rentals)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
		return
	}
	//err = cache.Rdb.Set(cache.Ctx, cacheKey, serializedRentals, time.Minute).Err() // Here TTL is set to one hour
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rentals})
}

// CreateRental godoc
// @Summary Create a new rental
// @Description Create a new rental with the given input data
// @Tags rentals
// @Security ApiKeyAuth
// @Security JwtAuth
// @Accept  json
// @Produce  json
// @Param   input     body   models.CreateRental   true   "Create rental object"
// @Success 201 {object} models.Rental "Successfully created rental"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /rentals [post]
func CreateRental(c *gin.Context) {
	var input models.CreateRental

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rental := models.Rental{

		Tag:       input.Tag,
		Title:     input.Title,
		Price:     input.Price,
		Location:  input.Location,
		Duration:  input.Duration,
		Reviews:   input.Reviews,
		Animation: input.Animation,
		Bedroom:   input.Bedroom,
		Bed:       input.Bed,
		Guest:     input.Guest,
	}

	database.Database.DB.Create(&rental)

	// Invalidate cache
	keysPattern := "rentals_offset_*"
	keys, err := cache.Rdb.Keys(cache.Ctx, keysPattern).Result()
	if err == nil {
		for _, key := range keys {
			cache.Rdb.Del(cache.Ctx, key)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": rental})
}

// FindRental godoc
// @Summary Find a rental by ID
// @Description Get details of a rental by its ID
// @Tags rentals
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Rental ID"
// @Success 200 {object} models.Rental "Successfully retrieved rental"
// @Failure 404 {string} string "Rental not found"
// @Router /rentals/{id} [get]
func FindRental(c *gin.Context) {
	var rental models.Rental

	galleryImages := []models.GalleryImage{}
	if err := database.Database.DB.Raw(`SELECT ID,
											tag,
											title,
											price,
											location,
											duration,
											reviews,
											ratings,
											animation,
											guest,
											bedroom,
											bed,
											Created_At,
											Updated_At
										FROM Rentals where id = ` + c.Param("id")).Scan(&rental).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "rental not found"})
		return
	}

	if err := database.Database.DB.Find(&galleryImages, "rental_id = ?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "rental not found"})
		return
	}
	for _, v := range galleryImages {
		rental.GalleryImg = append(rental.GalleryImg, v.Img)
	}
	rental.GalleryImages = galleryImages
	c.JSON(http.StatusOK, gin.H{"data": rental})

}

// UpdateRental godoc
// @Summary Update a rental by ID
// @Description Update the rental details for the given ID
// @Tags rentals
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "Rental ID"
// @Param input body models.UpdateRental true "Update rental object"
// @Success 200 {object} models.Rental "Successfully updated rental"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "rental not found"
// @Router /rentals/{id} [put]
func UpdateRental(c *gin.Context) {
	var rental models.Rental
	var input models.UpdateRental

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&rental).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "rental not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.DB.Model(&rental).Updates(models.Rental{
		Tag:       input.Tag,
		Title:     input.Title,
		Price:     input.Price,
		Location:  input.Location,
		Duration:  input.Duration,
		Reviews:   input.Reviews,
		Animation: input.Animation,
		Bedroom:   input.Bedroom,
		Bed:       input.Bed,
		Guest:     input.Guest,
	})

	c.JSON(http.StatusOK, gin.H{"data": rental})
}

// DeleteRental godoc
// @Summary Delete a rental by ID
// @Description Delete the rental with the given ID
// @Tags rentals
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Rental ID"
// @Success 204 {string} string "Successfully deleted rental"
// @Failure 404 {string} string "rental not found"
// @Router /rentals/{id} [delete]
func DeleteRental(c *gin.Context) {
	var rental models.Rental

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&rental).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "rental not found"})
		return
	}

	database.Database.DB.Delete(&rental)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
