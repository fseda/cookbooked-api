package repositories

import (
	"errors"
	"fmt"

	"github.com/fseda/cookbooked-api/internal/domain/models"
	"github.com/fseda/cookbooked-api/internal/domain/validator"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindOneById(id uint) (*models.User, error)
	FindOneBy(field string, value string) (*models.User, error)
	Delete(id uint) (int64, error)
	UserExists(field string, value string) (bool, error)
	FindOneForLogin(input string) (*models.User, error)
	FindOneByGithubID(githubID uint) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

var searchFields = []string{"username", "email"}

func (r *userRepository) Create(user *models.User) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if user.Email == "" {
		tx.Omit("email")
	}

	if user.Avatar == "" {
		tx.Omit("avatar_url")
	}

	if user.Name == "" {
		tx.Omit("name")
	}

	if user.Bio == "" {
		tx.Omit("bio")
	}

	if user.Location == "" {
		tx.Omit("location")
	}

	if user.GithubID == "" {
		tx.Omit("github_id")
	}

	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *userRepository) FindOneById(id uint) (*models.User, error) {
	var user models.User
	err := r.db.Omit("PasswordHash").First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
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
	if err := r.db.Where(queryStr, value).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindOneForLogin(input string) (*models.User, error) {
	var user *models.User
	var err error

	if validator.IsEmailLike(input) {
		user, err = r.FindOneBy("email", input)
		if user != nil && err == nil {
			return user, nil
		}
	} else {
		user, err := r.FindOneBy("username", input)
		if user != nil && err == nil {
			return user, nil
		}
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return nil, err
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

func (r *userRepository) FindOneByGithubID(githubID uint) (*models.User, error) {
	var user models.User
	if err := r.db.Where("github_id = ?", fmt.Sprint(githubID)).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func validateSearchField(field string) bool {
	for _, item := range searchFields {
		if item == field {
			return true
		}
	}
	return false
}
