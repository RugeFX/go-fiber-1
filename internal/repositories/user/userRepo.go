package userRepo

import (
	"github.com/RugeFX/go-fiber-1.git/internal/models"
	"github.com/RugeFX/go-fiber-1.git/internal/repositories"
	"gorm.io/gorm"
)

type UserRepository interface {
	repositories.BaseRepository[models.User]
	// GetByID(id uuid.UUID) (models.User, error)
}

type repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *repo {
	return &repo{db}
}

func (r *repo) Save(u models.User) (models.User, error) {
	if err := r.db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}
