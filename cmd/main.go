package main

import (
	"TestProjecct/internal/db"
	"TestProjecct/internal/handlers"
	"TestProjecct/internal/taskService"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	e := echo.New()

	taskRepo := taskService.NewTaskRepository(database)
	taskService := taskService.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(taskService)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/tasks", taskHandlers.GetTask)
	e.POST("/tasks", taskHandlers.PostTask)
	e.PATCH("/tasks/:id", taskHandlers.PatchTask)
	e.DELETE("/tasks/:id", taskHandlers.DeleteTask)

	e.Start("localhost:8080")
}
