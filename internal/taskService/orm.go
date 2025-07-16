package taskService

type Task struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
	UserId string `gorm:"type:uuid" json:"user_id"`
}

type TaskRequest struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
	UserId string `json:"user_id"`
}
