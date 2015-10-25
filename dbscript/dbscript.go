package dbscript

import (
	"strings"

	"github.com/dancannon/gorethink"
	"github.com/golang-vietnam/grox/api/rethink"
	"github.com/golang-vietnam/grox/stores"
	"github.com/golang-vietnam/grox/utils/logs"
)

var l = logs.New("dbscript")

type RethinkScript struct {
	re     *rethink.Instance
	dbname string
}

func NewRethinkScript(re *rethink.Instance, dbname string) *RethinkScript {
	if len(strings.TrimSpace(dbname)) == 0 {
		l.Fatalln("Empty database name")
	}
	return &RethinkScript{
		re:     re,
		dbname: dbname,
	}
}
func (this *RethinkScript) Setup() error {
	dbname := this.dbname
	err := this.re.Exec(gorethink.DB(dbname).Info())

	if err == nil {
		l.Printf("Database `%v` exists, skip creating.", dbname)
		return err
	}

	if !strings.Contains(err.Error(), "does not exist") {
		l.Printf("Error querying database `%v`: %v", dbname, err)
		return err
	}

	l.Printf("Creating database `%v`", dbname)
	err = this.re.Exec(gorethink.DBCreate(dbname))
	if err != nil {
		return err
	}

	err = this.createTableUser()
	return err
}

func (this *RethinkScript) createTableUser() error {
	err := this.re.Exec(this.re.DB().TableCreate(stores.UserTable))
	if err != nil {
		return err
	}

	l.Printf("Create index: username")
	err = this.re.Exec(this.re.Table(stores.UserTable).IndexCreate("username"))
	if err != nil {
		return err
	}

	err = this.re.Exec(this.re.Table(stores.UserTable).IndexWait())
	if err != nil {
		return err
	}

	return nil
}
