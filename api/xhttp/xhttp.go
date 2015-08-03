package xhttp

import (
	"net/http"

	gorillaContext "github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

type keyContextType int

const keyContext keyContextType = 0

type Context struct {
	Params httprouter.Params
}

func GetContext(r *http.Request) *Context {
	ctx := gorillaContext.Get(r, keyContext).(*Context)
	return ctx
}
