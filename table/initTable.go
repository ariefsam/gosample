package table

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"
)

type TableModule struct {
	Session   string
	db        *sqlx.DB
	redispool *redis.Pool
}

type TemplateData struct {
	Data    []User
	Visitor int64
}

func NewTableModule() TableModule {

	db, err := sqlx.Connect("postgres", "postgres://da161205:123Toped456@devel-postgre.tkpd/tokopedia-user?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	redispool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "devel-redis.tkpd:6379")
			if err != nil {
				return nil, err
			}
			return conn, err
		},
	}
	t := TableModule{"Sessid", db, redispool}
	return t
}

func (tableModule *TableModule) ShowData(w http.ResponseWriter, r *http.Request) {
	redispool := tableModule.redispool.Get()
	// Close redis after function end
	defer redispool.Close()
	visitor, err := redis.Int64(redispool.Do("INCR", "visitors"))

	searchName := "%" + r.FormValue("name") + "%"
	var dataUser = []User{}
	if searchName == "%%" {
		dataUser = SearchUser(tableModule)
	} else {
		dataUser = SearchUserByName(tableModule, searchName)
	}

	dataTable := TemplateData{dataUser, visitor}
	tmpl, err := template.ParseFiles("./files/htmltemplate/table.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, dataTable)
	if err != nil {
		panic(err)
	}

}
