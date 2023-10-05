package repositories

import (
	"errors"
	"fmt"

	"github.com/fseda/cookbooked-api/internal/domain/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindOneById(id uint) (*models.User, error)
	FindOneBy(field string, value string) (*models.User, error)
	Delete(id uint) (int64, error)
	UserExists(field string, value string) (bool, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

var searchFields = []string{"username", "email"}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindOneById(id uint) (*models.User, error) {
	var user models.User
	err := r.db.Omit("PasswordHash").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindOneBy(field string, value string) (*models.User, error) {
	if searchFieldIsValid := validateSearchField(field); !searchFieldIsValid {
		return nil, fmt.Errorf("invalid search field: %v, must be %v", field, searchFields)
	}

	var user models.User
	queryStr := fmt.Sprintf("%v = ?", field)
	err := r.db.Select(field).Where(queryStr, value).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Delete(id uint) (int64, error) {
	res := r.db.Delete(&models.User{}, id)
	return res.RowsAffected, res.Error
}

func (r *userRepository) UserExists(field string, value string) (bool, error) {
	if searchFieldIsValid := validateSearchField(field); !searchFieldIsValid {
		return false, fmt.Errorf("invalid search field: %v, must be %v", field, searchFields)
	}

	var user models.User
	queryStr := fmt.Sprintf("%v = ?", field)
	if err := r.db.Select(field).Where(queryStr, value).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func validateSearchField(field string) bool {
	for _, item := range searchFields {
		if item == field {
			return true
		}
	}
	return false
}
