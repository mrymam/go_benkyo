package model

type User struct {
	ID           int `gorm:"primary_key;column:id;"`
	Username     string
	PasswordHash string
}
