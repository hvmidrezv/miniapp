package repositories

import (
	"github.com/hvmidrezv/miniapp/internal/data/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) FindAll(page, pageSize int, status string) ([]models.Task, int64, error) {
	var tasks []models.Task
	var total int64

	offset := (page - 1) * pageSize

	query := r.DB.Model(&models.Task{})

	// Apply status filter if provided
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// Count total records
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Fetch paginated records
	if err := query.Offset(offset).Limit(pageSize).Find(&tasks).Error; err != nil {
		return nil, 0, err
	}

	return tasks, total, nil
}

func (r *TaskRepository) FindByID(id uint) (*models.Task, error) {
	var task models.Task
	if err := r.DB.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) Create(task *models.Task) error {
	return r.DB.Create(task).Error
}

func (r *TaskRepository) Update(task *models.Task) error {
	return r.DB.Save(task).Error
}

func (r *TaskRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Task{}, id).Error
}
