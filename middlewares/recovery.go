package middlewares

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/golang-vietnam/grox/utils/logs"
)

type Recovery struct {
	*log.Logger
}

func NewRecovery() func(http.Handler) http.Handler {
	r := Recovery{logs.New("Recovery")}
	return r.factory
}

func (r Recovery) factory(next http.Handler) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					debug.PrintStack()
				}
			}()

			next.ServeHTTP(w, r)
		})
}
