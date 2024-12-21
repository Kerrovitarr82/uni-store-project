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

// CreateRestrict godoc
// @Summary Create Restrict
// @Description Create a new restrict entry
// @Tags Restrict
// @Accept json
// @Produce json
// @Param restrict body dto.RestrictDTO true "Restrict Data"
// @Success 201 {object} models.Restrict
// @Failure 400 {object} map[string]interface{} "Validation Error"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/v1/restricts [post]
func CreateRestrict() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		var d dto.RestrictDTO
		if err := c.ShouldBindJSON(&d); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		restrict := models.Restrict{
			GameID: d.GameID,
			Region: d.Region,
		}

		if err := database.DB.Create(&restrict).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, restrict)
	}
}

// GetRestrict godoc
// @Summary Get Restrict
// @Description Get a restrict entry by ID
// @Tags Restrict
// @Accept json
// @Produce json
// @Param id path int true "Restrict ID"
// @Success 200 {object} models.Restrict
// @Failure 404 {object} map[string]interface{} "Restrict Not Found"
// @Router /api/v1/restricts/{id} [get]
func GetRestrict() gin.HandlerFunc {
	return func(c *gin.Context) {
		var restrict models.Restrict
		id := c.Param("restrict_id")

		if err := database.DB.First(&restrict, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Restrict not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, restrict)
	}
}

// GetAllRestricts godoc
// @Summary Get All Restricts
// @Description Get all restricts
// @Tags Restrict
// @Accept json
// @Produce json
// @Success 200 {array} models.Restrict
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/v1/restricts [get]
func GetAllRestricts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var restricts []models.Restrict
		if err := database.DB.Find(&restricts).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, restricts)
	}
}

// GetPaginatedRestricts godoc
// @Summary Get Paginated Restricts
// @Description Get restricts with pagination
// @Tags Restrict
// @Accept json
// @Produce json
// @Param page query int true "Page number"
// @Param limit query int true "Page size"
// @Success 200 {object} map[string]interface{} "Paginated Restricts"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/v1/restricts/paginated [get]
func GetPaginatedRestricts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var restricts []models.Restrict
		var total int64

		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
		offset := (page - 1) * limit

		if err := database.DB.Model(&models.Restrict{}).Count(&total).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := database.DB.Limit(limit).Offset(offset).Find(&restricts).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":       restricts,
			"total":      total,
			"page":       page,
			"totalPages": int(math.Ceil(float64(total) / float64(limit))),
		})
	}
}

// UpdateRestrict godoc
// @Summary Update Restrict
// @Description Update a restrict entry
// @Tags Restrict
// @Accept json
// @Produce json
// @Param id path int true "Restrict ID"
// @Param restrict body dto.RestrictDTO true "Restrict Data"
// @Success 200 {object} models.Restrict
// @Failure 400 {object} map[string]interface{} "Validation Error"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 404 {object} map[string]interface{} "Restrict Not Found"
// @Router /api/v1/restricts/{id} [patch]
func UpdateRestrict() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		var d dto.RestrictDTO
		id := c.Param("restrict_id")

		if err := c.ShouldBindJSON(&d); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var restrict models.Restrict
		if err := database.DB.First(&restrict, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Restrict not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if d.GameID != 0 {
			restrict.GameID = d.GameID
		}
		if d.Region != "" {
			restrict.Region = d.Region
		}

		if err := database.DB.Save(&restrict).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, restrict)
	}
}

// DeleteRestrict godoc
// @Summary Delete Restrict
// @Description Delete a restrict entry by ID
// @Tags Restrict
// @Accept json
// @Produce json
// @Param id path int true "Restrict ID"
// @Success 204 {string} string "No Content"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 404 {object} map[string]interface{} "Restrict Not Found"
// @Router /api/v1/restricts/{id} [delete]
func DeleteRestrict() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		id := c.Param("restrict_id")

		if err := database.DB.Delete(&models.Restrict{}, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Restrict not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
