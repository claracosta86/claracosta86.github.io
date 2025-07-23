
package logging

import (
    "log"
	"net/http"
	"time"

)

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &StatusRecorder{ResponseWriter: w, Status: http.StatusOK}
		start := time.Now()

		next.ServeHTTP(rec, r)

		log.Printf("[%s] %s %d %s", r.Method, r.URL.Path, rec.Status, time.Since(start))
	})
}

func (rec *StatusRecorder) WriteHeader(code int) {
	rec.Status = code
	rec.ResponseWriter.WriteHeader(code)
}
