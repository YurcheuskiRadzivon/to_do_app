package task

//easyjson:json
type (
	task struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Notes     string `json:"notes"`
		Completed bool   `json:"completed"`
		Priority  int    `json:"priority"`
	}

	Tasks struct {
		arr []task `json:"arr"`
	}
)
