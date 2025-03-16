package service

import (
	"errors"
	"github.com/kriswu/go_deepseek/internal/app/dao"
	"github.com/kriswu/go_deepseek/internal/app/model"
	"golang.org/x/crypto/bcrypt"
)

// UserService 用户服务
type UserService struct {
	userDAO *dao.UserDAO
}

// NewUserService 创建用户服务实例
func NewUserService(userDAO *dao.UserDAO) *UserService {
	return &UserService{userDAO: userDAO}
}

// Register 用户注册
func (s *UserService) Register(user *model.User) error {
	// 检查用户名是否已存在
	existingUser, err := s.userDAO.GetByUsername(user.Username)
	if err == nil && existingUser != nil {
		return errors.New("用户名已存在")
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// 创建用户
	return s.userDAO.Create(user)
}

// Login 用户登录
func (s *UserService) Login(username, password string) (*model.User, error) {
	// 获取用户信息
	user, err := s.userDAO.GetByUsername(username)
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	return user, nil
}

// GetUserByID 根据ID获取用户信息
func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	return s.userDAO.GetByID(id)
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(user *model.User) error {
	return s.userDAO.Update(user)
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(id uint) error {
	return s.userDAO.Delete(id)
}

// ListUsers 获取用户列表
func (s *UserService) ListUsers(page, pageSize int) ([]model.User, error) {
	offset := (page - 1) * pageSize
	return s.userDAO.List(offset, pageSize)
}