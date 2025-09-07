package auth

import (
	"errors"
	"time"

	"github.com/ardianilyas/go-auth-domain/internal/auth/models"
	"github.com/ardianilyas/go-auth-domain/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(username, email, password string) (*models.User, error)
	Login(email, password string) (*models.User, string, string, error)
	RefreshToken(oldToken string) (string, string, error)
	GetUserByID(userID uuid.UUID) (*models.User, error)
	Logout(token string) error
}

type authService struct {
	repo AuthRepository
	config *config.Config
}

func NewAuthService(repo AuthRepository, config *config.Config) AuthService {
	return &authService{repo, config}
}

func (s *authService) createAccessToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID.String(),
		"role": user.Role,
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.JWTSecret))
}

func (s *authService) Register(username, email, password string) (*models.User, error) {
	if _, err := s.repo.FindUserByEmail(email); err == nil {
		return nil, errors.New("email already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email: email,
		Username: username,
		Password: string(hashed),
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(email, password string) (*models.User, string, string, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return nil, "", "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", "", errors.New("invalid credentials")
	}

	accessToken, err := s.createAccessToken(user)
	if err != nil {
		return nil, "", "", err
	}

	refreshToken := uuid.NewString()
	rt := &models.RefreshToken{
		UserID: user.ID,
		Token: refreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
		CreatedAt: time.Now(),
	}

	if err := s.repo.CreateRefreshToken(rt); err != nil {
		return nil, "", "", err
	}

	return user, accessToken, refreshToken, nil
}

func (s *authService) RefreshToken(oldToken string) (string, string, error) {
	rt, err := s.repo.FindRefreshToken(oldToken)
	if err != nil {
		return "", "", errors.New("invalid token")
	}

	if rt.Revoked {
		return "", "", errors.New("token is revoked")
	}

	if time.Now().After(rt.ExpiresAt) {
		return "", "", errors.New("token is expired")
	}

	user, err := s.repo.FindUserByID(rt.UserID)
	if err != nil {
		return "", "", errors.New("user not found")
	}

	accessToken, err := s.createAccessToken(user)
	if err != nil {
		return "", "", err
	}

	newRefreshToken := uuid.NewString()
	newRT := &models.RefreshToken{
		UserID: user.ID,
		Token: newRefreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
		CreatedAt: time.Now(),
	}

	if err := s.repo.CreateRefreshToken(newRT); err != nil {
		return "", "", err
	}

	rt.Revoked = true
	if err := s.repo.UpdateRefreshToken(rt); err != nil {
		return "", "", err
	}

	return accessToken, newRefreshToken, nil
}

func (s *authService) GetUserByID(userID uuid.UUID) (*models.User, error) {
	user, err := s.repo.FindUserByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (s *authService) Logout(refreshToken string) error {
	rt, err := s.repo.FindRefreshToken(refreshToken)
	if err != nil {
		return errors.New("invalid token")
	}

	rt.Revoked = true
	return s.repo.UpdateRefreshToken(rt)
}