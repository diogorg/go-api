package middleware

import (
	"api/support"
	"net/http"
)

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		support.SetHeader(w)
		next.ServeHTTP(w, r)
	})
}
