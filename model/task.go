package model

import "time"

//Task is the object that holds task data
type Task struct {
	Uuid        string    `json:"uuid"`
	Title       string    `json:"title"`
	Creation    time.Time `json:"creation"`
	DueDate     time.Time `json:"dueDate"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
}
