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
	user, err := this.userStore.Get(xhttp.GetContext(r).Params.ByName("id"))
	if err != nil {
		xhttp.ResponseJson(w, 500, M{
			"error": err.Error(),
		})
		return
	}
	xhttp.ResponseJson(w, 200, user)
}

func (this *UserCtrl) List(w http.ResponseWriter, r *http.Request) {
	users, err := this.userStore.List()
	if err != nil {
		xhttp.ResponseJson(w, 500, M{
			"error": err.Error(),
		})
		return
	}
	xhttp.ResponseJson(w, 200, users)
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
	type updateUserData struct {
		Name string `json:"name",gorethink:"name"`
	}
	var data updateUserData
	if err := xhttp.BindJSON(r, &data); err != nil {
		xhttp.ResponseJson(w, 500, M{
			"error": err.Error(),
		})
		return
	}

	if err := this.userStore.Update(xhttp.GetContext(r).Params.ByName("id"), M{"name": data.Name}); err != nil {
		xhttp.ResponseJson(w, 500, M{
			"error": err.Error(),
		})
		return
	}
	xhttp.ResponseJson(w, 200, M{"message": "updated"})
}

func (this *UserCtrl) Delete(w http.ResponseWriter, r *http.Request) {
	if err := this.userStore.Delete(xhttp.GetContext(r).Params.ByName("id")); err != nil {
		xhttp.ResponseJson(w, 500, M{
			"error": err.Error(),
		})
		return
	}
	xhttp.ResponseJson(w, 200, M{"message": "deleted"})
}
