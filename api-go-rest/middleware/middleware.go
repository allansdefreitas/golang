package middleware

import "net/http"

func SetContentTypeMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json") // insert conten-type
		next.ServeHTTP(w, r)                               //put this configurarion to all handlers
	})
}
