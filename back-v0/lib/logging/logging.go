
package logging

import (
    "log"
	"net/http"

)
type StatusRecorder struct {
	http.ResponseWriter
	Status int
	Bytes  int
}

func (rec *StatusRecorder) WriteHeader(status int) {
	rec.Status = status
	rec.ResponseWriter.WriteHeader(status)
}

func (rec *StatusRecorder) Write(b []byte) (int, error) {
	n, err := rec.ResponseWriter.Write(b)
	rec.Bytes += n
	return n, err
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &StatusRecorder{ResponseWriter: w, Status: http.StatusOK}

		next.ServeHTTP(rec, r)

		log.Printf("[%s] %s %d",
			r.Method,
			r.URL.Path,
			rec.Status,
		)
	})
}
