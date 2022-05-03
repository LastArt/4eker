package pkg

import (
	"database/sql"
	"fmt"
	"kontroller/set"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/olekukonko/tablewriter"
)

type User struct {
	Cardid     string
	Fio        string
	Speciality string
	Salary     string
}

type Journal struct {
	Date string
	Fio  string
	Time string
}

type SuperUser struct {
	Login string
	Pass  string
	Email string
}

type EmplChanger interface {
	Add()
	Edit()
	AddInBot()
	DeleteRow()
	ClearDB()
	ShowAll()
	DeleteRowInBot()
	ShowAllInBot()
	EditFromBot()
}

//! ============= МЕТОДЫ ПО АДМИНИСТРАТОРАМ ====================
//Метод добавляющий администраторa через консоль
func (su SuperUser) Add() {
	db, err := sql.Open("sqlite3", "./scud.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var supUser = new(SuperUser)

	fmt.Println(string(set.Yellow), "Введите логин", string(set.ResetColor))
	supUser.Login = NewScan()

	fmt.Println(string(set.Yellow), "Введите пароль", string(set.ResetColor))
	supUser.Pass = NewScan()

	fmt.Println(string(set.Yellow), "Введите почту", string(set.ResetColor))
	supUser.Email = NewScan()

	if err != nil {
		fmt.Println(set.Red, "ОШИБКА:", set.ResetColor, err)
		panic(err)
	}
	var yesNo string
	fmt.Printf(set.Yellow+"Введенная Вами информация корректна?\n"+set.ResetColor+"\nЛогин: %v \nПароль: %v \nПочта: %v \n"+set.Yellow+"\nВведите букву Y-[Yes], если согласны или N-[N], если не согласны"+set.ResetColor+"\n", supUser.Login, supUser.Pass, supUser.Email)
	fmt.Scan(&yesNo)
	if yesNo == "Y" {
		_, err = db.Exec("insert into superuser (Login, Password, Email) values (?, ?, ?)", supUser.Login, supUser.Pass, supUser.Email)
		fmt.Println(set.Green, "Запись в БД успешно выполнена", set.ResetColor)
	} else if yesNo == "N" {
		fmt.Println(set.Red, "Запись в БД - ПРЕРВАНА", set.ResetColor)
	} else {
		fmt.Println(set.Red, "Введена некорректная информация, попробуйте еще раз ", set.ResetColor)
	}
}

//Метод удаляющий  администраторов из консоли
func (su SuperUser) DeleteRow(login string) {

	db, err := sql.Open("sqlite3", "./scud.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// удаляем строку с id=1
	_, err = db.Exec("DELETE FROM superuser WHERE Login = ?", login)
	if err != nil {
		fmt.Println(set.Red, "ОШИБКА:", set.ResetColor, err)
		panic(err)
	}
	fmt.Println("Запись по администратору - ", login, "удалена") // количество удаленных строк
}

//Метод удаляющий  администраторов из чата бота
func (su SuperUser) DeleteRowInBot(log string) string {
	var res string
	db, err := sql.Open("sqlite3", "./scud.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// удаляем строку с Login=log
	_, err = db.Exec("delete from superuser where Login = ?", log)
	if err != nil {
		res = "⛔️Ошибка выполнения удаления из БД\nПопробуйте еще раз!"
	} else {
		res = "✅Удаление прошло успешно!\nДля выхода из режима нажмите '🔙 Вернуться' или продолжайте удаление пользователей из БД!"
	}
	return res
}

//Метод выводящий весь список администраторов в консоль
func (su SuperUser) ShowAll() {
	db, err := sql.Open("sqlite3", "./scud.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from superuser")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Логин", "Пароль", "Почта"})
	for rows.Next() {
		var ssId, ssLog, ssPas, ssMail string

		err = rows.Scan(&ssId, &ssLog, &ssPas, &ssMail)
		if err != nil {
			log.Fatal(err)
		}
		data := [][]string{
			[]string{ssId, ssLog, ssPas, ssMail},
		}

		for _, v := range data {
			table.Append(v)
		}

	}

	table.Render()
}

//TODO Метод редактирующий администраторов  из консоли
func (su SuperUser) Edit() {
}

//TODO Метод редактирующий администраторов  из чат бота
func (su SuperUser) EditFromBot() {

}

//Метод выводящий весь список администраторов в чат бота
func (su SuperUser) ShowAllInBot() string {
	var id string
	db, err := sql.Open("sqlite3", "./scud.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from superuser")
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

// Метод добавляющий администратора через чат бота!
func (su SuperUser) AddInBot(log, pass, mail string) string {
	var res string
	db, err := sql.Open("sqlite3", "./scud.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("insert into superuser (Login, Password, Email) values (?, ?, ?)", log, pass, mail)
	if err != nil {
		res = "⛔️Ошибка выполнения записи в БД\nПопробуйте еще раз!"
	} else {
		res = "✅Запись прошла успешно!\nДля выхода из режима нажмите '🔙 Вернуться' или продолжайте заполнение базы данных!"
	}

	return res
}

//! ============= МЕТОДЫ ПО СОТРУДНИКАМ ====================
// Метод добавляющий сотрудника через консоль!
func (u User) Add() {
	db, err := sql.Open("sqlite3", "./scud.db")
	if err != nil {
		panic(err)
	}
	var usr = new(User)

	fmt.Println(set.Yellow, "Введите номер карты-пропуска:", set.ResetColor)
	usr.Cardid = NewScan()
	fmt.Println(set.Yellow, "Введите ФИО сотрудника:", set.ResetColor)
	usr.Fio = NewScan()
	fmt.Println(set.Yellow, "Введите специальность сотрудника:", set.ResetColor)
	usr.Speciality = NewScan()
	fmt.Println(set.Yellow, "Введите зарплату сотрудника:", set.ResetColor)
	usr.Salary = NewScan()

	var yesNo string
	fmt.Printf(set.Yellow+"Введенная Вами информация корректна?\n"+set.ResetColor+"\n№ Карты: %v \nФио сотрудника: %v \nСпециальность: %v \nЗарплата сотрудника: %v \n"+set.Yellow+"\nВведите букву Д-[Да], если согласны или Н-[Нет], если не согласны"+set.ResetColor+"\n", usr.Cardid, usr.Fio, usr.Speciality, usr.Salary)
	fmt.Scan(&yesNo)
	if yesNo == "Д" {
		_, err = db.Exec("insert into user (CardId, Fio, Speciality, Salary) values (?, ?, ?, ?)", usr.Cardid, usr.Fio, usr.Speciality, usr.Salary)
		if err != nil {
			fmt.Println(set.Red, "ОШИБКА:", set.ResetColor, err)
			panic(err)
		}
		fmt.Println(set.Green, "Запись в БД успешно выполнена", set.ResetColor)

	} else if yesNo == "Д" {
		fmt.Println(set.Red, "Запись в БД - ПРЕРВАНА", set.ResetColor)
	} else {

		fmt.Println(set.Red, "\nYВведена некорректная информация, попробуйте еще раз ", set.ResetColor)
	}
	db.Close()
}

// Метод добавляющий сотрудника через чат бота!
func (u User) AddInBot(card, fio, spec, sal string) string {
	var res string
	db, err := sql.Open("sqlite3", "./scud.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("insert into user (CardId, Fio, Speciality, Salary) values (?, ?, ?, ?)", card, fio, spec, sal)
	if err != nil {
		res = "⛔️Ошибка выполнения записи в БД\nПопробуйте еще раз!"
	} else {
		res = "✅Запись прошла успешно!\nДля выхода из режима нажмите '🔙 Вернуться' или продолжайте заполнение базы данных!"
	}

	return res
}

// Метод удаляющий сотрудника в консоли!
func (u User) DeleteRow(cardnum string) {
	db, err := sql.Open("sqlite3", "./scud.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("DELETE from USER where CardId = ?", cardnum)

	if err != nil {
		fmt.Println(set.Red, "ОШИБКА:", set.ResetColor, err)
		panic(err)
	}
	fmt.Println(set.Green, "Запись по № карты - ", cardnum, "удалена", set.ResetColor)
}

// Метод удаляющий сотрудника через чат бота!
func (u User) DeleteRowInBot(fio string) string {
	var res string
	db, err := sql.Open("sqlite3", "./scud.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("delete from user where Fio = ?", fio)

	if err != nil {
		res = "⛔️Ошибка выполнения удаления из БД\nПопробуйте еще раз!"
	} else {
		res = "✅Удаление прошло успешно!\nДля выхода из режима нажмите '🔙 Вернуться' или продолжайте удаление пользователей из БД!"
	}
	return res
}

// Метод выводящий весь список сотрудников в консоль!
func (u User) ShowAll() {
	db, err := sql.Open("sqlite3", "./scud.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Номер карты", "ФИО Сотрудника", "Должность", "Зарплата"})
	for rows.Next() {
		var sCardNum, sFio, sSpec, sSalary string

		err = rows.Scan(&sCardNum, &sFio, &sSpec, &sSalary)
		if err != nil {
			log.Fatal(err)
		}
		data := [][]string{
			[]string{sCardNum, sFio, sSpec, sSalary},
		}

		for _, v := range data {
			table.Append(v)
		}

	}

	table.Render()
}

// Метод выводящий весь список сотрудников в чат бота!
func (u User) ShowAllInBot() string {

	db, err := sql.Open("sqlite3", "./scud.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from user")
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

//TODO Метод редактирующий сотрудников через консоль
func (u User) Edit() {

}

//TODO Метод редактирующий сотрудников через чат бота
func (u User) EditFromBot() {

}

//! ============= МЕТОДЫ ПО РАБОТЕ С ТРЕКИНГОМ ====================

func (u User) CheckInTimeValidation(cardnum string) []string { // Определяем валидность номера карты
	var crd string
	var str = []string{"0", "0", "0"}

	db, err := sql.Open("sqlite3", "./scud.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM user WHERE CardId = ?", cardnum)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	users := []User{}

	for rows.Next() {
		p := User{}
		err := rows.Scan(&p.Cardid, &p.Fio, &p.Speciality, &p.Salary)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p)
		crd = p.Cardid
		str[0] = p.Cardid
		str[1] = p.Fio
		str[2] = p.Speciality
		//str = append(str, p.Cardid, p.Fio, p.Speciality)

	}
	if crd == "" {
		fmt.Println(set.Red, set.AccesDenied, set.ResetColor)
	} else {
		fmt.Println(set.Cyan, set.Acces, set.ResetColor)

	}
	return str
}

func (u User) AdCheckinToDb(fio, dt, tm string) { // Записываем трек в бд
	db, err := sql.Open("sqlite3", "./scud.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO journal (FioVisiter, Date, Time) values (?, ?, ?)", fio, dt, tm)
	if err != nil {
		fmt.Println(set.Red, "ОШИБКА:", set.ResetColor, err)
	}
}

//! ============= МЕТОДЫ ПО РАБОТЕ С ОТЧЕТАМИ =====================

//! ============= МЕТОДЫ ПО РАБОТЕ С СИСТЕМОЙ КОНТРОЛЯ КТО В ЦЕХУ =====================
func (j Journal) WhoInPlace() {
	today := "Sun May  1 03:21:54 2022"
	var id string
	fmt.Println(today)
	db, err := sql.Open("sqlite3", "./scud.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM journal WHERE Date = ?", today)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	journals := []Journal{}
	for rows.Next() {
		p := Journal{}
		err := rows.Scan(&id, &p.Date, &p.Fio, &p.Time)
		if err != nil {
			fmt.Println(err)
			continue
		}
		journals = append(journals, p)
	}
	fmt.Println(journals)
}
