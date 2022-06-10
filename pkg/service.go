package pkg

import (
	"4eker/set"
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/xuri/excelize/v2"
)

type Admin struct {
	Id    string
	Login string
	Pass  string
	Email string
}
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
	id    string
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

func NewScan() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, set.ERROR_INPUT, err)
	}

	return in.Text()
}

func StringScan() string {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(set.ERROR_INPUT, err)
	}
	return str
}

func CheckAdminUser(log, pass string) bool {
	var res bool = false
	var pwd, lg string
	db, err := sql.Open("sqlite3", "./scud.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM superuser WHERE Login=? AND Password =?", log, pass)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		p := new(Admin)
		err := rows.Scan(&p.Id, &p.Login, &p.Pass, &p.Email)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// admins = append(admins, p)
		lg = p.Login
		pwd = p.Pass
	}
	if lg == "" && pwd == "" {
		res = false
	} else {
		res = true
	}
	return res
}

func NumberValuator(msgIn string) []string {
	var res = []string{}
	words := strings.Split(msgIn, "/")
	for _, i := range words {
		res = append(res, i)
	}
	return res
}

func Question() {
	var do string
	fmt.Println("Для выхода нажмите В-[Выход] для продолжения работы нажмите любую кнопку и Enter")
	fmt.Scan(&do)
	if do == "В" || do == "в" {
		GetStart()
	}
	GetCliMenu()
}

func TmFormat() time.Time {
	timeFormat := time.Date(2022, time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), 0, 0, time.Local)
	return timeFormat
}

// Допфункция для получения ключа из json файла
func GetKey(path string) string {
	var TG_token string
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf(set.ERROR_TOKEN, err)
		panic(err.Error())
	}

	decoder := json.NewDecoder(file)
	conf := Config{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err.Error())
	}

	TG_token = conf.TelegramBotToken

	return TG_token
}

// Формируем журнал посещений за период
func PresentJournalToDay(data1, data2 string) {
	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "ФИО")
	f.SetCellValue("Sheet1", "C1", "ДАТА ОТМЕТКИ")
	f.SetCellValue("Sheet1", "D1", "ВРЕМЯ ОТМЕТКИ")

	jrnl := new(Journal)
	var id string

	db, err := sql.Open("sqlite3", "./scud.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM journal WHERE Date BETWEEN ? AND  ? ", data1, data2)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for i := 2; rows.Next(); i++ {
		err := rows.Scan(&id, &jrnl.User.Fio, &jrnl.Date, &jrnl.Time)
		if err != nil {
			fmt.Println(err)
			continue
		}

		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i), id)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i), jrnl.User.Fio)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i), jrnl.Date)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i), jrnl.Time)

	}

	if err := f.SaveAs(set.ExcelFile); err != nil {
		log.Fatal(err)
	}
}

// Формируем журнал посещений по сотруднику за период
func PresentJournalToEmpl(data1, data2, fio string) {
	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "ФИО")
	f.SetCellValue("Sheet1", "C1", "ДАТА ОТМЕТКИ")
	f.SetCellValue("Sheet1", "D1", "ВРЕМЯ ОТМЕТКИ")

	jrnl := new(Journal)
	var id string

	db, err := sql.Open("sqlite3", "./scud.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM journal WHERE Date BETWEEN ? AND  ? AND FioVisiter=? ", data1, data2, fio)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for i := 2; rows.Next(); i++ {
		err := rows.Scan(&id, &jrnl.User.Fio, &jrnl.Date, &jrnl.Time)
		if err != nil {
			fmt.Println(err)
			continue
		}

		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i), id)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i), jrnl.User.Fio)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i), jrnl.Date)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i), jrnl.Time)

	}

	if err := f.SaveAs(set.ExcelFile); err != nil {
		log.Fatal(err)
	}
}

// Формируем журнал посещений  за период
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
}

//Запись отметки в бд
func (u User) AdCheckinToDb(fio string, datetime time.Time) {
	db, err := sql.Open("sqlite3", "./scud.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	dt := datetime.Format("02.01.2006")
	tm := datetime.Format("15:04")
	_, err = db.Exec("INSERT INTO journal (FioVisiter, Date, Time) values (?, ?, ?)", fio, dt, tm)
	if err != nil {
		fmt.Println(set.ERROR_INSERT_TODB, err)
	}
}
