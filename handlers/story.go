package handlers

import (
	"net/http"

	"github.com/golang-vietnam/grox/api/xhttp"
	"github.com/golang-vietnam/grox/stores"
)

type StoryHandler struct {
	storyStore *stores.StoryStore
}

func NewStoryHandler(storyStore *stores.StoryStore) *StoryHandler {
	return &StoryHandler{
		storyStore: storyStore,
	}
}

func (this *StoryHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := xhttp.GetContext(r)
	id := ctx.Params.ByName("id")
	xhttp.ResponseJson(w, 200, map[string]interface{}{
		"id": id,
	})
}

func (this *StoryHandler) List(w http.ResponseWriter, r *http.Request) {

}

func (this *StoryHandler) Create(w http.ResponseWriter, r *http.Request) {

}

func (this *StoryHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (this *StoryHandler) Delete(w http.ResponseWriter, r *http.Request) {

}
