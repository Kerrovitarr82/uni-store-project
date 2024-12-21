package orderControllers

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
	"strconv"
	"time"
)

// CreateOrderFromCart godoc
// @Summary Create Order from Shopping Cart
// @Description Create a new order using the contents of the user's shopping cart
// @Tags Orders
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} dto.OrderDTO
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/orders/{user_id}/create [post]
func CreateOrderFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		userID := c.Param("user_id")
		if err := helpers.MatchUserTypeToUid(c, userID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user models.User
		var library models.Library

		if err := database.DB.WithContext(ctx).First(&user, userID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := database.DB.WithContext(ctx).Where("user_id = ?", userID).First(&library).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "User library not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Получение корзины пользователя
		var cart models.ShoppingCart
		if err := database.DB.WithContext(ctx).Preload("Games").Where("user_id = ?", userID).First(&cart).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Shopping cart not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Проверка, есть ли игры в корзине
		if len(cart.Games) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Shopping cart is empty"})
			return
		}

		// Расчет общей стоимости
		totalCost := 0.0
		for _, game := range cart.Games {
			totalCost += game.Price
		}

		// Создание нового заказа
		order := models.Order{
			UserID:    user.ID,
			Games:     cart.Games,
			TotalCost: totalCost,
		}

		if err := database.DB.WithContext(ctx).Create(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Очистка корзины после создания заказа
		if err := database.DB.WithContext(ctx).Model(&cart).Association("Games").Clear(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Добавление игр в библиотеку
		if err := database.DB.WithContext(ctx).Model(&library).Association("Games").Append(order.Games); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := database.DB.WithContext(ctx).
			Preload("Games.Developer").Preload("Games.Categories").Preload("Games.Restricts").Preload("Games").
			First(&order, order.ID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		respOrder := dto.OrderDTO{
			ID:        order.ID,
			UserID:    order.UserID,
			Games:     order.Games,
			TotalCost: order.TotalCost,
			CreatedAt: order.CreatedAt,
		}

		c.JSON(http.StatusOK, respOrder)
	}
}

// GetOrderByID godoc
// @Summary Get Order by ID
// @Description Get the details of a specific order by ID
// @Tags Orders
// @Accept json
// @Produce json
// @Param order_id path int true "Order ID"
// @Success 200 {object} dto.OrderDTO
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/orders/{order_id} [get]
func GetOrderByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderID := c.Param("order_id")

		// Создание контекста с таймаутом
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		var order models.Order
		if err := database.DB.WithContext(ctx).
			Preload("Games.Developer").Preload("Games.Categories").Preload("Games.Restricts").Preload("Games").
			First(&order, orderID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err := helpers.MatchUserTypeToUid(c, strconv.Itoa(order.UserID)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		respOrder := dto.OrderDTO{
			ID:        order.ID,
			UserID:    order.UserID,
			Games:     order.Games,
			TotalCost: order.TotalCost,
			CreatedAt: order.CreatedAt,
		}

		c.JSON(http.StatusOK, respOrder)
	}
}

// GetUserOrders godoc
// @Summary Get User's Orders
// @Description Get all orders of a specific user
// @Tags Orders
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} dto.OrderDTO
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/orders/user/{user_id} [get]
func GetUserOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("user_id")
		if err := helpers.MatchUserTypeToUid(c, userID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Создание контекста с таймаутом
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		var orders []models.Order
		if err := database.DB.WithContext(ctx).
			Preload("Games.Developer").Preload("Games.Categories").Preload("Games.Restricts").Preload("Games").
			Where("user_id = ?", userID).Find(&orders).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var respOrders []dto.OrderDTO

		for _, order := range orders {
			respOrder := dto.OrderDTO{
				ID:        order.ID,
				UserID:    order.UserID,
				Games:     order.Games,
				TotalCost: order.TotalCost,
				CreatedAt: order.CreatedAt,
			}
			respOrders = append(respOrders, respOrder)
		}

		c.JSON(http.StatusOK, respOrders)
	}
}

// GetAllOrders godoc
// @Summary Get All Orders
// @Description Get a list of all orders in the system
// @Tags Orders
// @Accept json
// @Produce json
// @Success 200 {array} dto.OrderDTO
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/orders [get]
func GetAllOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Создание контекста с таймаутом
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var orders []models.Order
		if err := database.DB.WithContext(ctx).Preload("Games").Find(&orders).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var respOrders []dto.OrderDTO

		for _, order := range orders {
			respOrder := dto.OrderDTO{
				ID:        order.ID,
				UserID:    order.UserID,
				Games:     order.Games,
				TotalCost: order.TotalCost,
				CreatedAt: order.CreatedAt,
			}
			respOrders = append(respOrders, respOrder)
		}

		c.JSON(http.StatusOK, respOrders)
	}
}
