package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"todo_list_roadmap/config"
	"todo_list_roadmap/db/migrations"
	"todo_list_roadmap/filter"
	"todo_list_roadmap/router"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func main() {
	env := config.InitEnv()
	config.InitLogger(env.EnvProfile)
	config.InitValidation()
	db := config.ConnectDB(env, migrations.EmbedMigrations)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	config.InitJWT(env.JWTSecretKey, env.JwtAccessTokenExpirationSeconds)

	r := chi.NewRouter()

	// Middleware
	r.Use(filter.ConfigCors)
	r.Use(filter.LoggingMiddleware)
	r.Use(middleware.Recoverer)

	// Router
	router.RegisterRouters(r, db)

	zap.L().Info("Starting HTTP server",
		zap.String("server_url", fmt.Sprintf("%s:%d", "http://localhost", env.ServerPort)),
		zap.Int("pid", os.Getpid()),
	)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", env.ServerPort), r); err != nil {
		zap.L().Fatal("Server startup failed",
			zap.Error(err), zap.Int("address", env.ServerPort),
		)
	}
}
