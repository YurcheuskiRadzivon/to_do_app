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

type TaskH struct {
	ID          int       `json:"id,omitempy"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	AddedTime   time.Time `json:"added_time,omitempy"`
	Images      []byte    `json:"images,omitempy"`
	UserId      int       `json:"user_id,omitempy"`
}
type T struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
