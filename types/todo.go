package types

type TodoCreateDTO struct {
	Title string `json:"title"`
}

type TodoResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
