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

type TaskService struct {
	taskRepo *repositories.TaskRepository
	userRepo *repositories.UserRepository
}

func NewTaskService(taskRepo *repositories.TaskRepository, userRepo *repositories.UserRepository) *TaskService {
	return &TaskService{
		taskRepo: taskRepo,
		userRepo: userRepo,
	}
}

func (s *TaskService) GetAllTasks(page, pageSize int, status string) ([]models.Task, utils.Pagination, error) {
	tasks, total, err := s.taskRepo.FindAll(page, pageSize, status)
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

	return tasks, pagination, nil
}

func (s *TaskService) GetTaskByID(id uint) (*models.Task, error) {
	task, err := s.taskRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("task not found")
		}
		return nil, err
	}
	return task, nil
}

func (s *TaskService) CreateTask(req *dto.CreateTaskRequest) (*models.Task, error) {
	// Verify user exists
	_, err := s.userRepo.FindByID(req.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	status := req.Status
	if status == "" {
		status = "pending"
	}

	task := &models.Task{
		UserID: req.UserID,
		Title:  req.Title,
		Status: status,
	}

	if err := s.taskRepo.Create(task); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) UpdateTask(id uint, req *dto.UpdateTaskRequest) (*models.Task, error) {
	task, err := s.taskRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("task not found")
		}
		return nil, err
	}

	if req.Title != "" {
		task.Title = req.Title
	}

	if req.Status != "" {
		task.Status = req.Status
	}

	if err := s.taskRepo.Update(task); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) DeleteTask(id uint) error {
	_, err := s.taskRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("task not found")
		}
		return err
	}

	return s.taskRepo.Delete(id)
}
