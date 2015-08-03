package server

import "net/http"

type Config struct {
	Server struct {
		Port string `json:"API_PORT"`
		Addr string `json:"API_ADDR"`
	} `json:"server"`

	RethinkDB struct {
		Port   string `json:"RETHINKDB_PORT"`
		Addr   string `json:"RETHINKDB_ADDR"`
		DBName string `json:"RETHINKDB_DBNAME"`
	} `json:"rethinkdb"`
}

func Start(cfg Config) {
	s := setup(cfg)

	listenAddr := cfg.Server.Addr + ":" + cfg.Server.Port

	l.Println("grox-server is listening on", listenAddr)
	http.ListenAndServe(listenAddr, s.Handler)
}
