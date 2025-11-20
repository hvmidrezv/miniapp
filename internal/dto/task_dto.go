package dto

type CreateTaskRequest struct {
	UserID uint   `json:"user_id" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Status string `json:"status" binding:"omitempty,oneof=pending done"`
}

type UpdateTaskRequest struct {
	Title  string `json:"title" binding:"omitempty"`
	Status string `json:"status" binding:"omitempty,oneof=pending done"`
}
