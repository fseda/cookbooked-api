package repositories

import (
	"fmt"

	"github.com/fseda/cookbooked-api/internal/domain/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) create(username, email, passwordHash string) *models.User {
	user := &models.User{
		// Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
	}

	result := r.db.Create(user)
	
	fmt.Println(result.RowsAffected)
	fmt.Println("user id", user.ID)

	return user
}
