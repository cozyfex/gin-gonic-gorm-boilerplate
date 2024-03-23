package repository

import (
	"gin-gonic-gorm-boilerplate/internal/database"
	"gin-gonic-gorm-boilerplate/internal/model"
)

// UserRepository 는 사용자 정보에 대한 데이터베이스 작업을 관리합니다.
type UserRepository struct {
	db *database.Manager
}

// NewUserRepository 는 새 UserRepository 인스턴스를 생성합니다.
func NewUserRepository(db *database.Manager) *UserRepository {
	return &UserRepository{db: db}
}

// FindAll 은 모든 사용자를 조회합니다.
func (r *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Reader().Find(&users).Error
	return users, err
}

// Create 는 사용자 정보를 저장합니다.
func (r *UserRepository) Create(user *model.User) error {
	err := r.db.Writer().Create(user).Error
	return err
}

// User 는 대상 사용자 정보를 조회합니다.
func (r *UserRepository) User(email string) (model.User, error) {
	var user model.User
	err := r.db.Reader().First(&user, email).Error
	return user, err
}
