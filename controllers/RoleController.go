package controllers

import (
	"maoelana/RBAC-project/initializers"
	"maoelana/RBAC-project/models"

	"github.com/gin-gonic/gin"
)

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

func GetAllRole(c *gin.Context) {
	var roles []models.Role
	initializers.DB.Preload("Permission").Find(&roles)

	c.JSON(200, gin.H{
		"role": roles,
	})
}

func GetRole(c *gin.Context) {
	id := c.Param("id")

	var role models.Role
	initializers.DB.Preload("Permission").First(&role, id)

	c.JSON(200, gin.H{
		"role": role,
	})
}

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

func RemoveRole(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Role{}, id)

	c.Status(200)
}
