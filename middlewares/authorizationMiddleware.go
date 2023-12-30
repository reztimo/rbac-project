package middlewares

import (
	"maoelana/RBAC-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware(requireRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		u, ok := user.(models.User)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if hasRequireRole(u.Roles, requireRole) {
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}

func hasRequireRole(roles []models.Role, requireRole string) bool {
	for _, role := range roles {
		if role.Name == requireRole {
			return true
		}
	}
	return false
}
