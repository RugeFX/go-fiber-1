package models

import "gorm.io/gorm"

// Model reference :
// 'id', 'class_id', 'subject_id', 'user_id', 'title', 'material', 'created_at',

type Material struct {
	gorm.Model
	Id    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
