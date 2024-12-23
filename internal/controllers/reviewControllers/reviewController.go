package reviewControllers

import (
	"TIPPr4/internal/database"
	"TIPPr4/internal/dto"
	"TIPPr4/internal/helpers"
	"TIPPr4/internal/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateReview godoc
// @Summary Create Review
// @Description Create a review for a specific game
// @Tags Review
// @Accept json
// @Produce json
// @Param review body dto.ReviewDTO true "Review details"
// @Success 201 {object} models.Review
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/reviews/{game_id}/user/{user_id} [post]
func CreateReview() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		var review models.Review
		var userID = c.Param("user_id")
		var gameID = c.Param("game_id")

		if err := c.ShouldBindJSON(&review); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if err := helpers.MatchUserTypeToUid(c, userID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Проверка пользователя и игры
		var user models.User
		if err := database.DB.WithContext(ctx).First(&user, userID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		var game models.Game
		if err := database.DB.WithContext(ctx).First(&game, gameID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
			return
		}

		// Сохранение отзыва
		review.CreatedAt = time.Now()
		review.UpdatedAt = time.Now()
		if err := database.DB.WithContext(ctx).Create(&review).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, review)
	}
}

// UpdateReview godoc
// @Summary Update Review
// @Description Update an existing review
// @Tags Review
// @Accept json
// @Produce json
// @Param review_id path int true "Review ID"
// @Param review body dto.ReviewUpdateDTO true "Updated Review details"
// @Success 200 {object} models.Review
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/reviews/{review_id}/user/{user_id} [patch]
func UpdateReview() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		reviewID := c.Param("review_id")
		userId := c.Param("user_id")

		if err := helpers.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var review models.Review
		if err := database.DB.WithContext(ctx).First(&review, reviewID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
			return
		}

		var updateData dto.ReviewUpdateDTO
		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if updateData.Title != "" {
			review.Title = updateData.Title
		}
		if updateData.Description != "" {
			review.Description = updateData.Description
		}
		if updateData.Rating != 0 {
			review.Rating = updateData.Rating
		}

		if err := database.DB.WithContext(ctx).Save(&review).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, review)
	}
}

// DeleteReview godoc
// @Summary Delete Review
// @Description Delete a specific review
// @Tags Review
// @Accept json
// @Produce json
// @Param review_id path int true "Review ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/reviews/{review_id}/user/{user_id} [delete]
func DeleteReview() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		reviewID := c.Param("review_id")
		userId := c.Param("user_id")

		if err := helpers.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := database.DB.WithContext(ctx).Delete(&models.Review{}, reviewID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Review deleted"})
	}
}

// GetReviewByID godoc
// @Summary Get Review by ID
// @Description Get a specific review by its ID
// @Tags Review
// @Accept json
// @Produce json
// @Param review_id path int true "Review ID"
// @Success 200 {object} models.Review
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/reviews/{review_id} [get]
func GetReviewByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		reviewID := c.Param("review_id")
		var review models.Review
		if err := database.DB.WithContext(ctx).
			Preload("User").
			Preload("Game").
			First(&review, reviewID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
			return
		}

		c.JSON(http.StatusOK, review)
	}
}

// GetReviewsByGameID godoc
// @Summary Get Reviews by Game ID
// @Description Get all reviews for a specific game
// @Tags Review
// @Accept json
// @Produce json
// @Param game_id path int true "Game ID"
// @Success 200 {array} models.Review
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/reviews/game/{game_id} [get]
func GetReviewsByGameID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		gameID := c.Param("game_id")
		var reviews []models.Review
		if err := database.DB.WithContext(ctx).
			Preload("User").
			Where("game_id = ?", gameID).
			Find(&reviews).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, reviews)
	}
}
