package main

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/sebastianwebber/app/web"
)

var (
	connStr = &pg.Options{
		User:     "seba",
		Password: "",
		Database: "seba",
	}
)

type dbLogger struct{}

func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {
	return
}

func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	fmt.Println(q.FormattedQuery())
	return
}

func main() {
	db := pg.Connect(connStr)
	db.AddQueryHook(dbLogger{})
	defer db.Close()

	r := web.GetRouter(db)
	r.Run()
}
