package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint64
	Username string
}
