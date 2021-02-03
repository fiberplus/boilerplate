package user

import (
	"boilerplate/models"

	"gorm.io/gorm"
)

// Repository
type Repository interface {
	Find(id int) (models.User, error)
}

//repository struct
type repository struct {
	db *gorm.DB
}

// NewRepo is the single instance repo that is being created.
func NewRepo(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

//Find - Find user repository
func (r *repository) Find(id int) (models.User, error) {
	var data models.User
	result := r.db.Where("id = ?", id).Find(&data)
	if result.Error != nil {
		return data, result.Error
	}
	return data, nil
}
