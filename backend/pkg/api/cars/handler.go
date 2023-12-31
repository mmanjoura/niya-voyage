package cars

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

// FindCars godoc
// @Summary Get all cars with pagination
// @Description Get a list of all cars with optional pagination
// @Tags cars
// @Security ApiKeyAuth
// @Produce json
// @Param offset query int false "Offset for pagination" default(0)
// @Param limit query int false "Limit for pagination" default(10)
// @Success 200 {array} models.Car "Successfully retrieved list of cars"
// @Router /cars [get]
func FindCars(c *gin.Context) {
	var cars []models.Car

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
	// cacheKey := "cars_offset_" + offsetQuery + "_limit_" + limitQuery

	// Try fetching the data from Redis first
	// cachedCars, err := cache.Rdb.Get(cache.Ctx, cacheKey).Result()
	// if err == nil {
	// 	err := json.Unmarshal([]byte(cachedCars), &cars)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal cached data"})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{"data": cars})
	// 	return
	// }

	// If cache missed, fetch data from the database
	// dB := database.Database.DB
	// dB.Model(&models.Car{}).Preload("SlideImages").Preload("GalleryImages").Offset(offset).Limit(limit).Find(&cars)
	database.Database.DB.Offset(offset).Limit(limit).Raw(`SELECT ID,
															tag,
															title,
															price,
															location,
															reviews,
															ratings,
															animation,
															seat,
															type,
															luggage,
															transmission,
															speed,
															Created_At,
															Updated_At
														FROM Cars`).Scan(&cars)

	galleryImages := []models.GalleryImage{}
	slideImages := []models.SlideImage{}

	for i, v := range cars {
		database.Database.DB.Find(&galleryImages, "car_id = ?", v.ID)
		database.Database.DB.Find(&slideImages, "car_id = ?", v.ID)
		cars[i].GalleryImages = galleryImages
		cars[i].SlideImages = slideImages
		for _, v := range slideImages {
			cars[i].SlideImg = append(cars[i].SlideImg, v.Img)
		}

	}

	// Serialize cars object and store it in Redis
	//serializedCars, err := json.Marshal(cars)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
		return
	}
	//err = cache.Rdb.Set(cache.Ctx, cacheKey, serializedCars, time.Minute).Err() // Here TTL is set to one hour
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cars})
}

// CreateCar godoc
// @Summary Create a new car
// @Description Create a new car with the given input data
// @Tags cars
// @Security ApiKeyAuth
// @Security JwtAuth
// @Accept  json
// @Produce  json
// @Param   input     body   models.CreateCar   true   "Create car object"
// @Success 201 {object} models.Car "Successfully created car"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /cars [post]
func CreateCar(c *gin.Context) {
	var input models.CreateCar

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	car := models.Car{

		Tag:          input.Tag,
		Title:        input.Title,
		Price:        input.Price,
		Location:     input.Location,
		Duration:     input.Duration,
		Reviews:      input.Reviews,
		Seat:         input.Seat,
		Type:         input.Type,
		Luggage:      input.Luggage,
		Transmission: input.Transmission,
		Speed:        input.Speed,
		Animation:    input.Animation,
	}

	database.Database.DB.Create(&car)

	// Invalidate cache
	keysPattern := "cars_offset_*"
	keys, err := cache.Rdb.Keys(cache.Ctx, keysPattern).Result()
	if err == nil {
		for _, key := range keys {
			cache.Rdb.Del(cache.Ctx, key)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": car})
}

// FindCar godoc
// @Summary Find a car by ID
// @Description Get details of a car by its ID
// @Tags cars
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Car ID"
// @Success 200 {object} models.Car "Successfully retrieved car"
// @Failure 404 {string} string "Car not found"
// @Router /cars/{id} [get]
func FindCar(c *gin.Context) {
	var car models.Car

	galleryImages := []models.GalleryImage{}
	if err := database.Database.DB.Raw(`SELECT ID,
								tag,
								title,
								price,
								location,
								reviews,
								ratings,
								animation,
								seat,
								type,
								luggage,
								transmission,
								speed,
								Created_At,
								Updated_At
							FROM Cars where id = ` + c.Param("id")).Scan(&car).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "car not found"})
		return
	}

	if err := database.Database.DB.Find(&galleryImages, "car_id = ?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "car not found"})
		return
	}

	for _, v := range galleryImages {
		car.GalleryImg = append(car.GalleryImg, v.Img)
	}
	c.JSON(http.StatusOK, gin.H{"data": car})
}

// UpdateCar godoc
// @Summary Update a car by ID
// @Description Update the car details for the given ID
// @Tags cars
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "Car ID"
// @Param input body models.UpdateCar true "Update car object"
// @Success 200 {object} models.Car "Successfully updated car"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "car not found"
// @Router /cars/{id} [put]
func UpdateCar(c *gin.Context) {
	var car models.Car
	var input models.UpdateCar

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "car not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.DB.Model(&car).Updates(models.Car{
		Tag:          input.Tag,
		Title:        input.Title,
		Price:        input.Price,
		Location:     input.Location,
		Duration:     input.Duration,
		Reviews:      input.Reviews,
		Seat:         input.Seat,
		Type:         input.Type,
		Luggage:      input.Luggage,
		Transmission: input.Transmission,
		Speed:        input.Speed,
		Animation:    input.Animation,
	})

	c.JSON(http.StatusOK, gin.H{"data": car})
}

// DeleteCar godoc
// @Summary Delete a car by ID
// @Description Delete the car with the given ID
// @Tags cars
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Car ID"
// @Success 204 {string} string "Successfully deleted car"
// @Failure 404 {string} string "car not found"
// @Router /cars/{id} [delete]
func DeleteCar(c *gin.Context) {
	var car models.Car

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "car not found"})
		return
	}

	database.Database.DB.Delete(&car)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
