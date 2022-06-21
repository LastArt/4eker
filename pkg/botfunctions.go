package pkg

import (
	"4eker/set"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//Метод добавляющий администратора через чат бота!
func (su SuperUser) AddInBot(log, pass, mail string) string {
	var res string
	db, err := sql.Open("mysql", ConnString())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO superuser (login, password, email) VALUES (?, ?, ?)", log, pass, mail)
	if err != nil {
		res = set.BOT_ERROR_ADDTODB
	} else {
		res = set.BOT_WARNING_ADD_FINE
	}

	return res
}

//Метод удаляющий  администраторов из чата бота
func (su SuperUser) DeleteRowInBot(id string) string {
	var res string
	db, err := sql.Open("mysql", ConnString())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM superuser WHERE id = ?", id)
	if err != nil {
		res = set.BOT_ERROR_DELETE_USER
	} else {
		res = set.BOT_WARNING_DELETE_FINE
	}
	return res
}

// Метод редактирующий администраторов  из чат бота
func (su SuperUser) EditFromBot(id, newLogin, newPass, newMail string) (string, string) {
	var res, show string
	db, err := sql.Open("mysql", ConnString())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE superuser SET login = ? WHERE id = ?", newLogin, id)
	_, err = db.Exec("UPDATE superuser SET password = ? WHERE id = ?", newPass, id)
	_, err = db.Exec("UPDATE superuser SET email = ? WHERE id = ?", newMail, id)
	if err != nil {
		res = set.BOT_ERROR_EDITUSER
	} else {
		res = set.BOT_WARNING_EDITUSER_FINE
		show = su.ShowAllInBot()
	}

	return res, show
}

// 3 Метод выводящий весь список администраторов в чат бота
func (su SuperUser) ShowAllInBot() string {
	var id string
	db, err := sql.Open("mysql", ConnString())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM superuser")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var str string

	for rows.Next() {
		err := rows.Scan(&id, &su.Login, &su.Pass, &su.Email)
		if err != nil {
			fmt.Println(err)
			continue
		}

		str += "\n ID: " + id + "\n Логин: " + su.Login + "\n Пароль: " + su.Pass + "\n Почта: " + su.Email + "\n=====================\n"
	}
	return str
}

// Метод выводящий в бот всех присутсвующих на текущую дату.
func (j Journal) WhoInPlaceForBot() string {

	datetime := TmFormat()
	dt := datetime.Format("02.01.2006")
	tm := datetime.Format("15:04")
	var id string
	fmt.Println("Сегодня: ", dt)
	db, err := sql.Open("mysql", ConnString())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM journal WHERE date = ? AND time < ?", dt, tm)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var str string

	for rows.Next() {
		err := rows.Scan(&id, &j.User.Fio, &j.Date, &j.Time)
		if err != nil {
			fmt.Println(err)
			continue
		}
		str += "\nID: " + id + "\nФИО: " + j.User.Fio + "\nДата отметки: " + j.Date + "\nВремя отметки: " + j.Time + "\n=====================\n"
	}
	return str
}

// Метод выводящий весь список сотрудников в бота!
func (u User) ShowAllInBot() string {
	db, err := sql.Open("mysql", ConnString())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var str string
	for rows.Next() {
		err := rows.Scan(&u.Cardid, &u.Fio, &u.Speciality, &u.Salary)
		if err != nil {
			fmt.Println(err)
			continue
		}
		str += "\n Номер карты: " + u.Cardid + "\n ФИО: " + u.Fio + "\n Должность: " + u.Speciality + "\n Зарплата: " + u.Salary + "\n=====================\n"
	}
	return str
}

// Метод удаляющий сотрудника через чат бота!
func (u User) DeleteRowInBot(card string) (string, string) {
	var res, show string
	db, err := sql.Open("mysql", ConnString())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM user WHERE cardid = ?", card)
	if err != nil {
		res = set.BOT_ERROR_DELETE_USER
	} else {
		res = set.BOT_WARNING_DELETE_FINE
		show = u.ShowAllInBot()
	}
	return res, show
}

// Метод добавляющий сотрудника через чат бота!
func (u User) AddInBot(card, fio, spec, sal string) string {
	var res string
	db, err := sql.Open("mysql", ConnString())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO user (cardid, fio, speciality, sallary) VALUES (?, ?, ?, ?)", card, fio, spec, sal)
	if err != nil {
		res = set.BOT_ERROR_ADDTODB

	} else {
		res = set.BOT_WARNING_ADD_FINE
	}

	return res
}

//Метод редактирующий сотрудников через чат бота
func (u User) EditFromBot(cardnum, newFio, newSpec, newSal string) (string, string) {
	var res, show string
	db, err := sql.Open("mysql", ConnString())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE user SET fio = ? WHERE cardid = ?", newFio, cardnum)
	_, err = db.Exec("UPDATE user SET speciality = ? WHERE cardid = ?", newSpec, cardnum)
	_, err = db.Exec("UPDATE user SET sallary = ? WHERE cardid = ?", newSal, cardnum)

	if err != nil {
		res = set.BOT_ERROR_EDITUSER
	} else {

		res = set.BOT_WARNING_EDITUSER_FINE
		show = u.ShowAllInBot()
	}

	return res, show
}

// Оповещение об отметке
/*func Check(fio, cardnum, spec, dt, tm string) {
	tg_key := GetKey(set.TokenFile)
	c := New(tg_key)
	chatId := -644032460
	if fio != "0" {
		cText := ("🟢Отметился:  \n" + fio + "\nНомер: " + cardnum + "\nДолжность: " + spec + "\nДата: " + dt + "\nВремя: " + tm)
		err := c.SendMessage(cText, int64(chatId))
		if err != nil {
			fmt.Println(set.ERROR_SEND_BOT, err)
		}
	}
}*/
