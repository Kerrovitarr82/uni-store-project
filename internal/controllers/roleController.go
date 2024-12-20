package controllers

import (
	"TIPPr4/internal/database"
	"TIPPr4/internal/helpers"
	"TIPPr4/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// CreateRole godoc
// @Summary Create a new role
// @Description This endpoint allows an admin to create a new role by providing the role type and description.
// @Tags Roles
// @Accept json
// @Produce json
// @Param role body models.Role true "Role data"
// @Success 201 {object} models.Role "Role created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 403 {object} map[string]interface{} "Forbidden, only admins can create roles"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/roles [post]
func CreateRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Проверка, что пользователь является администратором
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		var role models.Role
		if err := c.ShouldBindJSON(&role); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Сохранение роли в базу данных
		if err := database.DB.Create(&role).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.JSON(http.StatusCreated, role)
	}
}

// GetAllRoles godoc
// @Summary Get all roles
// @Description This endpoint allows an admin to fetch all roles.
// @Tags Roles
// @Produce json
// @Success 200 {array} models.Role "List of roles"
// @Failure 403 {object} map[string]interface{} "Forbidden, only admins can access roles"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/roles [get]
func GetAllRoles() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Проверка, что пользователь является администратором
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		var roles []models.Role
		if err := database.DB.Find(&roles).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, roles)
	}
}

// GetRoleById godoc
// @Summary Get role by ID
// @Description This endpoint allows an admin to fetch a role by its ID.
// @Tags Roles
// @Produce json
// @Param role_id path int true "Role ID"
// @Success 200 {object} models.Role "Role found"
// @Failure 404 {object} map[string]interface{} "Role not found"
// @Failure 403 {object} map[string]interface{} "Forbidden, only admins can access roles"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/roles/{role_id} [get]
func GetRoleById() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Проверка, что пользователь является администратором
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		roleID := c.Param("role_id")

		var role models.Role
		if err := database.DB.First(&role, "id = ?", roleID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
			return
		}

		c.JSON(http.StatusOK, role)
	}
}

// UpdateRole godoc
// @Summary Update a role
// @Description This endpoint allows an admin to update a role by providing its ID and new data.
// @Tags Roles
// @Accept json
// @Produce json
// @Param role_id path int true "Role ID"
// @Param role body models.Role true "Updated role data"
// @Success 200 {object} models.Role "Role updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 404 {object} map[string]interface{} "Role not found"
// @Failure 403 {object} map[string]interface{} "Forbidden, only admins can update roles"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/roles/{role_id} [patch]
func UpdateRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Проверка, что пользователь является администратором
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		roleID := c.Param("role_id")

		var role models.Role
		// Получаем роль по ID
		if err := database.DB.First(&role, "id = ?", roleID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
			return
		}

		// Получаем данные для обновления
		var roleUpdates models.Role
		if err := c.ShouldBindJSON(&roleUpdates); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Обновляем только те поля, которые были переданы
		if roleUpdates.Type != "" {
			role.Type = roleUpdates.Type
		}
		if roleUpdates.Description != "" {
			role.Description = roleUpdates.Description
		}

		// Обновляем дату обновления
		role.UpdatedAt = time.Now()

		// Сохраняем обновленную роль в базу данных
		if err := database.DB.Save(&role).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, role)
	}
}

// DeleteRole godoc
// @Summary Delete a role
// @Description This endpoint allows an admin to delete a role by its ID.
// @Tags Roles
// @Produce json
// @Param role_id path int true "Role ID"
// @Success 200 {object} map[string]interface{} "Role deleted successfully"
// @Failure 404 {object} map[string]interface{} "Role not found"
// @Failure 403 {object} map[string]interface{} "Forbidden, only admins can delete roles"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/roles/{role_id} [delete]
func DeleteRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Проверка, что пользователь является администратором
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		roleID := c.Param("role_id")

		var role models.Role
		if err := database.DB.First(&role, "id = ?", roleID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
			return
		}

		if err := database.DB.Delete(&role).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
	}
}
