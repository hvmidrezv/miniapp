package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hvmidrezv/miniapp/internal/dto"
	"github.com/hvmidrezv/miniapp/internal/services"
	"github.com/hvmidrezv/miniapp/pkg/utils"
)

type TaskHandler struct {
	taskService *services.TaskService
}

func NewTaskHandler(taskService *services.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

// GetAllTasks godoc
// @Summary List all tasks with pagination and filtering
// @Description Get a paginated list of all tasks with optional status filter
// @Tags Tasks
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Param status query string false "Filter by status" Enums(pending, done)
// @Success 200 {object} utils.PaginatedResponse
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/tasks [get]
func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	page, pageSize := utils.GetPaginationParams(c)
	status := c.Query("status")

	// Validate status if provided
	if status != "" && status != "pending" && status != "done" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid status. Must be 'pending' or 'done'")
		return
	}

	tasks, pagination, err := h.taskService.GetAllTasks(page, pageSize, status)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.PaginatedSuccessResponse(c, http.StatusOK, tasks, pagination)
}

// GetTaskByID godoc
// @Summary Get single task
// @Description Get a specific task by ID
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /api/tasks/{id} [get]
func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid task ID")
		return
	}

	task, err := h.taskService.GetTaskByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Task retrieved successfully", task)
}

// CreateTask godoc
// @Summary Create task for a user
// @Description Create a new task with title and user_id validation
// @Tags Tasks
// @Accept json
// @Produce json
// @Param task body dto.CreateTaskRequest true "Task data"
// @Success 201 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/tasks [post]
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req dto.CreateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, utils.FormatValidationError(err))
		return
	}

	task, err := h.taskService.CreateTask(&req)
	if err != nil {
		if err.Error() == "user not found" {
			utils.ErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Task created successfully", task)
}

// UpdateTask godoc
// @Summary Update task title or status
// @Description Update task's title and/or status (pending or done)
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param task body dto.UpdateTaskRequest true "Task data"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/tasks/{id} [put]
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

	task, err := h.taskService.UpdateTask(uint(id), &req)
	if err != nil {
		if err.Error() == "task not found" {
			utils.ErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Task updated successfully", task)
}

// DeleteTask godoc
// @Summary Delete task
// @Description Delete a specific task by ID
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/tasks/{id} [delete]
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
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Task deleted successfully", nil)
}
