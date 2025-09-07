package auth

import (
	"github.com/ardianilyas/go-auth-domain/internal/auth/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user *models.User) error
	FindUserByEmail(email string) (*models.User, error)
	FindUserByID(id uuid.UUID) (*models.User, error)
	CreateRefreshToken(token *models.RefreshToken) error
	FindRefreshToken(token string) (*models.RefreshToken, error)
	UpdateRefreshToken(rt *models.RefreshToken) error
	DeleteRefreshtoken(token string) error
	DeleteAllRefreshTokensByUser(userID uuid.UUID) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *authRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) FindUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) CreateRefreshToken(token *models.RefreshToken) error {
	return r.db.Create(token).Error
}

func (r *authRepository) FindRefreshToken(token string) (*models.RefreshToken, error) {
	var rt models.RefreshToken
	if err := r.db.Where("token = ? AND revoked = false", token).First(&rt).Error; err != nil {
		return nil, err
	}
	return &rt, nil
}

func (r *authRepository) UpdateRefreshToken(rt *models.RefreshToken) error {
	return r.db.Save(rt).Error
}

func (r *authRepository) DeleteRefreshtoken(token string) error {
	return r.db.Where("token = ?", token).Delete(&models.RefreshToken{}).Error
}

func (r *authRepository) DeleteAllRefreshTokensByUser(userID uuid.UUID) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.RefreshToken{}).Error
}