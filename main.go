package main

import (
	_ "embed"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

//go:embed list_tables.sql
var listTablesSql string

const output = "output"

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db, err := sqlx.Open("sqlite", `M:\!DB\Сидродельня Леврана\2025-03-24\030000855477.db`)
	if err != nil {
		log.Fatal(err.Error())
	}
	check(err)
	defer db.Close()

	os.RemoveAll(output)
	os.MkdirAll(output, 0644)

	tables := []string{}
	err = db.Select(&tables, listTablesSql)
	check(err)
	for _, table := range tables {
		err := listColumns(db, table)
		check(err)
	}
}
