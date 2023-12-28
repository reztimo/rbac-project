package controllers

import (
	"maoelana/RBAC-project/initializers"
	"maoelana/RBAC-project/models"

	"github.com/gin-gonic/gin"
)

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

func GetAllPermission(c *gin.Context) {
	var permissions []models.Permission
	initializers.DB.Find(&permissions)

	c.JSON(200, gin.H{
		"permission": permissions,
	})
}

func GetPermission(c *gin.Context) {
	id := c.Param("id")

	var permission models.Permission

	initializers.DB.First(&permission, id)

	c.JSON(200, gin.H{
		"permission": permission,
	})
}

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

func DeletePermission(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Permission{}, id)

	c.Status(200)
}
