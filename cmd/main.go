package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hvmidrezv/miniapp/internal/config"
	"github.com/hvmidrezv/miniapp/internal/handlers"
	"github.com/hvmidrezv/miniapp/internal/repositories"
	"github.com/hvmidrezv/miniapp/internal/routes"
	"github.com/hvmidrezv/miniapp/internal/services"
	"github.com/joho/godotenv"
)

func main() {
	// load .env params
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// connect to database
	config.ConnectDatabase()

	// initialize repositories
	userRepo := repositories.NewUserRepository(config.DB)
	taskRepo := repositories.NewTaskRepository(config.DB)

	// initialize services
	userService := services.NewUserService(userRepo)
	taskService := services.NewTaskService(taskRepo, userRepo)

	// initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	taskHandler := handlers.NewTaskHandler(taskService)

	//  gin router
	router := gin.Default()

	// setup routes
	routes.SetupRoutes(router, userHandler, taskHandler)

	// server
	port := ":8080"
	log.Printf("Server starting on port %s", port)
	if err := router.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
