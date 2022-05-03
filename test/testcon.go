package main

import (
	"database/sql"
	"fmt"
	"kontroller/pkg"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Admin struct {
	Id    string
	Login string
	Pass  string
	Email string
}

func main() {
	var lgn string
	var pwd string
	db, err := sql.Open("sqlite3", "./scud.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var log, pass string
	fmt.Println("Введите логин")
	fmt.Scan(&log)
	fmt.Println("Введите пароль")
	fmt.Scan(&pass)
	rows, err := db.Query("SELECT * FROM superuser WHERE Login = ? AND Password = ?", log, pass)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	admins := []Admin{}

	for rows.Next() {
		p := Admin{}
		err := rows.Scan(&p.Id, &p.Login, &p.Pass, &p.Email)
		if err != nil {
			fmt.Println(err)
			continue
		}
		admins = append(admins, p)
		lgn = p.Login
		pwd = p.Pass

	}
	fmt.Println("lgn - ", lgn)
	fmt.Println("pwd - ", pwd)
}

func Some(fio, cardnum, spec string) {
	c := pkg.New("5051681006:AAEU_3nVrrO5HMR8Ri3w4159NcshdclxgTI")
	chatId := -644032460
	cText := ("🟢Отметился:  " + fio + "\nНомер: " + cardnum + "\nДолжность: " + spec + "\nДата: " + time.Now().String())
	err := c.SendMessage(cText, int64(chatId))
	if err != nil {
		fmt.Println("ОШИБКА ОТПРАВКИ СООБЩЕНИЯ БОТУ: ", err)
	}
}
