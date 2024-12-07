package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/iancoleman/strcase"
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

//go:embed test.sql
var testSql string

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db, err := sqlx.Open("sqlite", `C:\!DB\UTSZ\030000527832.db`)
	check(err)
	defer db.Close()
	row, err := db.Query(testSql)
	check(err)
	f, err := os.Create("model.txt")
	check(err)
	defer f.Close()
	for row.Next() {
		if columns, err := row.Columns(); err == nil {
			_, err = fmt.Fprintln(f, "type model struct {")
			check(err)
			for _, col := range columns {
				log.Println(strcase.ToCamel(col))
				_, err = fmt.Fprintf(f, "  %s: string  `db:\"%s\"`\n", strcase.ToCamel(col), col)
				check(err)
			}
			_, err = fmt.Fprintln(f, "}")
			check(err)
			break
		}
	}
}
