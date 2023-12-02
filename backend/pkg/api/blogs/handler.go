package blogs

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

// FindBlogs godoc
// @Summary Get all blogs with pagination
// @Description Get a list of all blogs with optional pagination
// @Tags blogs
// @Security ApiKeyAuth
// @Produce json
// @Param offset query int false "Offset for pagination" default(0)
// @Param limit query int false "Limit for pagination" default(10)
// @Success 200 {array} models.Blog "Successfully retrieved list of blogs"
// @Router /blogs [get]
func FindBlogs(c *gin.Context) {
	var blogs []models.Blog

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
	// cacheKey := "blogs_offset_" + offsetQuery + "_limit_" + limitQuery

	// Try fetching the data from Redis first
	// cachedBooks, err := cache.Rdb.Get(cache.Ctx, cacheKey).Result()
	// if err == nil {
	// 	err := json.Unmarshal([]byte(cachedBooks), &blogs)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal cached data"})
	// 		return
	// 	}cd
	// 	c.JSON(http.StatusOK, gin.H{"data": blogs})
	// 	return
	// }

	// If cache missed, fetch data from the database
	database.Database.DB.Offset(offset).Limit(limit).Find(&blogs)

	// Serialize blogs object and store it in Redis
	//serializedBooks, err := json.Marshal(blogs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data"})
		return
	}
	//err = cache.Rdb.Set(cache.Ctx, cacheKey, serializedBooks, time.Minute).Err() // Here TTL is set to one hour
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": blogs})
}

// CreateBloggodoc
// @Summary Create a new Blog
// @Description Create a new Blog with the given input data
// @Tags blogs
// @Security ApiKeyAuth
// @Security JwtAuth
// @Accept  json
// @Produce  json
// @Param   input     body   models.CreateBook   true   "Create Blog object"
// @Success 201 {object} models.Blog "Successfully created Blog"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /blogs [post]
func CreateBlog(c *gin.Context) {
	var input models.CreateBlog

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Blog := models.Blog{Details: input.Details,
		Tag:            input.Tag,
		Tags:           input.Tags,
		Title:          input.Title,
		Img:            input.Img,
		DelayAnimation: input.DelayAnimation}

	database.Database.DB.Create(&Blog)

	// Invalidate cache
	keysPattern := "blogs_offset_*"
	keys, err := cache.Rdb.Keys(cache.Ctx, keysPattern).Result()
	if err == nil {
		for _, key := range keys {
			cache.Rdb.Del(cache.Ctx, key)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": Blog})
}

// FindBlog godoc
// @Summary Find a Blog by ID
// @Description Get details of a Blog by its ID
// @Tags blogs
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Blog ID"
// @Success 200 {object} models.Blog "Successfully retrieved Blog"
// @Failure 404 {string} string "Blog not found"
// @Router /blogs/{id} [get]
func FindBlog(c *gin.Context) {
	var Blog models.Blog

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Blog).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Blog})
}

// UpdateBlog godoc
// @Summary Update a Blog by ID
// @Description Update the Blog details for the given ID
// @Tags blogs
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "Blog ID"
// @Param input body models.UpdateBook true "Update Blog object"
// @Success 200 {object} models.Blog "Successfully updated Blog"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Blog not found"
// @Router /blogs/{id} [put]
func UpdateBlog(c *gin.Context) {
	var Blog models.Blog
	var input models.UpdateBlog

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Blog).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Database.DB.Model(&Blog).Updates(models.Blog{Details: input.Details,
		Tag:            input.Tag,
		Tags:           input.Tags,
		Title:          input.Title,
		Img:            input.Img,
		DelayAnimation: input.DelayAnimation})

	c.JSON(http.StatusOK, gin.H{"data": Blog})
}

// DeleteBlog godoc
// @Summary Delete a Blog by ID
// @Description Delete the Blog with the given ID
// @Tags blogs
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Blog ID"
// @Success 204 {string} string "Successfully deleted Blog"
// @Failure 404 {string} string "Blog not found"
// @Router /blogs/{id} [delete]
func DeleteBlog(c *gin.Context) {
	var Blog models.Blog

	if err := database.Database.DB.Where("id = ?", c.Param("id")).First(&Blog).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	database.Database.DB.Delete(&Blog)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
