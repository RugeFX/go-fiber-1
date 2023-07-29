package models

import (
	"time"

	"gorm.io/gorm"
)

// Model reference :
// 'id', 'class_id', 'subject_id', 'user_id', 'title', 'material', 'created_at',

type Material struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Title     string         `json:"title"`
	Body      string         `json:"body"`
}
