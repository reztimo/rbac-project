package main

import (
	"maoelana/RBAC-project/controllers"
	"maoelana/RBAC-project/initializers"
	"maoelana/RBAC-project/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariabels()
	initializers.ConnectToDB()
	//initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	// Endpoint-event
	r.POST("/event", middlewares.RequireAuth, middlewares.AuthorizationMiddleware("Admin"), controllers.CreateResource)
	r.GET("/event", middlewares.RequireAuth, controllers.GetAllResource)
	r.GET("/event/:id", middlewares.RequireAuth, controllers.GetResource)
	r.PUT("/event/:id", middlewares.RequireAuth, middlewares.AuthorizationMiddleware("Admin"), controllers.UpdateResource)
	r.DELETE("/event/:id", middlewares.RequireAuth, middlewares.AuthorizationMiddleware("Admin"), controllers.DeleteResource)

	// Endpoint-permission
	r.POST("/permission", middlewares.RequireAuth, middlewares.AuthorizationMiddleware("Admin"), controllers.CreatePermission)
	r.GET("/permission", middlewares.RequireAuth, middlewares.AuthorizationMiddleware("Admin"), controllers.GetAllPermission)
	r.GET("/permission/:id", middlewares.RequireAuth, middlewares.AuthorizationMiddleware("Admin"), controllers.GetPermission)
	r.PUT("/permission/:id", middlewares.RequireAuth, middlewares.AuthorizationMiddleware("Admin"), controllers.UpdatePermission)
	r.DELETE("/permission/:id", middlewares.RequireAuth, middlewares.AuthorizationMiddleware("Admin"), controllers.DeletePermission)

	// Endpoint-role
	r.POST("/role", middlewares.RequireAuth, middlewares.AuthorizationMiddleware("Admin"), controllers.CreateRole)
	r.GET("/role", middlewares.RequireAuth, middlewares.AuthorizationMiddleware("Admin"), controllers.GetAllRole)
	r.GET("/role/:id", middlewares.RequireAuth, middlewares.AuthorizationMiddleware("Admin"), controllers.GetRole)
	r.PUT("/role/:id", middlewares.RequireAuth, middlewares.AuthorizationMiddleware("Admin"), controllers.EditRole)
	r.DELETE("/role/:id", middlewares.RequireAuth, middlewares.AuthorizationMiddleware("Admin"), controllers.RemoveRole)

	// Endpoint-user
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middlewares.RequireAuth, controllers.Validate)
	r.GET("/profile", middlewares.RequireAuth, controllers.GetAllProfile)
	r.GET("/profile/:id", middlewares.RequireAuth, controllers.GetProfile)
	r.GET("/logout", middlewares.RequireAuth, controllers.Logout)

	r.Run()
}
