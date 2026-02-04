package router

import (
	"database/sql"
	"todo_list_roadmap/handle"
	"todo_list_roadmap/service"

	"github.com/go-chi/chi/v5"
)

func RegisterRouters(r *chi.Mux, sqlDB *sql.DB) {
	// Service
	todoService := service.NewTodoService(sqlDB)
	userService := service.NewUserService(sqlDB)

	// Rest
	handle.RegisterTodoRouter(todoService, r)
	handle.RegisterUserRouter(userService, r)
}
