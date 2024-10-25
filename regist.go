package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Regist(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/regist.html"))
	name := r.FormValue("name")
	password := r.FormValue("password")

	db, err := sql.Open("mysql", "root:1337@/calculator")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	insert, err := db.Query(fmt.Sprintf("INSERT INTO calculator.users (username, password) VALUES ('%s', '%s')", name, password))
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	tmpl.Execute(w, nil)
}
