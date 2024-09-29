package repository

import (
	"my-project/cmd/auth/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	FindByID(id uint) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) (*models.User, error) {
	result := r.db.Create(user)
	return user, result.Error
}

func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	return &user, result.Error
}

func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	result := r.db.Where("username = ?", username).First(&user)
	return &user, result.Error
}

func (r *userRepository) Update(user *models.User) (*models.User, error) {
	result := r.db.Save(user)
	return user, result.Error
}

func (r *userRepository) Delete(id uint) error {
	result := r.db.Delete(&models.User{}, id)
	return result.Error
}