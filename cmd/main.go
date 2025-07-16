package main

import (
	"TestProjecct/internal/db"
	"TestProjecct/internal/handlers"
	"TestProjecct/internal/taskService"
	userservice "TestProjecct/internal/userService"
	"TestProjecct/internal/web/tasks"
	"TestProjecct/internal/web/users"
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
	userRepo := userservice.NewUserRepository(database)
	userService := userservice.NewUserService(userRepo)
	userHandlers := handlers.NewUserHandler(userService)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	strictTaskHandler := tasks.NewStrictHandler(taskHandlers, nil)
	tasks.RegisterHandlers(e, strictTaskHandler)

	strictUserHandler := users.NewStrictHandler(userHandlers, nil)
	users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}
