package table

import (
	"time"

	"github.com/lib/pq"
)

type User struct {
	ID          int    `db:"user_id"`
	Name        string `db:"full_name"`
	MSISDN      string
	Email       string      `db:"user_email"`
	BirthTime   pq.NullTime `db:"birth_date"`
	BirthDate   string
	CreateTime  time.Time `db:"create_time"`
	CreatedTime string
	UpdatedTime pq.NullTime `db:"update_time"`
	UpdateTime  string
	UserAge     int    `db:"user_age"`
	Calculation string `db:"-"`
}

func SearchUser(module *TableModule) []User {
	users := []User{}
	err := module.db.Select(&users, "SELECT user_id, full_name, msisdn, user_email, birth_date, create_time, update_time, COALESCE(EXTRACT(YEAR from AGE(birth_date)),0) AS user_age FROM WS_USER LIMIT 10;")
	if err != nil {
		panic(err)
	}
	data := users
	return data
}

func SearchUserByName(module *TableModule, searchName string) []User {
	users := []User{}
	err := module.db.Select(&users, "SELECT user_id, full_name, msisdn, user_email, birth_date, create_time, update_time, COALESCE(EXTRACT(YEAR from AGE(birth_date)),0) AS user_age FROM WS_USER WHERE full_name like $1 LIMIT 10;", searchName)
	if err != nil {
		panic(err)
	}
	data := users
	return data
}
