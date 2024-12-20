package helpers

import (
	"TIPPr4/internal/database"
	"TIPPr4/internal/models"
	"context"
	"github.com/golang-jwt/jwt/v5"
	"log"
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
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	refreshClaims := SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("SECRET_KEY")))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		log.Panic(err)
		return
	}
	return token, refreshToken, nil
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userId id) (err error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var user models.User

	if err := database.DB.WithContext(ctx).First(&user, userId).Error; err != nil {
		return err // тут приколы
	}

	return
}
