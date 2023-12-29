package controllers

import (
	"maoelana/RBAC-project/initializers"
	"maoelana/RBAC-project/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

		var permissions []models.Permission
		initializers.DB.Model(&role).Association("Permission").Find(&permissions)

		role.Permission = permissions

		roles = append(roles, role)
	}

	initializers.DB.Model(&user).Association("Roles").Append(roles)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func Login(c *gin.Context) {
	var userData struct {
		Username string
		Password string
	}

	if c.Bind(&userData) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read userData",
		})
	}

	var user models.User
	initializers.DB.First(&user, "username = ?", userData.Username)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		//"token": tokenString,
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func GetAllProfile(c *gin.Context) {
	var users []models.User
	initializers.DB.Preload("Roles").Find(&users)

	for i := range users {
		for j := range users[i].Roles {
			initializers.DB.Model(&users[i].Roles[j]).Association("Permission").Find(&users[i].Roles[j].Permission)
		}
	}

	c.JSON(200, gin.H{
		"user": users,
	})
}

func GetProfile(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	initializers.DB.Preload("Roles").Find(&user, id)

	for i := range user.Roles {
		initializers.DB.Model(&user.Roles[i]).Association("Permission").Find(&user.Roles[i].Permission)
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout",
	})
}

/*func AssignRole(c *gin.Context) {

}*/
