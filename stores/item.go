package stores

import "github.com/golang-vietnam/grox/api/rethink"

type ItemStore struct {
	re *rethink.Instance
}

func NewItemStore(re *rethink.Instance) *ItemStore {
	return &ItemStore{
		re: re,
	}
}
