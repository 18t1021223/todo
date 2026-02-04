package filter

import (
	"net/http"
	"todo_list_roadmap/handle/response"

	"github.com/go-chi/jwtauth/v5"
)

func JWTAuthenticator(next http.Handler) http.Handler {
	hfn := func(w http.ResponseWriter, r *http.Request) {
		token, _, err := jwtauth.FromContext(r.Context())

		if err != nil {
			response.ErrorJSON(w, response.ErrUnauthorized)
			return
		}

		if token == nil {
			response.ErrorJSON(w, response.ErrUnauthorized)
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(hfn)
}
