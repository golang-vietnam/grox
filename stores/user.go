package stores

import (
	"errors"
	"github.com/golang-vietnam/grox/api/rethink"
	"github.com/golang-vietnam/grox/domain"
	"github.com/golang-vietnam/grox/utils/validator"
	"time"
)

type UserStore struct {
	re *rethink.Instance
}

func NewUserStore(re *rethink.Instance) *UserStore {
	return &UserStore{
		re: re,
	}
}

func (this *UserStore) Create(user *domain.User) error {
	if cussor, err := this.re.Run(this.re.Table(UserTable).Filter(M{"username": user.Username})); err != nil {
		return err
	} else {
		if cussor.IsNil() {
			if errs := validator.Validate(user); len(errs) > 0 {
				return errs[0]
			}
			user.CreatedTime = time.Now()
			user.Password = HashPassword(user.Password)
			ws, err := this.re.RunWrite(
				this.re.Table(UserTable).Insert(user))
			if err != nil {
				return err
			}
			if len(ws.GeneratedKeys) > 0 {
				user.Id = ws.GeneratedKeys[0]
			}
			return nil
		} else {
			return errors.New("User exist")
		}

	}
}
