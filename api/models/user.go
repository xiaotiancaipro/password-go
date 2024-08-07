package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	id       float64 `gorm:"primaryKey;autoIncrement"`
	username string
	password string
	email    string
}
