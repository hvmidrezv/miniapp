package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hvmidrezv/miniapp/internal/handlers"
)

func SetupRoutes(router *gin.Engine, userHandler *handlers.UserHandler, taskHandler *handlers.TaskHandler) {
	api := router.Group("/api")
	{
		// User routes
		users := api.Group("/users")
		{
			users.GET("", userHandler.GetAllUsers)
			users.GET("/:id", userHandler.GetUserByID)
			users.POST("", userHandler.CreateUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}

		// Task routes
		tasks := api.Group("/tasks")
		{
			tasks.GET("", taskHandler.GetAllTasks)
			tasks.GET("/:id", taskHandler.GetTaskByID)
			tasks.POST("", taskHandler.CreateTask)
			tasks.PUT("/:id", taskHandler.UpdateTask)
			tasks.DELETE("/:id", taskHandler.DeleteTask)
		}
	}
}
