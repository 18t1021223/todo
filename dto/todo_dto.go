package dto

type TodoCreateRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

type TodoCreateResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

type TodoUpdateRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

type TodoUpdateResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

type TodoListRequest struct {
	PageRequest
}

type TodoDetailResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}
