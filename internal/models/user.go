package models

import (
	"database/sql"

	"gorm.io/gorm"
)

// Model reference :
// 'reg_num', 'profile_picture', 'password', 'level_id', 'created_at',

type User struct {
	gorm.Model
	Username       sql.NullString `json:"username" gorm:"not null;unique;type:varchar(25)"`
	ProfilePicture string         `json:"profile_picture" gorm:"default:'default.png'"`
	Password       sql.NullString `json:"password" gorm:"not null;type:varchar(18)"`
	Level          int            `json:"level" gorm:"default:0"`
}
