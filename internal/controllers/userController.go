package controllers

import (
	"TIPPr4/internal/database"
	"TIPPr4/internal/helpers"
	"TIPPr4/internal/models"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true

	if err != nil {
		check = false
	}
	return check
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		password := HashPassword(user.Password)
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

		c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
	}
}

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
		err := database.DB.WithContext(ctx).First(&foundUser, "email = ?", user.Email).Error
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Email or password is incorrect"})
			return
		}

		passwordIsValid := VerifyPassword(user.Password, foundUser.Password)
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

func GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		if err := helpers.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User
		err := database.DB.WithContext(ctx).First(&user, userId).Error
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, user)
	}
}

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

		query := database.DB.WithContext(ctx).Limit(limit).Offset(offset)
		if err := query.Find(&users).Error; err != nil {
			log.Printf("Error fetching users: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		var users []models.User

		err := database.DB.WithContext(ctx).Find(&users).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}
