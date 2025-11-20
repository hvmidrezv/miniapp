package services

import (
	"errors"
	"math"

	"github.com/hvmidrezv/miniapp/internal/data/models"
	"github.com/hvmidrezv/miniapp/internal/dto"
	"github.com/hvmidrezv/miniapp/internal/repositories"
	"github.com/hvmidrezv/miniapp/pkg/utils"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetAllUsers(page, pageSize int) ([]models.User, utils.Pagination, error) {
	users, total, err := s.userRepo.FindAll(page, pageSize)
	if err != nil {
		return nil, utils.Pagination{}, err
	}

	totalPages := int64(math.Ceil(float64(total) / float64(pageSize)))

	pagination := utils.Pagination{
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
		TotalItems: total,
	}

	return users, pagination, nil
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

func (s *UserService) CreateUser(req *dto.CreateUserRequest) (*models.User, error) {
	// Check if email already exists
	existingUser, err := s.userRepo.FindByEmail(req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	user := &models.User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) UpdateUser(id uint, req *dto.UpdateUserRequest) (*models.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Check if email is being updated and if it already exists
	if req.Email != "" && req.Email != user.Email {
		existingUser, err := s.userRepo.FindByEmail(req.Email)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if existingUser != nil {
			return nil, errors.New("email already exists")
		}
		user.Email = req.Email
	}

	if req.Name != "" {
		user.Name = req.Name
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(id uint) error {
	_, err := s.userRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	return s.userRepo.Delete(id)
}
