package handlers

import (
	"net/http"

	"github.com/golang-vietnam/grox/stores"
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

}

func (this *UserCtrl) Update(w http.ResponseWriter, r *http.Request) {

}

func (this *UserCtrl) Delete(w http.ResponseWriter, r *http.Request) {

}
