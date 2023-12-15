package repositories

import "gorm.io/gorm"

type AuthRepository interface {
	SaveGithubAccessToken(userID uint, accessToken string) error
	GetAccessTokenByUserID(userID uint) (string, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) SaveGithubAccessToken(userID uint, accessToken string) error {
	return r.db.Exec("INSERT INTO oauth_tokens (user_id, access_token) VALUES (?, ?)", userID, accessToken).Error
}

func (r *authRepository) GetAccessTokenByUserID(userID uint) (string, error) {
	var accessToken string
	err := r.db.Raw("SELECT access_token FROM oauth_tokens WHERE user_id = ? ORDER BY id DESC LIMIT 1", userID).
		Scan(&accessToken).Error
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
