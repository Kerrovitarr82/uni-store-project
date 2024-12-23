package gameControllers

import (
	"TIPPr4/internal/database"
	"TIPPr4/internal/dto"
	"TIPPr4/internal/helpers"
	"TIPPr4/internal/models"
	"errors"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateDeveloper godoc
// @Summary Create Developer
// @Description Create a new developer
// @Tags Developer
// @Accept json
// @Produce json
// @Param developer body dto.DeveloperDTO true "Developer Data"
// @Success 201 {object} models.Developer
// @Failure 400 {object} map[string]interface{} "Validation Error"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/v1/developers [post]
func CreateDeveloper() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		var d dto.DeveloperDTO
		if err := c.ShouldBindJSON(&d); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		developer := models.Developer{
			Name:        d.Name,
			Email:       d.Email,
			Description: d.Description,
			Country:     d.Country,
		}

		if err := database.DB.Create(&developer).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, developer)
	}
}

// GetDeveloper godoc
// @Summary Get Developer
// @Description Get a developer by ID
// @Tags Developer
// @Accept json
// @Produce json
// @Param id path int true "Developer ID"
// @Success 200 {object} models.Developer
// @Failure 404 {object} map[string]interface{} "Developer Not Found"
// @Router /api/v1/developers/{id} [get]
func GetDeveloper() gin.HandlerFunc {
	return func(c *gin.Context) {
		var developer models.Developer
		id := c.Param("developer_id")

		if err := database.DB.Preload("Games").First(&developer, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Developer not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, developer)
	}
}

// GetAllDevelopers godoc
// @Summary Get All Developers
// @Description Get all developers
// @Tags Developer
// @Accept json
// @Produce json
// @Success 200 {array} models.Developer
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/v1/developers [get]
func GetAllDevelopers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var developers []models.Developer
		if err := database.DB.Find(&developers).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, developers)
	}
}

// GetPaginatedDevelopers godoc
// @Summary Get Paginated Developers
// @Description Get developers with pagination
// @Tags Developer
// @Accept json
// @Produce json
// @Param page query int true "Page number"
// @Param limit query int true "Page size"
// @Success 200 {object} map[string]interface{} "Paginated Developers"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/v1/developers/paginated [get]
func GetPaginatedDevelopers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var developers []models.Developer
		var total int64

		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
		offset := (page - 1) * limit

		if err := database.DB.Model(&models.Developer{}).Count(&total).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := database.DB.Limit(limit).Offset(offset).Find(&developers).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":       developers,
			"total":      total,
			"page":       page,
			"totalPages": int(math.Ceil(float64(total) / float64(limit))),
		})
	}
}

// UpdateDeveloper godoc
// @Summary Update Developer
// @Description Update a developer
// @Tags Developer
// @Accept json
// @Produce json
// @Param id path int true "Developer ID"
// @Param developer body dto.DeveloperUpdateDTO true "Developer Data"
// @Success 200 {object} models.Developer
// @Failure 400 {object} map[string]interface{} "Validation Error"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 404 {object} map[string]interface{} "Developer Not Found"
// @Router /api/v1/developers/{id} [patch]
func UpdateDeveloper() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		var d dto.DeveloperUpdateDTO
		id := c.Param("developer_id")

		if err := c.ShouldBindJSON(&d); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var developer models.Developer
		if err := database.DB.First(&developer, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Developer not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if d.Name != "" {
			developer.Name = d.Name
		}
		if d.Email != "" {
			developer.Email = d.Email
		}
		if d.Description != "" {
			developer.Description = d.Description
		}
		if d.Country != "" {
			developer.Country = d.Country
		}

		if err := database.DB.Save(&developer).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, developer)
	}
}

// DeleteDeveloper godoc
// @Summary Delete Developer
// @Description Delete a developer by ID
// @Tags Developer
// @Accept json
// @Produce json
// @Param id path int true "Developer ID"
// @Success 204 {string} string "No Content"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 404 {object} map[string]interface{} "Developer Not Found"
// @Router /api/v1/developers/{id} [delete]
func DeleteDeveloper() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		id := c.Param("developer_id")

		if err := database.DB.Delete(&models.Developer{}, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Developer not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
