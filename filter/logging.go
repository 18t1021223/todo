package filter

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)

		// Log response
		duration := time.Since(start)
		zap.L().Info("Response",
			zap.String("m", r.Method),
			zap.String("p", r.URL.Path),
			zap.Duration("d", duration),
		)
	})
}
