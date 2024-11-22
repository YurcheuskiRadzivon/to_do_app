package model

import "time"

type Task struct {
	ID          int       `json:"id,omitempy"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	AddedTime   time.Time `json:"added_time,omitempy"`
	Images      []string  `json:"images,omitempy"`
}
