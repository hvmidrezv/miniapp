package services

import (
	"errors"

	"github.com/hvmidrezv/miniapp/internal/data/models"
	"github.com/hvmidrezv/miniapp/internal/dto"
	"github.com/hvmidrezv/miniapp/internal/repositories"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetAllUsers(page, pageSize int) ([]models.User, int64, error) {
	return s.userRepo.FindAll(page, pageSize)
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *UserService) CreateUser(req dto.CreateUserRequest) (*models.User, error) {
	exists, err := s.userRepo.ExistsByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
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

func (s *UserService) UpdateUser(id uint, req dto.UpdateUserRequest) (*models.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		exists, err := s.userRepo.ExistsByEmailExcludingID(req.Email, id)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, errors.New("email already exists")
		}
		user.Email = req.Email
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
