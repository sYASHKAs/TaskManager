package handlers

import (
	"TestProjecct/internal/taskService"
	"TestProjecct/internal/web/tasks"
	"context"
)

type TaskHandler struct {
	service taskService.TaskService
}

func NewTaskHandler(s taskService.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

// GetTasks implements tasks.StrictServerInterface.
func (h *TaskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}
	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Name:   &tsk.Name,
			Status: &tsk.Status,
		}
		response = append(response, task)
	}

	return response, nil
}

// PostTasks implements tasks.StrictServerInterface.
func (h *TaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := taskService.TaskRequest{
		Name:   *taskRequest.Name,
		Status: *taskRequest.Status,
	}
	createdTask, err := h.service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Name:   &createdTask.Name,
		Status: &createdTask.Status,
	}
	// Просто возвращаем респонс!
	return response, nil
}

// PatchTasksId implements tasks.StrictServerInterface.
func (h *TaskHandler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskID := request.Id
	taskRequest := request.Body

	taskToUpdate := taskService.TaskRequest{
		Name:   *taskRequest.Name,
		Status: *taskRequest.Status,
	}
	updatedTask, err := h.service.UpdateTask(taskID, taskToUpdate)
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Name:   &updatedTask.Name,
		Status: &updatedTask.Status,
	}
	return response, nil
}

// DeleteTasksId implements tasks.StrictServerInterface.
func (h *TaskHandler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskId := request.Id

	if err := h.service.DeleteTask(taskId); err != nil {
		return nil, err
	}
	return tasks.DeleteTasksId204Response{}, nil
}

// func (h *TaskHandler) GetTask(c echo.Context) error {
// 	tasks, err := h.service.GetAllTasks()
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get task"})
// 	}
// 	return c.JSON(http.StatusOK, tasks)
// }

// func (h *TaskHandler) PostTask(c echo.Context) error {
// 	var req taskService.TaskRequest
// 	if err := c.Bind(&req); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
// 	}

// 	task, err := h.service.CreateTask(req)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not get task"})
// 	}
// 	return c.JSON(http.StatusCreated, task)
// }

// func (h *TaskHandler) PatchTask(c echo.Context) error {
// 	id := c.Param("id")

// 	var req taskService.TaskRequest
// 	if err := c.Bind(&req); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
// 	}

// 	updatedTask, err := h.service.UpdateTask(id, req)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not update task"})
// 	}
// 	return c.JSON(http.StatusOK, updatedTask)
// }

// func (h *TaskHandler) DeleteTask(c echo.Context) error {
// 	id := c.Param("id")

// 	if err := h.service.DeleteTask(id); err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete task"})
// 	}
// 	return c.NoContent(http.StatusNoContent)
// }
