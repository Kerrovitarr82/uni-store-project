package favoriteControllers

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
	"uniStore/internal/database"
	"uniStore/internal/dto"
	"uniStore/internal/helpers"
	"uniStore/internal/models"
)

// AddGameToFavorite godoc
// @Summary Add Game to favorite
// @Description Add a specific game to the user's favorite
// @Tags Favorite
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param game_id path int true "Game ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/favorite/{user_id}/add/{game_id} [post]
func AddGameToFavorite() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		userID := c.Param("user_id")
		gameID := c.Param("game_id")

		if err := helpers.MatchUserTypeToUid(c, userID); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("Adding to favorite: user_id = %s, game_id = %s\n", userID, gameID)

		// Проверка наличия пользователя
		var user models.User
		if err := database.DB.WithContext(ctx).First(&user, userID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Проверка наличия игры
		var game models.Game
		if err := database.DB.WithContext(ctx).First(&game, gameID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(game)

		var favorite models.Favorite
		if err := database.DB.WithContext(ctx).Where("user_id = ?", userID).First(&favorite).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Favorite item not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(favorite)

		if err := database.DB.WithContext(ctx).Model(&favorite).Association("Games").Append(&game); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Game added to favorite"})
	}
}

// RemoveGameFromFavorite godoc
// @Summary Remove Game from favorite
// @Description Remove a specific game from the user's favorite
// @Tags Favorite
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param game_id path int true "Game ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/favorite/{user_id}/remove/{game_id} [delete]
func RemoveGameFromFavorite() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		userID := c.Param("user_id")
		gameID := c.Param("game_id")

		if err := helpers.MatchUserTypeToUid(c, userID); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		// Проверка наличия пользователя
		var user models.User
		if err := database.DB.WithContext(ctx).First(&user, userID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Проверка наличия игры
		var game models.Game
		if err := database.DB.WithContext(ctx).First(&game, gameID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var favorite models.Favorite
		if err := database.DB.WithContext(ctx).Where("user_id = ?", userID).First(&favorite).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Favorite not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Удаление игры из корзины
		if err := database.DB.WithContext(ctx).Model(&favorite).Association("Games").Delete(&game); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Game removed from favorite"})
	}
}

// GetFavorite godoc
// @Summary Get favorite
// @Description Get the user's favorite contents
// @Tags Favorite
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} dto.FavoriteDTO
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/favorite/{user_id} [get]
func GetFavorite() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		userID := c.Param("user_id")

		if err := helpers.MatchUserTypeToUid(c, userID); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		var favorite models.Favorite
		if err := database.DB.WithContext(ctx).Preload("User").Preload("Games").Where("user_id = ?", userID).First(&favorite).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Favorite item not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		respFav := dto.FavoriteDTO{
			ID:     favorite.ID,
			UserID: favorite.UserID,
			Games:  favorite.Games,
		}

		c.JSON(http.StatusOK, respFav)
	}
}

// ClearFavorite godoc
// @Summary Clear favorite
// @Description Clear all games from the user's favorite
// @Tags Favorite
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/favorite/{user_id}/clear [delete]
func ClearFavorite() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		userID := c.Param("user_id")

		if err := helpers.MatchUserTypeToUid(c, userID); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		var favorite models.Favorite
		if err := database.DB.WithContext(ctx).Where("user_id = ?", userID).First(&favorite).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Favorite item not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := database.DB.WithContext(ctx).Model(&favorite).Association("Games").Clear(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Favorite cleared"})
	}
}
