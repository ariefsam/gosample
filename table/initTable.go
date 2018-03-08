package table

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type TableModule struct {
	Session string
	db      *sqlx.DB
}

type TemplateData struct {
	Data []User
}

func NewTableModule() TableModule {

	db, err := sqlx.Connect("postgres", "postgres://da161205:123Toped456@devel-postgre.tkpd/tokopedia-user?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	t := TableModule{"Sessid", db}
	return t
}

func (tableModule *TableModule) ShowData(w http.ResponseWriter, r *http.Request) {
	searchName := "%" + r.FormValue("name") + "%"
	var dataUser = []User{}
	if searchName == "%%" {
		dataUser = SearchUser(tableModule)
	} else {
		dataUser = SearchUserByName(tableModule, searchName)
	}
	dataTable := TemplateData{dataUser}
	tmpl, err := template.ParseFiles("../bigproject/files/htmltemplate/table.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, dataTable)
	if err != nil {
		panic(err)
	}

}
