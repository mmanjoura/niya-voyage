package hotels

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

// FindHotels godoc
// @Summary Get all hotels with pagination
// @Description Get a list of all hotels with optional pagination
// @Tags hotels
// @Security ApiKeyAuth
// @Produce json
// @Param offset query int false "Offset for pagination" default(0)
// @Param limit query int false "Limit for pagination" default(10)
// @Success 200 {array} models.Hotel "Successfully retrieved list of hotels"
// @Router /hotels [get]
func FindHotels(c *gin.Context) {
	var hotels []models.Hotel

	// Get query params
	offsetQuery := c.DefaultQuery("offset", "0")
	limitQuery := c.DefaultQuery("limit", "4")

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
	// cacheKey := "hotels_offset_" + offsetQuery + "_limit_" + limitQuery

	// Try fetching the data from Redis first
	// cachedBooks, err := cache.Rdb.Get(cache.Ctx, cacheKey).Result()
	// if err == nil {
	// 	err := json.Unmarshal([]byte(cachedBooks), &hotels)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal cached data"})
	// 		return
	// 	}cd
	// 	c.JSON(http.StatusOK, gin.H{"data": hotels})
	// 	return
	// }

	// If cache missed, fetch data from the database
	// dB := database.Database.DB

	if err := database.Database.DB.Offset(offset).Limit(limit).Raw(`SELECT id,
															category_id,
															tag,
															img,
															title,
															location,
															ratings,
															Reviews,
															price,
															Animation,
															city,
															Created_At,
															Updated_At
														FROM hotels`).Scan(&hotels).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})

		return
	}

	galleryImages := []models.GalleryImage{}
	slideImages := []models.SlideImage{}

	for i, v := range hotels {
		database.Database.DB.Find(&galleryImages, "hotel_id = ?", v.ID)
		database.Database.DB.Find(&slideImages, "hotel_id = ?", v.ID)
		hotels[i].GalleryImages = galleryImages
		hotels[i].SlideImages = slideImages
		for _, v := range slideImages {
			hotels[i].SlideImg = append(hotels[i].SlideImg, v.Img)
		}

	}

	// Serialize hotels object and store it in Redis
	//serializedBooks, err := json.Marshal(hotels)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
		return
	}
	//err = cache.Rdb.Set(cache.Ctx, cacheKey, serializedBooks, time.Minute).Err() // Here TTL is set to one hour
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": hotels})
}

// CreateHotelgodoc
// @Summary Create a new Hotel
// @Description Create a new Hotel with the given input data
// @Tags hotels
// @Security ApiKeyAuth
// @Security JwtAuth
// @Accept  json
// @Produce  json
// @Param   input     body   models.CreateBook   true   "Create Hotel object"
// @Success 201 {object} models.Hotel "Successfully created Hotel"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /hotels [post]
func CreateHotel(c *gin.Context) {
	var input models.CreateHotel

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Hotel := models.Hotel{Category_ID: input.Category_ID,
		Tag:       input.Tag,
		Img:       input.Img,
		Title:     input.Title,
		Location:  input.Location,
		Ratings:   input.Ratings,
		Reviews:   input.Reviews,
		Price:     input.Price,
		Animation: input.Animation,
		City:      input.City}

	database.Database.DB.Create(&Hotel)

	// Invalidate cache
	keysPattern := "hotels_offset_*"
	keys, err := cache.Rdb.Keys(cache.Ctx, keysPattern).Result()
	if err == nil {
		for _, key := range keys {
			cache.Rdb.Del(cache.Ctx, key)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": Hotel})
}

// FindHotel godoc
// @Summary Find a Hotel by ID
// @Description Get details of a Hotel by its ID
// @Tags hotels
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Hotel ID"
// @Success 200 {object} models.Hotel "Successfully retrieved Hotel"
// @Failure 404 {string} string "Hotel not found"
// @Router /hotels/{id} [get]
func FindHotel(c *gin.Context) {
	var hotel models.Hotel

	galleryImages := []models.GalleryImage{}
	if err := database.Database.DB.Raw(`SELECT id,
											category_id,
											tag,
											img,
											title,
											location,
											ratings,
											Reviews,
											price,
											Animation,
											city,
											Created_At,
											Updated_At
										FROM hotels where id = ` + c.Param("id")).Scan(&hotel).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "hotel not found"})
		return
	}

	if err := database.Database.DB.Find(&galleryImages, "hotel_id = ?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "hotel not found"})
		return
	}
	for _, v := range galleryImages {
		hotel.GalleryImg = append(hotel.GalleryImg, v.Img)
	}
	hotel.GalleryImages = galleryImages
	c.JSON(http.StatusOK, gin.H{"data": hotel})
}

// UpdateHotel godoc
// @Summary Update a Hotel by ID
// @Description Update the Hotel details for the given ID
// @Tags hotels
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "Hotel ID"
// @Param input body models.UpdateBook true "Update Hotel object"
// @Success 200 {object} models.Hotel "Successfully updated Hotel"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Hotel not found"
// @Router /hotels/{id} [put]
func UpdateHotel(c *gin.Context) {
	var Hotel models.Hotel
	var input models.UpdateHotel

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Hotel).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.DB.Model(&Hotel).Updates(models.Hotel{Category_ID: input.Category_ID,
		Tag:       input.Tag,
		Img:       input.Img,
		Title:     input.Title,
		Location:  input.Location,
		Ratings:   input.Ratings,
		Reviews:   input.Reviews,
		Price:     input.Price,
		Animation: input.Animation,
		City:      input.City})

	c.JSON(http.StatusOK, gin.H{"data": Hotel})
}

// DeleteHotel godoc
// @Summary Delete a Hotel by ID
// @Description Delete the Hotel with the given ID
// @Tags hotels
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Hotel ID"
// @Success 204 {string} string "Successfully deleted Hotel"
// @Failure 404 {string} string "Hotel not found"
// @Router /hotels/{id} [delete]
func DeleteHotel(c *gin.Context) {
	var Hotel models.Hotel

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Hotel).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
		return
	}

	database.Database.DB.Delete(&Hotel)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
