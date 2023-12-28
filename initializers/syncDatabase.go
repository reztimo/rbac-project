package initializers

import "maoelana/RBAC-project/models"

func SyncDatabase() {
	DB.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.UserRole{},
		&models.RolePermission{},
		&models.Event{},
	)
}
