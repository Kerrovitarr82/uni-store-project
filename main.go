package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strconv"
	"sync"
)

type Mouse struct {
	ID       int     `json:"id"`
	Brand    string  `json:"brand"`
	Model    string  `json:"model"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

// Моделируем базу данных в виде карты и мьютекса для синхронизации доступа
var mouseStore = make(map[int]Mouse)
var idCounter = 1
var mu sync.Mutex

func createMouse(c *gin.Context) {
	var mice []Mouse
	if err := c.ShouldBindJSON(&mice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	for i, m := range mice {
		m.ID = idCounter
		mouseStore[idCounter] = m
		idCounter++
		mice[i] = m
	}
	mu.Unlock()

	c.JSON(http.StatusCreated, mice)
}

func getAllMice(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	var mice []Mouse
	for _, mouse := range mouseStore {
		mice = append(mice, mouse)
	}

	sort.Slice(mice, func(i, j int) bool {
		return mice[i].ID < mice[j].ID
	})

	c.JSON(http.StatusOK, mice)
}

func getMouseByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	mu.Lock()
	mouse, exists := mouseStore[id]
	mu.Unlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mouse not found"})
		return
	}

	c.JSON(http.StatusOK, mouse)
}

func updateMouse(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedMouse Mouse
	if err := c.ShouldBindJSON(&updatedMouse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	mouse, exists := mouseStore[id]
	mu.Unlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mouse not found"})
		return
	}

	if updatedMouse.Brand != "" {
		mouse.Brand = updatedMouse.Brand
	}
	if updatedMouse.Model != "" {
		mouse.Model = updatedMouse.Model
	}
	if updatedMouse.Price != 0 {
		mouse.Price = updatedMouse.Price
	}
	if updatedMouse.Quantity != 0 {
		mouse.Quantity = updatedMouse.Quantity
	}

	mu.Lock()
	mouseStore[id] = mouse
	mu.Unlock()

	c.JSON(http.StatusOK, mouse)
}

func deleteMouse(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	mu.Lock()
	_, exists := mouseStore[id]
	if exists {
		delete(mouseStore, id)
	}
	mu.Unlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mouse not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Mouse deleted successfully"})
}

func main() {
	r := gin.Default()

	r.POST("/mice", createMouse)       // Создать новый товар
	r.GET("/mice", getAllMice)         // Получить список всех товаров
	r.GET("/mice/:id", getMouseByID)   // Получить товар по ID
	r.PATCH("/mice/:id", updateMouse)  // Обновить товар по ID
	r.DELETE("/mice/:id", deleteMouse) // Удалить товар по ID

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
