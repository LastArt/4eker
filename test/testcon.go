package main

import (
	"4eker/pkg"
	"4eker/set"
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	var data1, data2, fio string
	fmt.Println("Введите начало периода")
	data1 = pkg.NewScan()
	fmt.Println("Введите конец периода")
	data2 = pkg.NewScan()
	fmt.Println("Введите ФИО")
	fio = pkg.NewScan()
	ExeclWrite(data1, data2, fio)

}

type Journal struct {
	Date string
	Time string
	User pkg.User
}

func ExeclWrite(data1, data2, fio string) {
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
	rows, err := db.Query("SELECT * FROM journal WHERE Date BETWEEN ? AND  ? AND FioVisiter =?", data1, data2, fio)
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

		fmt.Println(id)
		fmt.Println(jrnl.User.Fio)
		fmt.Println(jrnl.Date)
		fmt.Println(jrnl.Time)

		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i), id)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i), jrnl.User.Fio)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i), jrnl.Date)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i), jrnl.Time)

	}

	if err := f.SaveAs(set.ExcelFile); err != nil {
		log.Fatal(err)
	}
}
