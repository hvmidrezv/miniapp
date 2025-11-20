package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/hvmidrezv/miniapp/docs"
	"github.com/hvmidrezv/miniapp/internal/config"
	"github.com/hvmidrezv/miniapp/internal/handlers"
	"github.com/hvmidrezv/miniapp/internal/repositories"
	"github.com/hvmidrezv/miniapp/internal/routes"
	"github.com/hvmidrezv/miniapp/internal/services"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Mini App API
// @version 1.0
// @description REST API for User and Task Management
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host miniapplianassignment.liara.run
// @BasePath /
// @schemes https http

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

	// CORS configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * 3600,
	}))

	// setup routes
	routes.SetupRoutes(router, userHandler, taskHandler)

	// swagger docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// server
	port := ":8080"
	log.Printf("Server starting on port %s", port)
	if err := router.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
