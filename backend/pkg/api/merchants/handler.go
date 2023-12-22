package merchants

import (
	"encoding/json"
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

// FindMerchants godoc
// @Summary Get all Merchants with pagination
// @Description Get a list of all Merchants with optional pagination
// @Tags Merchants
// @Security ApiKeyAuth
// @Produce json
// @Param offset query int false "Offset for pagination" default(0)
// @Param limit query int false "Limit for pagination" default(10)
// @Success 200 {array} models.Merchant "Successfully retrieved list of Merchants"
// @Router /Merchants [get]
func FindMerchants(c *gin.Context) {
	var merchants []models.Merchant

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
	cacheKey := "Merchants_offset_" + offsetQuery + "_limit_" + limitQuery

	// Try fetching the data from Redis first
	cachedMerchants, err := cache.Rdb.Get(cache.Ctx, cacheKey).Result()
	if err == nil {
		err := json.Unmarshal([]byte(cachedMerchants), &merchants)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal cached data"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": merchants})
		return
	}

	// If cache missed, fetch data from the database
	database.Database.DB.Offset(offset).Limit(limit).Find(&merchants)

	for i, v := range merchants {
		locationInfo := models.LocationInfo{}
		passwordChan := models.ChangePass{}
		database.Database.DB.Find(&locationInfo, "merchant_id = ?", v.ID)
		database.Database.DB.Find(&passwordChan, "merchant_id = ?", v.ID)
		// v.LocationInformation = locationInfo
		// v.ChangePassword = passwordChan
		merchants[i].LocationInfo = locationInfo
		merchants[i].ChangePass = passwordChan

	}

	// for i, v := range Merchants {
	// 	// Serialize each Merchant object and store it in Redis
	// 	serializedMerchant, err := json.Marshal(Merchants)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
	// 		return
	// 	}

	// Serialize Merchants object and store it in Redis
	// serializedMerchants, err := json.Marshal(Merchants)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
	// 	return
	// }
	// err = cache.Rdb.Set(cache.Ctx, cacheKey, serializedMerchants, time.Minute).Err() // Here TTL is set to one hour
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"data": merchants})
}

// CreateMerchant godoc
// @Summary Create a new Merchant
// @Description Create a new Merchant with the given input data
// @Tags Merchants
// @Security ApiKeyAuth
// @Security JwtAuth
// @Accept  json
// @Produce  json
// @Param   input     body   models.CreateMerchant   true   "Create Merchant object"
// @Success 201 {object} models.Merchant "Successfully created Merchant"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /Merchants [post]
func CreateMerchant(c *gin.Context) {
	var input models.CreateMerchant

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Merchant := models.Merchant{
		BusinessName: input.BusinessName,
		UserName:     input.UserName,
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
		BirthDate:    input.BirthDate,
		About:        input.About,
	}

	database.Database.DB.Create(&Merchant)

	// Invalidate cache
	keysPattern := "Merchants_offset_*"
	keys, err := cache.Rdb.Keys(cache.Ctx, keysPattern).Result()
	if err == nil {
		for _, key := range keys {
			cache.Rdb.Del(cache.Ctx, key)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": Merchant})
}

// FindMerchant godoc
// @Summary Find a Merchant by ID
// @Description Get details of a Merchant by its ID
// @Tags Merchants
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Merchant ID"
// @Success 200 {object} models.Merchant "Successfully retrieved Merchant"
// @Failure 404 {string} string "Merchant not found"
// @Router /Merchants/{id} [get]
func FindMerchant(c *gin.Context) {
	var Merchant models.Merchant

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Merchant).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Merchant not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Merchant})
}

// UpdateMerchant godoc
// @Summary Update a Merchant by ID
// @Description Update the Merchant details for the given ID
// @Tags Merchants
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "Merchant ID"
// @Param input body models.UpdateMerchant true "Update Merchant object"
// @Success 200 {object} models.Merchant "Successfully updated Merchant"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Merchant not found"
// @Router /Merchants/{id} [put]
func UpdateMerchant(c *gin.Context) {
	var Merchant models.Merchant
	var input models.UpdateMerchant

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Merchant).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Merchant not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.DB.Model(&Merchant).Updates(models.Merchant{
		BusinessName: input.BusinessName,
		UserName:     input.UserName,
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
		BirthDate:    input.BirthDate,
		About:        input.About,
	})

	c.JSON(http.StatusOK, gin.H{"data": Merchant})
}

// DeleteMerchant godoc
// @Summary Delete a Merchant by ID
// @Description Delete the Merchant with the given ID
// @Tags Merchants
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Merchant ID"
// @Success 204 {string} string "Successfully deleted Merchant"
// @Failure 404 {string} string "Merchant not found"
// @Router /Merchants/{id} [delete]
func DeleteMerchant(c *gin.Context) {
	var Merchant models.Merchant

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Merchant).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Merchant not found"})
		return
	}

	database.Database.DB.Delete(&Merchant)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
