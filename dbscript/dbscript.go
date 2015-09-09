package dbscript

import (
	"github.com/golang-vietnam/grox/api/rethink"
	"github.com/golang-vietnam/grox/stores"
)

type RethinkScript struct {
	rt *rethink.Instance
}

func NewRethinkScript(rt *rethink.Instance) *RethinkScript {
	return &RethinkScript{
		rt: rt,
	}
}
func (this *RethinkScript) Settup() error {
	if err := this.createTableUser(); err != nil {
		return err
	}
	return nil
}

func (this *RethinkScript) createTableUser() error {
	this.rt.Exec(this.rt.DB().TableCreate(stores.UserTable))
	if _, err := this.rt.Run(this.rt.Table(stores.UserTable).IndexStatus("username")); err != nil {
		if err := this.rt.Exec(this.rt.Table(stores.UserTable).IndexCreate("username")); err != nil {
			return err
		}
		if err := this.rt.Exec(this.rt.Table(stores.UserTable).IndexWait()); err != nil {
			return err
		}
	}
	return nil
}
