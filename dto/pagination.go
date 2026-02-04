package dto

type PageRequest struct {
	Page  int
	Limit int
}

type PageResponse[T any] struct {
	Data  []T   `json:"data"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Total int64 `json:"total"`
}
