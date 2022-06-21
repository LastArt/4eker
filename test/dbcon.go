package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Admin struct {
	id       int
	Login    string
	Password string
	Email    string
}

func main() {
	db, err := sql.Open("mysql", "u0813820_artur:Zmkstaltex2019@tcp(31.31.198.44:3306)/u0813820_urv")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM u0813820_urv.superuser")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	admins := []Admin{}

	for rows.Next() {
		p := Admin{}
		err := rows.Scan(&p.id, &p.Login, &p.Password, &p.Email)
		if err != nil {
			fmt.Println(err)
			continue
		}
		admins = append(admins, p)
	}
	for _, p := range admins {
		fmt.Println(p.id, p.Login, p.Password, p.Email)
	}
}

/*
db, err := sql.Open("mysql", "u0813820_artur:Zmkstaltex2019@tcp(31.31.198.44:3306)/u0813820_urv")
if err != nil {
	panic(err)
}
defer db.Close()
result, err := db.Exec("INSERT INTO u0813820_urv.superuser (login, password, email) VALUES (?, ?, ?)", "ar", "qw123", "qwe2344@mail.ru")
if err != nil {
	panic(err)
}
log.Println(result.LastInsertId())
*/
