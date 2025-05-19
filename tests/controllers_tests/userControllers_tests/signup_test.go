package userControllers_tests

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"uniStore/internal/database"
	"uniStore/internal/helpers"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"uniStore/internal/controllers/userControllers"
	"uniStore/internal/middleware"
)

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, *sql.DB) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}

	dialector := postgres.New(postgres.Config{
		Conn: db,
	})

	gormDB, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("failed to open gorm DB: %v", err)
	}

	return gormDB, mock, db
}

func setupRouterWithMockDB(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// JWT middleware
	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.Use(middleware.Authenticate())
			users.GET("/", userControllers.GetAllUsers())
		}
	}

	return router
}

func TestGetUsers_WithAuthCookie(t *testing.T) {
	gdb, mock, sqlDB := setupMockDB(t)
	defer sqlDB.Close()

	database.DB = gdb

	router := setupRouterWithMockDB(gdb)

	// Мокаем SELECT * FROM "users"
	rows := sqlmock.NewRows([]string{"id", "name", "second_name", "email"}).
		AddRow(1, "Alice", "Wonderland", "alice@example.com").
		AddRow(2, "Bob", "Builder", "bob@example.com")
	mock.ExpectQuery(`SELECT \* FROM "users"`).WillReturnRows(rows)

	// Генерация токенов
	os.Setenv("SECRET_KEY", "test-secret") // для подписи JWT
	token, refreshToken, err := helpers.GenerateAllTokens(
		"alice@example.com", "Alice", "Wonderland", "ADMIN", 1,
	)
	assert.NoError(t, err)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/", nil)
	req.AddCookie(&http.Cookie{Name: "token", Value: token})
	req.AddCookie(&http.Cookie{Name: "refreshToken", Value: refreshToken})

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var users []map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &users)
	assert.NoError(t, err)
	assert.Len(t, users, 2)
}

func TestGetUsers_WithoutAuthCookie(t *testing.T) {
	gdb, _, sqlDB := setupMockDB(t)
	defer sqlDB.Close()

	database.DB = gdb

	router := setupRouterWithMockDB(gdb)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
