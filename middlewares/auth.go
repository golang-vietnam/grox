package middlewares

import (
	"fmt"
	"net/http"
)

type Auth struct {
}

func NewAuth() func(http.Handler) http.Handler {
	a := Auth{}
	return a.factory
}

func (a Auth) factory(next http.Handler) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(403)
			fmt.Fprintf(w, "You are not allowed to access this page.")
		})
}
