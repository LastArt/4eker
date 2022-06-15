package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type employee struct {
	card uint
	fio  string
	spec string
	sal  string
}

func main() {

	connStr := "user=postgres password=angiolog dbname=scudpg sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	empl := []employee{}

	for rows.Next() {
		p := employee{}
		err := rows.Scan(&p.card)
		if err != nil {
			fmt.Println(err)
			continue
		}
		empl = append(empl, p)
	}
	for _, p := range empl {
		fmt.Println(p)
	}
}
