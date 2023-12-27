package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name       string       `gorm:"unique"`
	Permission []Permission `gorm:"many2many:role_permission"`
}
