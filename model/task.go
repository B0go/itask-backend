package model

import "time"

//Task is the object that holds task data
type Task struct {
	Uuid        string    `json:"uuid" pg:"uuid" gorm:"primaryKey"`
	Title       string    `json:"title" pg:"title"`
	Creation    time.Time `json:"creation" pg:"creation"`
	DueDate     time.Time `json:"dueDate" pg:"due_date"`
	Description string    `json:"description" pg:"description"`
	Completed   bool      `json:"completed" pg:"completed"`
}
