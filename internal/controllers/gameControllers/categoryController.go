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

// CreateCategory godoc
// @Summary Create Category
// @Description Create a new category
// @Tags Category
// @Accept json
// @Produce json
// @Param category body dto.CategoryDTO true "Category Data"
// @Success 201 {object} models.Category
// @Failure 400 {object} map[string]interface{} "Validation Error"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/v1/categories [post]
func CreateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		var d dto.CategoryDTO
		if err := c.ShouldBindJSON(&d); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		category := models.Category{
			Name:        d.Name,
			Description: d.Description,
		}

		if err := database.DB.Create(&category).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, category)
	}
}

// GetCategory godoc
// @Summary Get Category
// @Description Get a category by ID
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} models.Category
// @Failure 404 {object} map[string]interface{} "Category Not Found"
// @Router /api/v1/categories/{id} [get]
func GetCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var category models.Category
		id := c.Param("category_id")

		if err := database.DB.Preload("Games").First(&category, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, category)
	}
}

// GetAllCategories godoc
// @Summary Get All Categories
// @Description Get all categories
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {array} models.Category
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/v1/categories [get]
func GetAllCategories() gin.HandlerFunc {
	return func(c *gin.Context) {
		var categories []models.Category
		if err := database.DB.Preload("Games").Find(&categories).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, categories)
	}
}

// GetPaginatedCategories godoc
// @Summary Get Paginated Categories
// @Description Get categories with pagination
// @Tags Category
// @Accept json
// @Produce json
// @Param page query int true "Page number"
// @Param limit query int true "Page size"
// @Success 200 {object} map[string]interface{} "Paginated Categories"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/v1/categories/paginated [get]
func GetPaginatedCategories() gin.HandlerFunc {
	return func(c *gin.Context) {
		var categories []models.Category
		var total int64

		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
		offset := (page - 1) * limit

		if err := database.DB.Model(&models.Category{}).Count(&total).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := database.DB.Preload("Games").Limit(limit).Offset(offset).Find(&categories).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":       categories,
			"total":      total,
			"page":       page,
			"totalPages": int(math.Ceil(float64(total) / float64(limit))),
		})
	}
}

// UpdateCategory godoc
// @Summary Update Category
// @Description Update a category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body dto.CategoryDTO true "Category Data"
// @Success 200 {object} models.Category
// @Failure 400 {object} map[string]interface{} "Validation Error"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 404 {object} map[string]interface{} "Category Not Found"
// @Router /api/v1/categories/{id} [patch]
func UpdateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		var d dto.CategoryDTO
		id := c.Param("category_id")

		if err := c.ShouldBindJSON(&d); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var category models.Category
		if err := database.DB.First(&category, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if d.Name != "" {
			category.Name = d.Name
		}
		if d.Description != "" {
			category.Description = d.Description
		}

		if err := database.DB.Save(&category).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, category)
	}
}

// DeleteCategory godoc
// @Summary Delete Category
// @Description Delete a category by ID
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 204 {string} string "No Content"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 404 {object} map[string]interface{} "Category Not Found"
// @Router /api/v1/categories/{id} [delete]
func DeleteCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		id := c.Param("category_id")

		if err := database.DB.Delete(&models.Category{}, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
