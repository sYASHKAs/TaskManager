package handlers

import (
	"TestProjecct/internal/taskService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service taskService.TaskService
}

func NewTaskHandler(s taskService.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) GetTask(c echo.Context) error {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get task"})
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) PostTask(c echo.Context) error {
	var req taskService.TaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	task, err := h.service.CreateTask(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not get task"})
	}
	return c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) PatchTask(c echo.Context) error {
	id := c.Param("id")

	var req taskService.TaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	updatedTask, err := h.service.UpdateTask(id, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not update task"})
	}
	return c.JSON(http.StatusOK, updatedTask)
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteTask(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete task"})
	}
	return c.NoContent(http.StatusNoContent)
}
