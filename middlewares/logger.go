package middlewares

import (
	"log"
	"net/http"
	"time"

	"github.com/golang-vietnam/grox/utils/logs"
)

type Logger struct {
	*log.Logger
}

func NewLogger() func(http.Handler) http.Handler {
	l := Logger{logs.New("HTTP")}
	return l.factory
}

func (l Logger) factory(next http.Handler) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			l.Printf("Started %v %v", r.Method, r.URL.Path)

			next.ServeHTTP(w, r)

			l.Printf("Completed %v %v in %v", r.Method, r.URL.Path, time.Since(start))
		})
}
