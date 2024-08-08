package models

import "gorm.io/gorm"

type User struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Username string
	Password string
	Email    string `gorm:"unique"`
}

type UserOperate struct{}

func (u UserOperate) IsTableEmpty(db *gorm.DB) bool {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		panic("Error querying database:")
	}
	return result.RowsAffected == 0
}
