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
	Time string
	User User
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

//*============= МЕТОДЫ ПО АДМИНИСТРАТОРАМ ====================
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

// Метод добавляющий администратора через чат бота!
func (su SuperUser) AddInBot(log, pass, mail string) string {
	var res string
	db, err := sql.Open("sqlite3", "./scud.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO superuser (Login, Password, Email) VALUES (?, ?, ?)", log, pass, mail)
	if err != nil {
		res = "⛔️Ошибка выполнения записи в БД\nПопробуйте еще раз!"
	} else {
		res = "✅Запись прошла успешно!\nДля выхода из режима нажмите '🔙 Вернуться' или продолжайте заполнение базы данных!"
	}

	return res
}

//============= МЕТОДЫ ПО СОТРУДНИКАМ ====================
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
		_, err = db.Exec("INSERT INTO user (CardId, Fio, Speciality, Salary) VALUES (?, ?, ?, ?)", usr.Cardid, usr.Fio, usr.Speciality, usr.Salary)
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
	_, err = db.Exec("INSERT INTO user (CardId, Fio, Speciality, Salary) VALUES (?, ?, ?, ?)", card, fio, spec, sal)
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
	_, err = db.Exec("DELETE FROM user WHERE Fio = ?", fio)

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
		//var sCardNum, sFio, sSpec, sSalary string
		err = rows.Scan(&u.Cardid, &u.Fio, &u.Speciality, &u.Salary)
		if err != nil {
			log.Fatal(err)
		}
		data := [][]string{
			[]string{u.Cardid, u.Fio, u.Speciality, u.Salary},
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

//TODO Метод редактирующий сотрудников через консоль
func (u User) Edit() {

}

//TODO Метод редактирующий сотрудников через чат бота
func (u User) EditFromBot() {

}

//Проверка валидности карты пропуска
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
		fmt.Println(set.ERROR_INSERT_TODB, err)
	}
}

//! ============= МЕТОДЫ ПО РАБОТЕ С ОТЧЕТАМИ =====================

// Метод выводящий в консоль всех присутсвующих на текущую дату.
func (j Journal) WhoInPlace() {
	today, totime := TmFormat()
	var id string
	fmt.Println("Сегодня: ", today)
	db, err := sql.Open("sqlite3", "./scud.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM journal WHERE Date = ? AND Time < ?", today, totime) // Требуется доработка SQL запроса Вывести список сотрудников За сегодня и За время с 8 утра до текущего часа
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	journals := []Journal{}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID записи", "ФИО", "Дата отметки", "Время отметки"})
	for rows.Next() {
		p := Journal{}
		err := rows.Scan(&id, &j.User.Fio, &j.Date, &j.Time)
		j.Date, _ = TmFormat()
		if err != nil {
			fmt.Println(err)
			continue
		}
		journals = append(journals, p)
		data := [][]string{
			[]string{id, j.User.Fio, j.Date, j.Time},
		}
		for _, v := range data {
			table.Append(v)
		}
	}
	table.Render()

}

func (j Journal) WhoInPlaceForBot() string {

	today, totime := TmFormat()
	var id string
	fmt.Println("Сегодня: ", today)
	db, err := sql.Open("sqlite3", "./scud.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM journal WHERE Date = ? AND Time < ?", today, totime)
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
		//j.Date, _ = TmFormat()
		str += "\nID: " + id + "\nФИО: " + j.User.Fio + "\nДата отметки: " + j.Date + "\nВремя отметки: " + j.Time + "\n=====================\n"
	}
	return str
}

// TODO Работа с файлом excel
func (j Journal) ExportToExcel() {
	db, err := sql.Open("sqlite3", "./scud.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM journal WHERE Data = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	//var str string
	// for rows.Next() {
	// 	err := rows.Scan(&u.Cardid, &u.Fio, &u.Speciality, &u.Salary)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	}
	// 	str += "\n Номер карты: " + u.Cardid + "\n ФИО: " + u.Fio + "\n Должность: " + u.Speciality + "\n Зарплата: " + u.Salary + "\n=====================\n"
	// }
}

//TODO  Модуль добавления TOKENA в меню программы
//TODO  Распределить все настройки по группам
