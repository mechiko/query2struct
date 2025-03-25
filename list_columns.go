package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/jmoiron/sqlx"
)

//go:embed list_column.sql
var listColumnTableSql string

//go:embed tmpl_struct.txt
var tmplStruct string

type tmplData struct {
	Table           string
	TableName       string
	ColumnList      string
	ColumnListPK    string
	VarColumnList   string
	VarColumnListPK string
	BindList        string
	BindListPK      string
	SetList         string
	UpdateList      string
	FieldPK         string
	VarFieldPK      string
	Columns         columnSlice
}

func listColumns(db *sqlx.DB, table string) error {
	var str string
	tableColumns := make(columnSlice, 0)
	err := db.Select(&tableColumns, listColumnTableSql, table)
	check(err)
	listColumnsPK := []string{}
	listColumns := []string{}
	listVarColumnsPK := []string{}
	listVarColumns := []string{}
	bindColumns := []string{}
	bindColumnsPK := []string{}
	updateList := []string{}
	updatePK := ""
	setList := []string{}
	fieldPK := ""
	for i, cl := range tableColumns {
		tableColumns[i].NameCamel = CamelCase(cl.Name)
		typeCol := strings.ToLower(cl.Type)
		if cl.Pk == 1 {
			listColumnsPK = append(listColumnsPK, cl.Name)
			listVarColumnsPK = append(listVarColumnsPK, "r."+cl.NameCamel)
			bindColumnsPK = append(bindColumnsPK, "?")
			updatePK = "r." + cl.NameCamel
			fieldPK = cl.Name
			if strings.HasPrefix(typeCol, "integer") {
				tableColumns[i].Type = "int64"
			} else if strings.HasPrefix(typeCol, "bool") {
				tableColumns[i].Type = "bool"
			} else {
				tableColumns[i].Type = "string"
			}
		} else {
			listColumns = append(listColumns, cl.Name)
			listVarColumns = append(listVarColumns, "r."+cl.NameCamel)
			bindColumns = append(bindColumns, "?")
			setList = append(setList, fmt.Sprintf("%s=?", cl.Name))
			updateList = append(updateList, "r."+cl.NameCamel)
			if strings.HasPrefix(typeCol, "integer") {
				tableColumns[i].Type = "int"
			} else if strings.HasPrefix(typeCol, "bool") {
				tableColumns[i].Type = "bool"
			} else {
				tableColumns[i].Type = "string"
			}
		}
	}
	updateList = append(updateList, updatePK)
	data := tmplData{
		Table:           table,
		TableName:       CamelCase(table),
		Columns:         tableColumns,
		ColumnList:      strings.Join(listColumns, ","),
		ColumnListPK:    strings.Join(listColumnsPK, ","),
		VarColumnList:   strings.Join(listVarColumns, ","),
		VarColumnListPK: strings.Join(listVarColumnsPK, ","),
		BindList:        strings.Join(bindColumns, ","),
		BindListPK:      strings.Join(bindColumnsPK, ","),
		SetList:         strings.Join(setList, ","),
		UpdateList:      strings.Join(updateList, ","),
		FieldPK:         fieldPK,
		VarFieldPK:      updatePK,
	}
	str, err = PrintStruct(data)
	check(err)
	fileName := filepath.Join("output", fmt.Sprintf("%s.go", table))
	err = os.WriteFile(fileName, []byte(str), 0644)
	check(err)
	return nil
}

func PrintStruct(k interface{}) (string, error) {
	var err error
	var buf bytes.Buffer
	var result string = ""

	funcMap := template.FuncMap{
		// The name "inc" is what the function will be called in the template text.
		"inc": func(i int) int {
			return i + 1
		},
	}

	t := template.Must(template.New("struct").Funcs(funcMap).Parse(tmplStruct))
	err = t.ExecuteTemplate(&buf, "struct", k)
	if err != nil {
		return result, err
	}
	result = buf.String()
	return result, err
}

func CamelCase(str string) string {
	words := strings.Split(str, "_")
	key := strings.Title(words[0])
	for _, word := range words[1:] {
		key += strings.Title(word)
	}
	return key
}
