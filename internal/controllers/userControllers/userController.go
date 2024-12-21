package userControllers

import (
	"TIPPr4/internal/database"
	"TIPPr4/internal/helpers"
	"TIPPr4/internal/models"
	"TIPPr4/internal/myUtils"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
)

var validate = validator.New()

// Signup godoc
// @Summary Registers a new User
// @Description This endpoint allows you to register a new User by providing required fields: name, second_name, email, phone_number, password, and role_id. It validates the input, hashes the password, and saves the userControllers in the database. Also, it creates shoppingCart for this user
// @Tags Users
// @Accept json
// @Produce json
// @Param signup body dto.UserSignupDTO true "User data to register"
// @Success 201 {object} map[string]interface{} "User registered successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 409 {object} map[string]interface{} "Email or phone number already in use"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/auth/signup [post]
func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		var user models.User
		var cart models.ShoppingCart
		var favorite models.Favorite
		var library models.Library

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		password, err := myUtils.HashPassword(user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		user.Password = password
		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		token, refreshToken, _ := helpers.GenerateAllTokens(user.Email, user.Name, user.SecondName, user.Role.Type, user.ID)
		user.Token = token
		user.RefreshToken = refreshToken

		// Сохранение пользователя в базу
		if err := database.DB.WithContext(ctx).Create(&user).Error; err != nil {
			// Проверяем нарушение уникальности
			var pqErr *pgconn.PgError
			if errors.As(err, &pqErr) && pqErr.Code == "23505" {
				if pqErr.ConstraintName == "users_email_key" {
					c.JSON(http.StatusConflict, gin.H{"error": "Email already in use"})
				} else if pqErr.ConstraintName == "users_phone_number_key" {
					c.JSON(http.StatusConflict, gin.H{"error": "Phone number already in use"})
				} else {
					c.JSON(http.StatusConflict, gin.H{"error": "Unique constraint violation"})
				}
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "User item was not created"})
			}
			return
		}

		cart.User = user
		if err := database.DB.WithContext(ctx).Create(&cart).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cart item was not created"})
			return
		}

		favorite.User = user
		if err := database.DB.WithContext(ctx).Create(&favorite).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Favorite item was not created"})
			return
		}

		library.User = user
		if err := database.DB.WithContext(ctx).Create(&library).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Favorite item was not created"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
	}
}

// Login godoc
// @Summary Logs in a User and returns user data
// @Description This endpoint allows the userControllers to log in by providing email and password. It checks if the userControllers exists, verifies the password, generates access and refresh tokens, updates the tokens in the database, and sets them as cookies in the response.
// @Tags Users
// @Accept json
// @Produce json
// @Param login body dto.UserLoginDTO true "User credentials (email and password)"
// @Success 200 {object} models.User "Successfully logged in and returned userControllers data"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Email or password is incorrect"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/auth/login [post]
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := database.DB.WithContext(ctx).Preload("Role").First(&foundUser, "email = ?", user.Email).Error
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Email or password is incorrect"})
			return
		}

		passwordIsValid := myUtils.VerifyPassword(user.Password, foundUser.Password)
		if !passwordIsValid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Email or password is incorrect"})
			return
		}

		token, refreshToken, err := helpers.GenerateAllTokens(foundUser.Email, foundUser.Name, foundUser.SecondName, foundUser.Role.Type, foundUser.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		}

		err = helpers.UpdateAllTokens(token, refreshToken, foundUser.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't update tokens in DB"})
			return
		}

		// Установка токена в cookie
		c.SetCookie("token", token, 3600*2, "/", "localhost", false, true)
		c.SetCookie("refreshToken", refreshToken, 3600*24*7, "/", "localhost", false, true)

		c.JSON(http.StatusOK, foundUser)
	}
}

// GetUserById godoc
// @Summary Get a User by ID
// @Description Fetches a Users by their ID from the database. The userControllers making the request must be authorized to access the requested user data. User can get access only to their data. Admin can get access to all users data.
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} models.User "Successfully retrieved the userControllers data"
// @Failure 400 {object} map[string]interface{} "Invalid userControllers ID"
// @Failure 500 {object} map[string]interface{} "Failed to fetch userControllers"
// @Router /api/v1/users/{user_id} [get]
func GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		if err := helpers.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User
		err := database.DB.WithContext(ctx).Preload("Role").First(&user, userId).Error
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, user)
	}
}

// GetPaginatedUsers godoc
// @Summary Get users with pagination
// @Description Fetches a paginated list of users from the database. Only admins can access this endpoint.
// @Tags Users
// @Accept json
// @Produce json
// @Param limit query int false "Limit number of results" default(10)
// @Param offset query int false "Offset for pagination" default(0)
// @Success 200 {array} models.User "Successfully retrieved the paginated users"
// @Failure 400 {object} map[string]interface{} "Invalid query parameters"
// @Failure 403 {object} map[string]interface{} "Forbidden"
// @Failure 500 {object} map[string]interface{} "Failed to fetch users"
// @Router /api/v1/users/paginated [get]
func GetPaginatedUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		var users []models.User

		// Извлекаем параметры пагинации (limit и offset) из запроса
		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10")) // По умолчанию 10 записей
		if err != nil || limit <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
			return
		}
		offset, err := strconv.Atoi(c.DefaultQuery("offset", "0")) // По умолчанию без пропуска
		if err != nil || offset < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset value"})
			return
		}

		query := database.DB.WithContext(ctx).Preload("Role").Limit(limit).Offset(offset)
		if err := query.Find(&users).Error; err != nil {
			log.Printf("Error fetching users: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Fetches a list of all users from the database. Only admins can access this endpoint.
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} models.User "Successfully retrieved all users"
// @Failure 403 {object} map[string]interface{} "Forbidden"
// @Failure 500 {object} map[string]interface{} "Failed to fetch users"
// @Router /api/v1/users [get]
func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		var users []models.User

		err := database.DB.WithContext(ctx).Preload("Role").Find(&users).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

// UpdateUser godoc
// @Summary Update a User data
// @Description Allows updating specific fields of a User. The User making the request must be authorized to update the specified User data. User can update only their own data excluding role. Admin can update all users data.
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Param user body dto.UserUpdateDTO true "User data to update"
// @Success 200 {object} models.User "Successfully updated the user"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 404 {object} map[string]interface{} "User not found"
// @Failure 500 {object} map[string]interface{} "Failed to update user"
// @Router /api/v1/users/{user_id} [patch]
func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем ID пользователя из URL
		userID := c.Param("user_id")

		if err := helpers.MatchUserTypeToUid(c, userID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Создаем контекст с тайм-аутом
		var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// Получаем данные пользователя, которые нужно обновить
		var userUpdates models.User
		if err := c.ShouldBindJSON(&userUpdates); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
			return
		}

		// Проверяем, существует ли пользователь в базе данных
		var user models.User
		if err := database.DB.WithContext(ctx).First(&user, "id = ?", userID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
			return
		}

		// Обновляем только те поля, которые были переданы в запросе
		if userUpdates.Name != "" {
			user.Name = userUpdates.Name
		}
		if userUpdates.SecondName != "" {
			user.SecondName = userUpdates.SecondName
		}
		if userUpdates.ThirdName != "" {
			user.ThirdName = userUpdates.ThirdName
		}
		if userUpdates.Email != "" {
			user.Email = userUpdates.Email
		}
		if userUpdates.PhoneNumber != "" {
			user.PhoneNumber = userUpdates.PhoneNumber
		}
		if userUpdates.Password != "" {
			// Необходимо хэшировать пароль перед сохранением в базе
			hashedPassword, err := myUtils.HashPassword(userUpdates.Password)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
				return
			}
			user.Password = hashedPassword
		}
		if userUpdates.PaymentInfo != "" {
			user.PaymentInfo = userUpdates.PaymentInfo
		}
		err := helpers.CheckUserType(c, "ADMIN")
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		if userUpdates.RoleID != 0 {
			user.RoleID = userUpdates.RoleID
		}

		// Обновляем дату обновления
		user.UpdatedAt = time.Now()

		// Сохраняем обновленные данные пользователя в базу данных
		if err := database.DB.WithContext(ctx).Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		// Возвращаем успешный ответ с обновленными данными
		c.JSON(http.StatusOK, gin.H{
			"message": "User updated successfully",
			"user":    user,
		})
	}
}
