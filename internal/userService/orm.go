package userservice

import "TestProjecct/internal/taskService"

type User struct {
	ID       string             `gorm:"primaryKey" json:"id"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Tasks    []taskService.Task `gorm:"foreignKey:UserId"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
