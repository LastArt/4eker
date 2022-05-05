package main

import (
	"kontroller/pkg"
	"kontroller/set"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	ExeclWrite()
}

type Journal struct {
	Date string
	Time string
	User pkg.User
}

func ExeclWrite() {
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
