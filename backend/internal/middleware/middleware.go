package middleware

import (
	"git.sample.ru/sample/internal/entity"
	"git.sample.ru/sample/internal/logger"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info.Printf("{'uri': '%s', 'X-Request-Id': '%s'}", r.RequestURI, r.Header.Get("x-request-id"))

		next.ServeHTTP(w, r)
	})
}

func WithRequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := string(r.Header.Get("x-request-id"))

		if requestId == "" {
			requestId = entity.GenerateID()
			logger.Info.Printf("X-Request-Id not set; New X-Request-Id value is %s", requestId)
		}

		w.Header().Add("x-request-id", requestId)

		next.ServeHTTP(w, r)
	})
}
