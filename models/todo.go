package models

import "time"

type Todo struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
}
