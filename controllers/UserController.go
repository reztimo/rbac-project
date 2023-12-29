package controllers

import (
	"maoelana/RBAC-project/initializers"
	"maoelana/RBAC-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var userData struct {
		Username string
		Password string
		Roles    []string
	}

	if c.Bind(&userData) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read userData",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	user := models.User{
		Username: userData.Username,
		Password: string(hash),
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
	}

	var roles []models.Role
	for _, roleName := range userData.Roles {
		var role models.Role
		initializers.DB.Where("name = ?", roleName).First(&role)
		roles = append(roles, role)
	}

	initializers.DB.Model(&user).Association("Roles").Append(roles)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func Login(c *gin.Context) {
}

func Validate(c *gin.Context) {

}

func GetProfile(c *gin.Context) {

}

func Logout(c *gin.Context) {

}

/*func AssignRole(c *gin.Context) {

}*/
