package handlers

import (
	"github.com/golang-vietnam/grox/api/xhttp"
	"github.com/golang-vietnam/grox/domain"
	"github.com/golang-vietnam/grox/stores"
	"net/http"
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
	var story domain.Story
	if err := xhttp.BindJSON(r, &story); err != nil {
		xhttp.ResponseJson(w, 500, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}
	xhttp.ResponseJson(w, 200, map[string]interface{}{
		"message": "ok",
	})
}

func (this *StoryHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (this *StoryHandler) Delete(w http.ResponseWriter, r *http.Request) {

}
