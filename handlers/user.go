package handlers

import (
	"github.com/golang-vietnam/grox/api/xhttp"
	"github.com/golang-vietnam/grox/domain"
	"github.com/golang-vietnam/grox/stores"
	"net/http"
)

type UserCtrl struct {
	userStore *stores.UserStore
}

func NewUserCtrl(userStore *stores.UserStore) *UserCtrl {
	return &UserCtrl{
		userStore: userStore,
	}
}

func (this *UserCtrl) Get(w http.ResponseWriter, r *http.Request) {

}

func (this *UserCtrl) List(w http.ResponseWriter, r *http.Request) {

}

func (this *UserCtrl) Create(w http.ResponseWriter, r *http.Request) {
	var user = domain.User{}

	if err := xhttp.BindJSON(r, &user); err != nil {
		xhttp.ResponseJson(w, 500, M{
			"error": err.Error(),
		})
		return
	}

	if err := this.userStore.Create(&user); err != nil {
		xhttp.ResponseJson(w, 200, M{
			"error": err.Error(),
		})
		return
	}
	xhttp.ResponseJson(w, 200, user)

}

func (this *UserCtrl) Update(w http.ResponseWriter, r *http.Request) {

}

func (this *UserCtrl) Delete(w http.ResponseWriter, r *http.Request) {

}
