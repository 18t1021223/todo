package handle

import (
	"net/http"
	"todo_list_roadmap/config"
	"todo_list_roadmap/dto"
	"todo_list_roadmap/filter"
	"todo_list_roadmap/handle/response"
	"todo_list_roadmap/service"
	"todo_list_roadmap/util"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type todoHandle struct {
	s *service.TodoService
}

func RegisterTodoRouter(s *service.TodoService, r chi.Router) {
	handle := &todoHandle{s}
	r.Route("/todos", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(config.TokenAuth))
			r.Use(filter.JWTAuthenticator)

			r.Post("/", handle.create)
			r.Get("/", handle.get)
			r.Put("/{id}", handle.update)
			r.Delete("/{id}", handle.delete)
		})
	})
}

func (h *todoHandle) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req dto.TodoCreateRequest
	if err := util.BindAndValidate(r, &req); err != nil {
		response.ErrValidation(w, err)
		return
	}

	todo, err := h.s.Save(ctx, req)
	if err != nil {
		response.ErrorJSON(w, err)
		return
	}
	response.Created(w, todo)
}

func (h *todoHandle) get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pageReq := util.ParsePageRequest(r)
	todos, total, err := h.s.GetTodos(ctx, dto.TodoListRequest{PageRequest: pageReq})
	if err != nil {
		response.ErrorJSON(w, err)
		return
	}
	resp := util.BuildPageResponse(todos, total, pageReq)
	response.OK(w, resp)
}

func (h *todoHandle) update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ID := chi.URLParam(r, "id")
	if len(ID) == 0 {
		response.ErrValidation(w, response.ErrIDRequired)
		return
	}
	var request dto.TodoUpdateRequest
	if err := util.BindAndValidate(r, &request); err != nil {
		response.InvalidReq(w)
		return
	}
	todo, err := h.s.UpdateTodo(ctx, request, ID)
	if err != nil {
		response.ErrorJSON(w, err)
		return
	}
	response.OK(w, todo)
}

func (h *todoHandle) delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ID := chi.URLParam(r, "id")
	if len(ID) == 0 {
		response.ErrValidation(w, response.ErrIDRequired)
		return
	}
	if err := h.s.DeleteTodo(ctx, ID); err != nil {
		response.ErrorJSON(w, err)
		return
	}
	response.NoContent(w)
}
