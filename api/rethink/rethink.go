package rethink

import (
	r "github.com/dancannon/gorethink"

	"github.com/golang-vietnam/grox/utils/logs"
)

var (
	l = logs.New("api/rethink")
)

type Instance struct {
	opts    r.ConnectOpts
	session *r.Session

	db string
}

func NewInstance(opts r.ConnectOpts) (ins *Instance, err error) {
	ins = &Instance{
		opts: opts,
		db:   opts.Database,
	}

	ins.session, err = r.Connect(opts)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (this *Instance) connect() (*r.Session, error) {
	session, err := r.Connect(this.opts)
	return session, err
}

func (this *Instance) DB() r.Term {
	return r.DB(this.db)
}

func (this *Instance) Exec(term r.Term) error {
	session, err := this.connect()
	if err != nil {
		return err
	}
	return term.Exec(session)
}

func (this *Instance) Table(name string) r.Term {
	return r.DB(this.db).Table(name)
}

func (this *Instance) Run(term r.Term) (*r.Cursor, error) {
	session, err := this.connect()
	if err != nil {
		return nil, err
	}

	return term.Run(session)
}

func (this *Instance) RunWrite(term r.Term) (r.WriteResponse, error) {
	session, err := this.connect()
	if err != nil {
		return r.WriteResponse{}, err
	}
	return term.RunWrite(session)
}

func (this *Instance) One(term r.Term, result interface{}) error {
	cursor, err := term.Run(this.session)
	if err != nil {
		return err
	}

	return cursor.One(result)
}

func (this *Instance) All(term r.Term, result interface{}) error {
	cursor, err := term.Run(this.session)
	if err != nil {
		return err
	}

	return cursor.All(result)
}

func (this *Instance) Between(table string, index interface{}, start, end interface{}) r.Term {
	return r.DB(this.db).Table(table).OrderBy(r.OrderByOpts{Index: index}).Between(start, end)
}

func (this *Instance) GetAll() {

}
