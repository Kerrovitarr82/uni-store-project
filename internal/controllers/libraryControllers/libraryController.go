package libraryControllers

import (
	"TIPPr4/internal/database"
	"TIPPr4/internal/dto"
	"TIPPr4/internal/helpers"
	"TIPPr4/internal/models"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// GetLibrary godoc
// @Summary Get favorite
// @Description Get the user's library
// @Tags Library
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} dto.LibraryDTO
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/library/{user_id} [get]
func GetLibrary() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		userID := c.Param("user_id")

		if err := helpers.MatchUserTypeToUid(c, userID); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		var library models.Library
		if err := database.DB.WithContext(ctx).Preload("User").Preload("Games").Where("user_id = ?", userID).First(&library).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Library item not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		respLib := dto.LibraryDTO{
			ID:     library.ID,
			UserID: library.UserID,
			Games:  library.Games,
		}

		c.JSON(http.StatusOK, respLib)
	}
}
