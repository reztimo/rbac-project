package controllers

import (
	"maoelana/RBAC-project/initializers"
	"maoelana/RBAC-project/models"

	"github.com/gin-gonic/gin"
)

// CreatePermission
// @Summary Create a new permission
// @Description Create a new permission by providing its name.
// @Tags permissions
// @Produce json
// @Param name formData string true "Name of the new permission"
// @Success 200 "Permission created successfully"
// @Router /permission [post]
func CreatePermission(c *gin.Context) {
	var permissionData struct {
		Name string
	}

	c.Bind(&permissionData)

	permission := models.Permission{
		Name: permissionData.Name,
	}

	result := initializers.DB.Create(&permission)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"permission": permission,
	})
}

// GetAllPermission
// @Summary Get all permissions
// @Description Retrieve a list of all permissions.
// @Tags permissions
// @Produce json
// @Success 200 "List of permissions retrieved successfully"
// @Router /permission [get]
func GetAllPermission(c *gin.Context) {
	var permissions []models.Permission
	initializers.DB.Find(&permissions)

	c.JSON(200, gin.H{
		"permission": permissions,
	})
}

// GetPermission
// @Summary Get a specific permission by ID
// @Description Retrieve a specific permission by providing its ID.
// @Tags permissions
// @Produce json
// @Param id path string true "Permission ID"
// @Success 200 "Permission retrieved successfully"
// @Router /permission/{id} [get]
func GetPermission(c *gin.Context) {
	id := c.Param("id")

	var permission models.Permission

	initializers.DB.First(&permission, id)

	c.JSON(200, gin.H{
		"permission": permission,
	})
}

// UpdatePermission
// @Summary Update a specific permission by ID
// @Description Update a specific permission by providing its ID and new data.
// @Tags permissions
// @Produce json
// @Param id path string true "Permission ID"
// @Param name formData string true "New name for the permission"
// @Success 200 "Permission updated successfully"
// @Router /permission/{id} [put]
func UpdatePermission(c *gin.Context) {
	id := c.Param("id")

	var permissionData struct {
		Name string
	}

	c.Bind(&permissionData)

	var permission models.Permission
	initializers.DB.First(&permission, id)

	initializers.DB.Model(&permission).Updates(models.Permission{
		Name: permissionData.Name,
	})

	c.JSON(200, gin.H{
		"permission": permission,
	})
}

// DeletePermission
// @Summary Delete a specific permission by ID
// @Description Delete a specific permission by providing its ID.
// @Tags permissions
// @Produce json
// @Param id path string true "Permission ID"
// @Success 200 "Permission deleted successfully"
// @Router /permission/{id} [delete]
func DeletePermission(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Permission{}, id)

	c.Status(200)
}
