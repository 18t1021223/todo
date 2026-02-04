package util

import (
	"net/http"
	"strconv"
	"todo_list_roadmap/dto"
)

const (
	DefaultPage = 1
	DefaultSize = 20
	MaxSize     = 100
)

func Normalize(p *dto.PageRequest) {
	if p.Page <= 0 {
		p.Page = DefaultPage
	}
	if p.Limit <= 0 {
		p.Limit = DefaultSize
	}
	if p.Limit > MaxSize {
		p.Limit = MaxSize
	}
}

func Offset(p dto.PageRequest) int {
	return (p.Page - 1) * p.Limit
}

func BuildPageResponse[T any](
	items []T, totalItems int64, req dto.PageRequest,
) dto.PageResponse[T] {
	//totalPages := int(math.Ceil(float64(totalItems) / float64(req.Limit)))

	return dto.PageResponse[T]{
		Data:  items,
		Page:  req.Page,
		Limit: req.Limit,
		Total: totalItems,
	}
}

func ParsePageRequest(r *http.Request) dto.PageRequest {
	q := r.URL.Query()

	page, _ := strconv.Atoi(q.Get("page"))
	limit, _ := strconv.Atoi(q.Get("limit"))

	p := dto.PageRequest{Page: page, Limit: limit}
	Normalize(&p)
	return p
}
