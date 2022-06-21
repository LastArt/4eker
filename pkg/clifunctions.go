package pkg

import (
	"4eker/set"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
)

//==========================
// CLI -- АДМИНИСТРАТОРЫ
//==========================

// Метод добавляющий администраторa через консоль (Меню 6)
func (su SuperUser) Add() {
	db, err := sql.Open("mysql", "u0813820_artur:Zmkstaltex2019@tcp(31.31.198.44:3306)/u0813820_urv")
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
	fmt.Printf(set.Yellow+"Введенная Вами информация корректна?\n"+set.ResetColor+"\nЛогин: %v \nПароль: %v \nПочта: %v \n"+set.Yellow+"\nВведите букву Д-[Да], если согласны или Н-[Нет], если не согласны"+set.ResetColor+"\n", supUser.Login, supUser.Pass, supUser.Email)
	fmt.Scan(&yesNo)
	if yesNo == "Д" || yesNo == "д" || yesNo == "Да" || yesNo == "да" {
		_, err = db.Exec("insert into superuser (Login, Password, Email) values (?, ?, ?)", supUser.Login, supUser.Pass, supUser.Email)
		fmt.Println(set.Green, "Запись в БД успешно выполнена", set.ResetColor)
	} else if yesNo == "Н" || yesNo == "н" || yesNo == "Нет" || yesNo == "нет" {
		fmt.Println(set.Red, "Запись в БД - ПРЕРВАНА", set.ResetColor)
	} else {
		fmt.Println(set.Red, "Введена некорректная информация, попробуйте еще раз ", set.ResetColor)
	}
}

//Метод редактирующий администраторов  из консоли (Меню 7)
func (su SuperUser) Edit(id string) {
	db, err := sql.Open("mysql", "u0813820_artur:Zmkstaltex2019@tcp(31.31.198.44:3306)/u0813820_urv")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println(set.Yellow, "Измените логин суперпользователя:", set.ResetColor)
	su.Login = NewScan()
	fmt.Println(set.Yellow, "Измените пароль суперпользователя:", set.ResetColor)
	su.Pass = NewScan()
	fmt.Println(set.Yellow, "Измените почту суперпользователя:", set.ResetColor)
	su.Email = NewScan()

	var yesNo string
	fmt.Printf(set.Yellow+"Введенная Вами информация корректна?\n"+set.ResetColor+"\nId: %v \nЛогин: %v \nПароль: %v \nПочтовый ящик: %v \n"+set.Yellow+"\nВведите букву Д-[Да], если согласны или Н-[Нет], если не согласны"+set.ResetColor+"\n", id, su.Login, su.Pass, su.Email)
	fmt.Scan(&yesNo)
	if yesNo == "Д" || yesNo == "д" {
		_, err = db.Exec("UPDATE superuser SET login = ? WHERE id = ?", su.Login, id)
		_, err = db.Exec("UPDATE superuser SET password = ? WHERE id = ?", su.Pass, id)
		_, err = db.Exec("UPDATE superuser SET email = ? WHERE id = ?", su.Email, id)
		if err != nil {
			fmt.Println(set.Red, "ОШИБКА:", set.ResetColor, err)
			panic(err)
		}
		fmt.Println(set.Green, "Данные успешно обновлены", set.ResetColor)
		su.ShowAll()
		// Добавить метод отправляющий логин и пароль на почту
	} else if yesNo == "Н" || yesNo == "н" {
		fmt.Println(set.Red, "Процедура прервана", set.ResetColor)
	} else {
		fmt.Println(set.Red, "\nВведена некорректная информация, попробуйте еще раз ", set.ResetColor)
	}
}

//Метод удаляющий  администраторов из консоли (Меню 8)
func (su SuperUser) DeleteRow(id string) {

	db, err := sql.Open("mysql", "u0813820_artur:Zmkstaltex2019@tcp(31.31.198.44:3306)/u0813820_urv")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM superuser WHERE id = ?", id)
	if err != nil {
		fmt.Println(set.Red, "ОШИБКА:", set.ResetColor, err)
		panic(err)
	}
	fmt.Println("Запись по администратору - ", id, "удалена")
}

//Метод выводящий весь список администраторов в консоль (Меню 9)
func (su SuperUser) ShowAll() {
	db, err := sql.Open("mysql", "u0813820_artur:Zmkstaltex2019@tcp(31.31.198.44:3306)/u0813820_urv")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM superuser")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Логин", "Пароль", "Почта"})
	for rows.Next() {
		//var ssId, ssLog, ssPas, ssMail string

		err = rows.Scan(&su.id, &su.Login, &su.Pass, &su.Email)
		if err != nil {
			log.Fatal(err)
		}
		data := [][]string{
			[]string{su.id, su.Login, su.Pass, su.Email},
		}

		for _, v := range data {
			table.Append(v)
		}

	}

	table.Render()
}

//===================
//CLI -- СОТРУДНИКИ
//===================

// Метод добавляющий сотрудника через консоль!
func (u User) Add() {
	db, err := sql.Open("mysql", "u0813820_artur:Zmkstaltex2019@tcp(31.31.198.44:3306)/u0813820_urv")
	if err != nil {
		panic(err)
	}
	var usr = new(User)
	defer db.Close()
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
	if yesNo == "Д" || yesNo == "д" {
		_, err = db.Exec("INSERT INTO user (cardid, fio, speciality, sallary) VALUES (?, ?, ?, ?)", usr.Cardid, usr.Fio, usr.Speciality, usr.Salary)
		if err != nil {
			fmt.Println(set.Red, "ОШИБКА:", set.ResetColor, err)
			panic(err)
		}
		fmt.Println(set.Green, "Запись в БД успешно выполнена", set.ResetColor)

	} else if yesNo == "Н" || yesNo == "н" {
		fmt.Println(set.Red, "Запись в БД - ПРЕРВАНА", set.ResetColor)
	} else {

		fmt.Println(set.Red, "\nYВведена некорректная информация, попробуйте еще раз ", set.ResetColor)
	}
}

// Метод редактирующий сотрудников через консоль
func (u User) Edit(cardnum string) {
	db, err := sql.Open("mysql", "u0813820_artur:Zmkstaltex2019@tcp(31.31.198.44:3306)/u0813820_urv")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println(set.Yellow, "Измените ФИО:", set.ResetColor)
	u.Fio = NewScan()
	fmt.Println(set.Yellow, "Измените специальность сотрудника:", set.ResetColor)
	u.Speciality = NewScan()
	fmt.Println(set.Yellow, "Измените зарплату сотрудника:", set.ResetColor)
	u.Salary = NewScan()

	var yesNo string
	fmt.Printf(set.Yellow+"Введенная Вами информация корректна?\n"+set.ResetColor+"\n№ Карты: %v \nФио сотрудника: %v \nСпециальность: %v \nЗарплата сотрудника: %v \n"+set.Yellow+"\nВведите букву Д-[Да], если согласны или Н-[Нет], если не согласны"+set.ResetColor+"\n", cardnum, u.Fio, u.Speciality, u.Salary)
	fmt.Scan(&yesNo)
	if yesNo == "Д" || yesNo == "д" {
		_, err = db.Exec("UPDATE user SET fio = ? WHERE cardid = ?", u.Fio, cardnum)
		_, err = db.Exec("UPDATE user SET speciality = ? WHERE cardid = ?", u.Speciality, cardnum)
		_, err = db.Exec("UPDATE user SET salary = ? WHERE cardid = ?", u.Salary, cardnum)
		if err != nil {
			fmt.Println(set.Red, "ОШИБКА:", set.ResetColor, err)
			panic(err)
		}
		fmt.Println(set.Green, "Данные успешно обновлены", set.ResetColor)
		u.ShowAll()
	} else if yesNo == "Н" || yesNo == "н" {
		fmt.Println(set.Red, "Процедура прервана", set.ResetColor)
	} else {

		fmt.Println(set.Red, "\nYВведена некорректная информация, попробуйте еще раз ", set.ResetColor)
	}

}

// Метод удаляющий сотрудника в консоли!
func (u User) DeleteRow(cardnum string) {
	db, err := sql.Open("mysql", "u0813820_artur:Zmkstaltex2019@tcp(31.31.198.44:3306)/u0813820_urv")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM user WHERE cardid = ?", cardnum)

	if err != nil {
		fmt.Println(set.Red, "ОШИБКА:", set.ResetColor, err)
		panic(err)
	}
	fmt.Println(set.Green, "Запись по № карты - ", cardnum, "удалена", set.ResetColor)
}

// Метод выводящий весь список сотрудников в консоль!
func (u User) ShowAll() {
	db, err := sql.Open("mysql", "u0813820_artur:Zmkstaltex2019@tcp(31.31.198.44:3306)/u0813820_urv")
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

// Метод выводящий в консоль присутсвующих на текущую дату сотрудников.
func (j Journal) WhoInPlace() {
	datetime := TmFormat()
	dt := datetime.Format("02.01.2006")
	tm := datetime.Format("15:04")
	var id string
	fmt.Println("Сегодня: ", dt)
	db, err := sql.Open("mysql", "u0813820_artur:Zmkstaltex2019@tcp(31.31.198.44:3306)/u0813820_urv")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM journal WHERE date = ? AND time < ?", dt, tm) // Требуется доработка SQL запроса Вывести список сотрудников За сегодня и За время с 8 утра до текущего часа
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
		//jDt := j.Date.Format("02.01.2006") //!  Внес поправку для тестирования нового формата
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

//Проверка валидности карты пропуска
func (u User) CheckInTimeValidation(cardnum string) []string {
	var crd string
	var str = []string{"n/a", "n/a", "n/a", "n/a"}

	db, err := sql.Open("mysql", "u0813820_artur:Zmkstaltex2019@tcp(31.31.198.44:3306)/u0813820_urv")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM user WHERE cardid = ?", cardnum)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	users := []User{}

	for rows.Next() {
		p := User{}
		err := rows.Scan(&u.Cardid, &u.Fio, &u.Speciality, &u.Salary)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p)
		crd = u.Cardid
		str[0] = u.Cardid
		str[1] = u.Fio
		str[2] = u.Speciality
		str[3] = u.Salary

	}
	if crd == "" {
		fmt.Println(set.Red, set.AccesDenied, set.ResetColor)
	} else {
		fmt.Println(set.Cyan, set.Acces, set.ResetColor)

	}
	return str
}
