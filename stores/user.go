package stores

import "github.com/golang-vietnam/grox/api/rethink"

type UserStore struct {
	re *rethink.Instance
}

func NewUserStore(re *rethink.Instance) *UserStore {
	return &UserStore{
		re: re,
	}
}
