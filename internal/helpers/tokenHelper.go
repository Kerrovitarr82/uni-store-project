package helpers

import (
	"TIPPr4/internal/database"
	"TIPPr4/internal/models"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type SignedDetails struct {
	Name       string
	SecondName string
	Email      string
	Uid        int
	UserRole   string
	jwt.RegisteredClaims
}

func GenerateAllTokens(email string, name string, secondName string, userRole string, uid int) (signedToken string, signedRefreshToken string, err error) {
	claims := SignedDetails{
		Email:      email,
		Name:       name,
		SecondName: secondName,
		Uid:        uid,
		UserRole:   userRole,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	refreshClaims := SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("SECRET_KEY")))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return token, refreshToken, err
	}
	return token, refreshToken, nil
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userId int) error {
	// Устанавливаем тайм-аут для операции
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Подготавливаем обновляемые данные
	updates := map[string]interface{}{
		"token":         signedToken,
		"refresh_token": signedRefreshToken,
		"updated_at":    time.Now(),
	}

	// Выполняем обновление токенов в бд
	if err := database.DB.WithContext(ctx).Model(&models.User{}).Where("id = ?", userId).Updates(updates).Error; err != nil {
		return err
	}

	return nil
}

func ValidateToken(signedToken string, signedRefreshToken string) (*SignedDetails, error) {
	// Парсим токен с передачей структуры, которая будет содержать данные о токене
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		},
	)

	refreshToken, err := jwt.ParseWithClaims(
		signedRefreshToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		},
	)

	// Если произошла ошибка при парсинге токена, возвращаем ошибку
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	// Преобразуем claims в ожидаемый тип и проверяем
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		return nil, fmt.Errorf("the token has invalid claims")
	}
	refreshClaims, ok := refreshToken.Claims.(*SignedDetails)
	if !ok {
		return nil, fmt.Errorf("the refreshToken has invalid claims")
	}

	// Проверяем, не истек ли срок действия токена
	if claims.RegisteredClaims.ExpiresAt.Unix() < time.Now().UTC().Unix() {
		if refreshClaims.RegisteredClaims.ExpiresAt.Unix() < time.Now().UTC().Unix() {
			return nil, fmt.Errorf("tokens are expired")
		}
		newToken, newRefreshToken, err := GenerateAllTokens(claims.Email, claims.Name, claims.SecondName, claims.UserRole, claims.Uid)
		if err != nil {
			return nil, fmt.Errorf("token generation failed: %w", err)
		}
		err = UpdateAllTokens(newToken, newRefreshToken, claims.Uid)
		if err != nil {
			return nil, fmt.Errorf("token update failed: %w", err)
		}
	}

	return claims, nil
}
