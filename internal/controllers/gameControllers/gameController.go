package gameControllers

import (
	"TIPPr4/internal/database"
	"TIPPr4/internal/dto"
	"TIPPr4/internal/helpers"
	"TIPPr4/internal/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"net/http"
	"strconv"
)

// CreateGame godoc
// @Summary Create Game
// @Description Create a new gameControllers
// @Tags Game
// @Accept json
// @Produce json
// @Param gameControllers body dto.GameDTO true "Game Data"
// @Success 201 {object} models.Game
// @Failure 400 {object} map[string]interface{} "Validation Error"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/v1/games [post]
func CreateGame() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		var d dto.GameDTO
		if err := c.ShouldBindJSON(&d); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var categories []models.Category
		for _, id := range d.CategoryIds {
			categories = append(categories, models.Category{ID: id})
		}

		var restricts []models.Restrict
		for _, id := range d.RestrictIds {
			restricts = append(restricts, models.Restrict{ID: id})
		}

		game := models.Game{
			Name:           d.Name,
			Description:    d.Description,
			Size:           d.Size,
			Price:          d.Price,
			AgeRestriction: d.AgeRestriction,
			DeveloperId:    d.DeveloperId,
			Categories:     categories,
			Restricts:      restricts,
		}

		if err := database.DB.Create(&game).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, game)
	}
}

// GetGame godoc
// @Summary Get Game
// @Description Get a gameControllers by ID
// @Tags Game
// @Accept json
// @Produce json
// @Param id path int true "Game ID"
// @Success 200 {object} models.Game
// @Failure 404 {object} map[string]interface{} "Game Not Found"
// @Router /api/v1/games/{id} [get]
func GetGame() gin.HandlerFunc {
	return func(c *gin.Context) {
		var game models.Game
		id := c.Param("game_id")

		if err := database.DB.Preload("Developer").Preload("Categories").Preload("Restricts").First(&game, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, game)
	}
}

// GetAllGames godoc
// @Summary Get All Games
// @Description Get all games
// @Tags Game
// @Accept json
// @Produce json
// @Success 200 {array} models.Game
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/v1/games [get]
func GetAllGames() gin.HandlerFunc {
	return func(c *gin.Context) {
		var games []models.Game
		if err := database.DB.Preload("Developer").Preload("Categories").Preload("Restricts").Find(&games).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, games)
	}
}

// GetPaginatedGames godoc
// @Summary Get Paginated Games
// @Description Get games with pagination
// @Tags Game
// @Accept json
// @Produce json
// @Param page query int true "Page number"
// @Param limit query int true "Page size"
// @Success 200 {object} map[string]interface{} "Paginated Games"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/v1/games/paginated [get]
func GetPaginatedGames() gin.HandlerFunc {
	return func(c *gin.Context) {
		var games []models.Game
		var total int64

		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
		offset := (page - 1) * limit

		if err := database.DB.Model(&models.Game{}).Count(&total).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := database.DB.Preload("Developer").Preload("Categories").Preload("Restricts").Limit(limit).Offset(offset).Find(&games).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":       games,
			"total":      total,
			"page":       page,
			"totalPages": int(math.Ceil(float64(total) / float64(limit))),
		})
	}
}

// UpdateGame godoc
// @Summary Update Game
// @Description Update a gameControllers
// @Tags Game
// @Accept json
// @Produce json
// @Param id path int true "Game ID"
// @Param gameControllers body dto.GameDTO true "Game Data"
// @Success 200 {object} models.Game
// @Failure 400 {object} map[string]interface{} "Validation Error"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 404 {object} map[string]interface{} "Game Not Found"
// @Router /api/v1/games/{id} [patch]
func UpdateGame() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		var d dto.GameDTO
		id := c.Param("game_id")

		if err := c.ShouldBindJSON(&d); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var game models.Game
		if err := database.DB.First(&game, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var categories []models.Category
		for _, id := range d.CategoryIds {
			categories = append(categories, models.Category{ID: id})
		}

		var restricts []models.Restrict
		for _, id := range d.RestrictIds {
			restricts = append(restricts, models.Restrict{ID: id})
		}

		if d.Name != "" {
			game.Name = d.Name
		}
		if d.Description != "" {
			game.Description = d.Description
		}
		if d.Size != 0 {
			game.Size = d.Size
		}
		if d.Price != 0 {
			game.Price = d.Price
		}
		if d.AgeRestriction != "" {
			game.AgeRestriction = d.AgeRestriction
		}
		if d.DeveloperId != 0 {
			game.DeveloperId = d.DeveloperId
		}
		if categories != nil {
			game.Categories = categories
		}
		if restricts != nil {
			game.Restricts = restricts
		}

		if err := database.DB.Save(&game).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, game)
	}
}

// DeleteGame godoc
// @Summary Delete Game
// @Description Delete a gameControllers by ID
// @Tags Game
// @Accept json
// @Produce json
// @Param id path int true "Game ID"
// @Success 204 {string} string "No Content"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 404 {object} map[string]interface{} "Game Not Found"
// @Router /api/v1/games/{id} [delete]
func DeleteGame() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		id := c.Param("game_id")

		if err := database.DB.Delete(&models.Game{}, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
