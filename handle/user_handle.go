package handle

import (
	"net/http"
	"todo_list_roadmap/dto"
	"todo_list_roadmap/handle/response"
	"todo_list_roadmap/service"
	"todo_list_roadmap/util"

	"github.com/go-chi/chi/v5"
)

type userHandle struct {
	s *service.UserService
}

func RegisterUserRouter(s *service.UserService, r chi.Router) {
	handle := &userHandle{s}
	r.Post("/register", handle.register)
	r.Post("/login", handle.login)
}

func (h *userHandle) register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req dto.UserRegisterRequest
	if err := util.BindAndValidate(r, &req); err != nil {
		response.ErrValidation(w, err)
		return
	}

	responseDTO, err := h.s.Create(ctx, req)
	if err != nil {
		response.ErrorJSON(w, err)
		return
	}
	response.Created(w, responseDTO)
}

func (h *userHandle) login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req dto.UserLoginRequest
	if err := util.BindAndValidate(r, &req); err != nil {
		response.ErrValidation(w, err)
		return
	}

	responseDTO, err := h.s.Login(ctx, req)
	if err != nil {
		response.ErrorJSON(w, err)
		return
	}
	response.OK(w, responseDTO)
}
