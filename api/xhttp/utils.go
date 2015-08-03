package xhttp

import (
	"net/http"

	gorillaContext "github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

type Composer struct {
	middlewares []func(http.Handler) http.Handler
}

func Compose(middlewares ...func(http.Handler) http.Handler) Composer {
	c := Composer{
		middlewares: middlewares,
	}

	return c
}

func (c Composer) Then(h http.Handler) http.Handler {
	if h == nil {
		h = http.DefaultServeMux
	}

	for i := len(c.middlewares) - 1; i >= 0; i-- {
		h = c.middlewares[i](h)
	}

	return h
}

func Adapt(handler http.Handler) httprouter.Handle {

	return httprouter.Handle(
		func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
			ctx := &Context{
				Params: params,
			}
			gorillaContext.Set(r, keyContext, ctx)

			handler.ServeHTTP(w, r)
		})
}
