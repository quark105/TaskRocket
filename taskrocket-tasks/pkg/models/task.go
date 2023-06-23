package models

import "time"

type Task struct {
	Id          int64     `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      int8      `json:"status"`
	DueDate     time.Time `json:"due_date"`
}
