package filter

import (
	"net/http"

	"github.com/go-chi/cors"
)

func ConfigCors(next http.Handler) http.Handler {
	c := cors.New(
		cors.Options{
			AllowedOrigins:     []string{"https://*", "http://*"},
			AllowOriginFunc:    nil,
			AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:     nil,
			AllowCredentials:   false,
			MaxAge:             300,
			OptionsPassthrough: false,
			Debug:              false,
		},
	)
	return c.Handler(next)
}
