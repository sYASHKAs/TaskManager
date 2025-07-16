package taskService

import "github.com/google/uuid"

type TaskService interface {
	CreateTask(req TaskRequest) (Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	UpdateTask(id string, req TaskRequest) (Task, error)
	DeleteTask(id string) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s *taskService) CreateTask(req TaskRequest) (Task, error) {
	task := Task{
		ID:     uuid.NewString(),
		Name:   req.Name,
		Status: req.Status,
	}
	if err := s.repo.CreateTask(task); err != nil {
		return Task{}, err
	}
	return task, nil
}

func (s *taskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *taskService) GetTaskByID(id string) (Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *taskService) UpdateTask(id string, req TaskRequest) (Task, error) {
	task, err := s.repo.GetTaskByID(id)
	if err != nil {
		return Task{}, err
	}

	task.Name = req.Name
	task.Status = req.Status

	if err := s.repo.UpdateTask(task); err != nil {
		return Task{}, err
	}
	return task, nil
}

func (s *taskService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}
