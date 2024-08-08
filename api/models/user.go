package models

type User struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Username string
	Password string
	Email    string `gorm:"unique"`
}
