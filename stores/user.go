package stores

import (
	"errors"
	"github.com/golang-vietnam/grox/api/rethink"
	"github.com/golang-vietnam/grox/domain"
	"gopkg.in/bluesuncorp/validator.v8"
)

const (
	UserTable = "user"
)

type UserStore struct {
	re       *rethink.Instance
	validate *validator.Validate
}

func NewUserStore(re *rethink.Instance, validate *validator.Validate) *UserStore {
	return &UserStore{
		re:       re,
		validate: validate,
	}
}

func (this *UserStore) Create(user *domain.User) (err error) {
	_, err = this.re.RunWrite(
		this.re.Table(UserTable).Filter(map[string]interface{}{
			"username": user.Username}).Insert(user))
	return
}

func (this *UserStore) Validate(user *domain.User) []error {
	var errorList []error
	errs := this.validate.Struct(user)
	for _, v := range errs.(validator.ValidationErrors) {
		switch v.Field {
		case "Username":
			switch v.Tag {
			case "required":
				errorList = append(errorList, errors.New("Username required"))
			case "max":
				errorList = append(errorList, errors.New("Username max length 20 charater"))
			}
		}
	}
	return errorList
}
