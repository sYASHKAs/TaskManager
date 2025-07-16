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
			UserId: &tsk.UserId,
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
		UserId: *taskRequest.UserId,
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
		UserId: &createdTask.UserId,
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
		UserId: *taskRequest.UserId,
	}
	updatedTask, err := h.service.UpdateTask(taskID, taskToUpdate)
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Name:   &updatedTask.Name,
		Status: &updatedTask.Status,
		UserId: &updatedTask.UserId,
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

// GetUsersIdTasks implements tasks.StrictServerInterface.
func (h *TaskHandler) GetUsersIdTasks(ctx context.Context, request tasks.GetUsersIdTasksRequestObject) (tasks.GetUsersIdTasksResponseObject, error) {
	userId := request.Id

	taskUser, err := h.service.GetTasksForUser(userId)
	if err != nil {
		return nil, err
	}
	response := tasks.GetUsersIdTasks200JSONResponse{}

	for _, tsk := range taskUser {
		task := tasks.Task{
			Id:     &tsk.ID,
			Name:   &tsk.Name,
			Status: &tsk.Status,
			UserId: &tsk.UserId,
		}
		response = append(response, task)
	}

	return response, nil
}
