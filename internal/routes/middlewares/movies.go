package middlewares

import "net/http"

// MovieCtx TODO add context getter and Value
func MovieCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

	})
}
