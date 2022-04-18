package model

type User struct {
	ID           int    `gorm:"primary_key;column:id;type:int;"`
	Username     string `gorm:"column:username;type:varchar(255)"`
	PasswordHash string `gorm:"column:password_hash;type:text;"`
}
