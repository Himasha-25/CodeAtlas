package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service struct {
	repo   *repository
	secret string
}

func NewService(db *gorm.DB, secret string) *Service {
	return &Service{repo: newRepository(db), secret: secret}
}

func (s *Service) Register(req RegisterRequest) (*AuthResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u := &User{Email: req.Email, PasswordHash: string(hash)}
	if err := s.repo.create(u); err != nil {
		return nil, err
	}
	token, err := generateToken(u.ID, s.secret)
	if err != nil {
		return nil, err
	}
	return &AuthResponse{Token: token}, nil
}

func (s *Service) Login(req LoginRequest) (*AuthResponse, error) {
	u, err := s.repo.findByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}
	token, err := generateToken(u.ID, s.secret)
	if err != nil {
		return nil, err
	}
	return &AuthResponse{Token: token}, nil
}
