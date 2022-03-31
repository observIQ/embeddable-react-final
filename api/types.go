package api

type AllTodos map[string]todo

type ListResponse struct {
	Todos []todo `json:"todos"`
}

type CreatePayload struct {
	Description string `json:"description"`
}

type CreateResponse struct {
	Todo todo `json:"todo"`
}

type CheckPayload struct {
	Completed bool `json:"completed"`
}

type CheckResponse struct {
	Todo todo `json:"todo"`
}
