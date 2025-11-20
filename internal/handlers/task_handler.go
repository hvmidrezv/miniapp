package handlers

import (
	"errors"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hvmidrezv/miniapp/internal/dto"
	"github.com/hvmidrezv/miniapp/internal/services"
	"github.com/hvmidrezv/miniapp/pkg/utils"
	"gorm.io/gorm"
)

type TaskHandler struct {
	taskService *services.TaskService
}

func NewTaskHandler(taskService *services.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	page, pageSize := utils.GetPaginationParams(c)

	tasks, total, err := h.taskService.GetAllTasks(page, pageSize)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch tasks")
		return
	}

	totalPages := int64(math.Ceil(float64(total) / float64(pageSize)))

	pagination := utils.Pagination{
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
		TotalItems: total,
	}

	utils.PaginatedSuccessResponse(c, http.StatusOK, tasks, pagination)
}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid task ID")
		return
	}

	task, err := h.taskService.GetTaskByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.ErrorResponse(c, http.StatusNotFound, "Task not found")
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch task")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Task retrieved successfully", task)
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req dto.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, utils.FormatValidationError(err))
		return
	}

	task, err := h.taskService.CreateTask(req)
	if err != nil {
		if err.Error() == "user not found" {
			utils.ErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create task")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Task created successfully", task)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid task ID")
		return
	}

	var req dto.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, utils.FormatValidationError(err))
		return
	}

	task, err := h.taskService.UpdateTask(uint(id), req)
	if err != nil {
		if err.Error() == "task not found" {
			utils.ErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update task")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Task updated successfully", task)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid task ID")
		return
	}

	if err := h.taskService.DeleteTask(uint(id)); err != nil {
		if err.Error() == "task not found" {
			utils.ErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete task")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Task deleted successfully", nil)
}
