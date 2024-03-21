package service

import (
	"gin-gonic-gorm-boilerplate/internal/model"
	"gin-gonic-gorm-boilerplate/internal/repository"
)

// UserService 는 사용자 관련 서비스를 정의합니다.
type UserService struct {
	repo *repository.UserRepository
}

// NewUserService 는 새 UserService 인스턴스를 생성합니다.
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetAllUsers 는 모든 사용자를 조회합니다.
func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.repo.FindAll()
}

// CreateUser 는 새로운 사용자를 생성합니다.
func (s *UserService) CreateUser(user *model.User) error {
	return s.repo.Save(user)
}
