package handlers

import (
	"net/http"

	"github.com/golang-vietnam/grox/api/xhttp"
	"github.com/golang-vietnam/grox/stores"
)

type ItemCtrl struct {
	itemStore *stores.ItemStore
}

func NewItemCtrl(itemStore *stores.ItemStore) *ItemCtrl {
	return &ItemCtrl{
		itemStore: itemStore,
	}
}

func (this *ItemCtrl) Get(w http.ResponseWriter, r *http.Request) {
	ctx := xhttp.GetContext(r)
	id := ctx.Params.ByName("id")

	xhttp.ResponseJson(w, 200, map[string]interface{}{
		"id": id,
	})
}

func (this *ItemCtrl) List(w http.ResponseWriter, r *http.Request) {

}

func (this *ItemCtrl) Create(w http.ResponseWriter, r *http.Request) {

}

func (this *ItemCtrl) Update(w http.ResponseWriter, r *http.Request) {

}

func (this *ItemCtrl) Delete(w http.ResponseWriter, r *http.Request) {

}
