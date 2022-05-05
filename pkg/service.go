package pkg

import (
	"bufio"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"kontroller/set"
	"log"
	"os"
	"os/exec"
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

func CheckAdminUser(pass string) bool {
	var pwd string
	var res bool = false
	db, err := sql.Open("sqlite3", "./scud.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM superuser WHERE Password = ?", pass)
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
		pwd = p.Pass

	}
	if pwd == "" {
		fmt.Println("Неверный код доступа!")
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

func TmFormat() (string, string) {
	timeFormat := time.Date(2022, time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), 0, 0, time.Local)
	//func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	dm := timeFormat.Format("02.01.2006")
	tm := timeFormat.Format("15:04")
	return dm, tm
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

func NewExcelExport(data string) {
	f := excelize.NewFile()

	f.SetCellValue("Sheet1", "A1", "НОМЕР КАРТЫ")
	f.SetCellValue("Sheet1", "B1", "ФИО")
	f.SetCellValue("Sheet1", "C1", "СПЕЦИАЛЬНОСТЬ")
	f.SetCellValue("Sheet1", "D1", "ДАТА ОТМЕТКИ")
	f.SetCellValue("Sheet1", "E1", "ВРЕМЯ ОТМЕТКИ")

	if err := f.SaveAs(set.ExcelFile); err != nil {
		log.Fatal(err)
	}

}

// Организовать зачистку файла после выдачи пользователю!
func OldExcelDel() {
	f := excelize.NewFile()

	f.SetCellValue("Sheet1", "A1", "")
	f.SetCellValue("Sheet1", "B1", "")
	f.SetCellValue("Sheet1", "C1", "")
	f.SetCellValue("Sheet1", "D1", "")
	f.SetCellValue("Sheet1", "E1", "")

	if err := f.SaveAs(set.ExcelFile); err != nil {
		log.Fatal(err)
	}
}
