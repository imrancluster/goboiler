package models

import (
	"github.com/jinzhu/gorm"
)

// User : user model
type User struct {
	gorm.Model
	Username  string `gorm:"not null;type:varchar(100);unique_index" json:"username"`
	Password  string `gorm:"null;type:varchar(255)" json:"password"`
	CreatedBy int    `gorm:"not null;" json:"created_by"`
	UpdatedBy int    `gorm:"not null;" json:"updated_by"`
}
