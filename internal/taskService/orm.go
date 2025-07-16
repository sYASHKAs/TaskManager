package taskService

type Task struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

type TaskRequest struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}
