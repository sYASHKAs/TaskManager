package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Task struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

type TaskRequest struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

var tasks = []Task{}

func getTask(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

func postTask(c echo.Context) error {
	var req TaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	task := Task{
		ID:     uuid.New().String(),
		Name:   req.Name,
		Status: req.Status,
	}

	tasks = append(tasks, task)
	return c.JSON(http.StatusCreated, task)
}

func patchTask(c echo.Context) error {
	id := c.Param("id")

	var req TaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Name = req.Name
			tasks[i].Status = req.Status
			return c.JSON(http.StatusOK, tasks[i])
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
}

func main() {

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/tasks", getTask)
	e.POST("/tasks", postTask)
	e.PATCH("/tasks/:id", patchTask)
	e.DELETE("/tasks/:id", deleteTask)

	e.Start("localhost:8080")
}
