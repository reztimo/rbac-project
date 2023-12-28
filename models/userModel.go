package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Roles    []Role `gorm:"many2many:user_roles"` //merujuk ke Role di roleModel.go
}
