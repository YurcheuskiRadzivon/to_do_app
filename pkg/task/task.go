package task

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Notes     string `json:"notes"`
	Completed bool   `json:"completed"`
	Priority  int    `json:"priority"`
}
