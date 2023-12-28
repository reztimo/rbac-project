package main

import (
	"maoelana/RBAC-project/controllers"
	"maoelana/RBAC-project/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariabels()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	// resource event
	r.POST("/event", controllers.CreateResource)
	r.GET("/event", controllers.GetAllResource)
	r.GET("/event/:id", controllers.GetResource)
	r.PUT("/event/:id", controllers.UpdateResource)
	r.DELETE("/event/:id", controllers.DeleteResource)

	r.Run()
}
