package main

import (
	"maoelana/RBAC-project/controllers"
	"maoelana/RBAC-project/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariabels()
	initializers.ConnectToDB()
	//initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	// resource event
	r.POST("/event", controllers.CreateResource)
	r.GET("/event", controllers.GetAllResource)
	r.GET("/event/:id", controllers.GetResource)
	r.PUT("/event/:id", controllers.UpdateResource)
	r.DELETE("/event/:id", controllers.DeleteResource)

	// permission
	r.POST("/permission", controllers.CreatePermission)
	r.GET("/permission", controllers.GetAllPermission)
	r.GET("/permission/:id", controllers.GetPermission)
	r.PUT("/permission/:id", controllers.UpdatePermission)
	r.DELETE("/permission/:id", controllers.DeletePermission) //warning

	// role
	r.POST("/role", controllers.CreateRole)
	r.GET("/role", controllers.GetAllRole)
	r.GET("/role/:id", controllers.GetRole)
	r.PUT("/role/:id", controllers.EditRole)
	r.DELETE("/role/:id", controllers.RemoveRole) //warning

	// user
	r.POST("/register", controllers.Register)

	r.Run()
}
