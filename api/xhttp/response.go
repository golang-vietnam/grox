package xhttp

import (
	"encoding/json"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, status int, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(status)
	w.Write(data)
}
