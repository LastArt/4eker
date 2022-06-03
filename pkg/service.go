package pkg

import (
	"4eker/set"
	"bufio"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

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

type Admin struct {
	Id    string
	Login string
	Pass  string
	Email string
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
		fmt.Printf("Логин - %s", lg)
		fmt.Printf("\nПароль - %s", pwd)
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

	if do == "В" {
		GetStart()
	}
	GetSettings()
}

func ClearConsole() {
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("clear")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	out.String()
}

func Check(fio, cardnum, spec, dt, tm string) {
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
}

func TmFormat() time.Time {
	timeFormat := time.Date(2022, time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), 0, 0, time.Local)
	//func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	// dm := timeFormat.Format("02.01.2006")
	// tm := timeFormat.Format("15:04")
	return timeFormat
}

//
type Config struct {
	TelegramBotToken string
}

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

// Журнал посещений за период
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

// Журнал посещений по сотруднику за период
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
