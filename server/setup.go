package server

import (
	"net/http"

	"github.com/dancannon/gorethink"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"

	"github.com/golang-vietnam/grox/api/rethink"
	"github.com/golang-vietnam/grox/api/xhttp"
	"github.com/golang-vietnam/grox/dbscript"
	"github.com/golang-vietnam/grox/handlers"
	"github.com/golang-vietnam/grox/middlewares"
	"github.com/golang-vietnam/grox/stores"
	"github.com/golang-vietnam/grox/utils/logs"
)

var l = logs.New("grox/server")

type setupStruct struct {
	Config

	Rethink *rethink.Instance
	Handler http.Handler
}

func setup(cfg Config) *setupStruct {
	s := &setupStruct{Config: cfg}
	s.setupRethink()
	s.setupScript()
	s.setupRoutes()

	return s
}

func (s *setupStruct) setupRethink() {
	cfg := s.Config
	re, err := rethink.NewInstance(gorethink.ConnectOpts{
		Address:  cfg.RethinkDB.Addr + ":" + cfg.RethinkDB.Port,
		Database: cfg.RethinkDB.DBName,
	})

	if err != nil {
		l.Fatalln("Could not connect to RethinkDB")
	}

	s.Rethink = re
}

func (s *setupStruct) setupScript() {
	script := dbscript.NewRethinkScript(s.Rethink)
	if err := script.Settup(); err != nil {
		l.Fatalln("Could generate Table to RethinkDB")
	}
}

func commonMiddlewares() func(http.Handler) http.Handler {
	logger := middlewares.NewLogger()
	recovery := middlewares.NewRecovery()

	return func(h http.Handler) http.Handler {
		return recovery(logger(h))
	}
}

func authMiddlewares() func(http.Handler) http.Handler {
	auth := middlewares.NewAuth()

	return func(h http.Handler) http.Handler {
		return auth(h)
	}
}

func (s *setupStruct) setupRoutes() {
	commonMids := commonMiddlewares()
	authMids := authMiddlewares()

	normal := func(h http.HandlerFunc) httprouter.Handle {
		return xhttp.Adapt(commonMids(h))
	}

	auth := func(h http.HandlerFunc) httprouter.Handle {
		return xhttp.Adapt(commonMids(authMids(h)))
	}

	router := httprouter.New()
	itemStore := stores.NewItemStore(s.Rethink)
	userStore := stores.NewUserStore(s.Rethink)
	storyStore := stores.NewStoryStore(s.Rethink)

	{
		itemCtrl := handlers.NewItemCtrl(itemStore)
		router.GET("/v1/item", normal(itemCtrl.List))
		router.GET("/v1/item/:id", normal(itemCtrl.Get))
		router.POST("/v1/item", auth(itemCtrl.Create))
	}

	{
		userCtrl := handlers.NewUserCtrl(userStore)
		router.GET("/v1/user", normal(userCtrl.List))
		router.GET("/v1/user/:id", normal(userCtrl.Get))
		router.POST("/v1/user", normal(userCtrl.Create))
	}

	{
		storyHandler := handlers.NewStoryHandler(storyStore)
		router.GET("/v1/story", normal(storyHandler.List))
		router.GET("/v1/story/:id", normal(storyHandler.Get))
		router.POST("/v1/story", normal(storyHandler.Create))
	}

	s.Handler = context.ClearHandler(router)
}
