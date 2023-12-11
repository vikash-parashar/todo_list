package models

import "time"

type Todo struct {
	ID            int       `json:""`
	Title         string    ``
	CurrentStatus string    ``
	CreatedAt     time.Time ``
	UpdatedAt     time.Time ``
	DeleteAt      time.Time ``
	DueDate       time.Time ``
}

const (
	StatusComplete = "Complete"
	StatusPending  = "Pending"
)
