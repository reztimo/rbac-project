package controllers

import (
	"maoelana/RBAC-project/initializers"
	"maoelana/RBAC-project/models"

	"github.com/gin-gonic/gin"
)

// CreateRole
// @Summary Create a new role
// @Description Create a new role by providing its name and permissions.
// @Tags roles
// @Produce json
// @Param name formData string true "Name of the role"
// @Param permissions formData array true "Array of permission names"
// @Success 200 "Role created successfully"
// @Router /role [post]
func CreateRole(c *gin.Context) {
	var roleData struct {
		Name       string
		Permission []string
	}

	c.Bind(&roleData)

	role := models.Role{
		Name: roleData.Name,
	}

	result := initializers.DB.Create(&role)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// get premission from table permission
	var permissions []models.Permission
	for _, permissionName := range roleData.Permission {
		var permission models.Permission
		initializers.DB.Where("name = ?", permissionName).First(&permission)
		permissions = append(permissions, permission)
	}

	initializers.DB.Model(&role).Association("Permission").Append(permissions)

	c.JSON(200, gin.H{
		"role": role,
	})
}

// GetAllRole
// @Summary Get all roles with associated permissions
// @Description Get a list of all roles along with their associated permissions.
// @Tags roles
// @Produce json
// @Success 200 "Roles retrieved successfully"
// @Router /role [get]
func GetAllRole(c *gin.Context) {
	var roles []models.Role
	initializers.DB.Preload("Permission").Find(&roles)

	c.JSON(200, gin.H{
		"role": roles,
	})
}

// GetRole
// @Summary Get a specific role by ID
// @Description Get details of a specific role by providing its ID.
// @Tags roles
// @Produce json
// @Param id path string true "Role ID"
// @Success 200 "Role retrieved successfully"
// @Router /role/{id} [get]
func GetRole(c *gin.Context) {
	id := c.Param("id")

	var role models.Role
	initializers.DB.Preload("Permission").First(&role, id)

	c.JSON(200, gin.H{
		"role": role,
	})
}

// EditRole
// @Summary Edit a specific role by ID
// @Description Edit details of a specific role by providing its ID and new data.
// @Tags roles
// @Produce json
// @Param id path string true "Role ID"
// @Param name body string true "name for role"
// @Param permissions body []string true "permission for user"
// @Success 200 "Role edited successfully"
// @Router /role/{id} [put]
func EditRole(c *gin.Context) {
	id := c.Param("id")

	var roleData struct {
		Name       string
		Permission []string
	}

	c.Bind(&roleData)

	var role models.Role
	initializers.DB.First(&role, id)

	initializers.DB.Model(&role).Updates(models.Role{
		Name: roleData.Name,
	})

	var permissions []models.Permission
	for _, permissionName := range roleData.Permission {
		var permission models.Permission
		initializers.DB.Where("name = ?", permissionName).First(&permission)
		permissions = append(permissions, permission)
	}

	// delete permission
	initializers.DB.Model(&role).Association("Permission").Clear()

	// add new permission
	initializers.DB.Model(&role).Association("Permission").Append(permissions)

	c.JSON(200, gin.H{
		"role": role,
	})
}

// RemoveRole
// @Summary Remove a specific role by ID
// @Description Remove a specific role by providing its ID.
// @Tags roles
// @Produce json
// @Param id path string true "Role ID"
// @Success 200 "Role removed successfully"
// @Router /role/{id} [delete]
func RemoveRole(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Role{}, id)

	c.Status(200)
}
