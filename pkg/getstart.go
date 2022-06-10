package pkg

import (
	"4eker/set"
	"fmt"
	"regexp"
	"time"
)

func GetStart() {

	var emplID, adminLog, adminPass string
	var check bool
	pattern := "[[:digit:]]"
	reg := regexp.MustCompile(pattern)

	fmt.Println(set.Descr)

	emplID = NewScan()

	isNumbers := reg.MatchString(emplID)
	if emplID == "--admin" {
		fmt.Println(string(set.Yellow), "Введите логин администратора", string(set.ResetColor))
		adminLog = NewScan()
		fmt.Println(string(set.Yellow), "Введите пароль администратора", string(set.ResetColor))
		adminPass = NewScan()
		check = CheckAdminUser(adminLog, adminPass)
		if check != true {
			fmt.Println(set.Red + set.AccesDenied + set.ResetColor)
		} else {
			GetCliMenu()
		}
		time.Sleep(3 * time.Second)
		fmt.Println(set.Descr)
		GetStart()
	} else {
		if isNumbers {
			var usr = new(User)
			datetime := TmFormat()
			dt := datetime.Format("02.01.2006")
			tm := datetime.Format("15:04")
			fmt.Printf("Дата: - %s\nВремя: - %s", dt, tm)
			whoEntered := usr.CheckInTimeValidation(emplID)
			usr.AdCheckinToDb(whoEntered[1], datetime)
			//Check(whoEntered[1], whoEntered[0], whoEntered[2], dt, tm) // Функция оповещения через бота !Допилить чтобы ее работа не мешала CLI
			time.Sleep(2 * time.Second)
			GetStart()
		} else {
			fmt.Println(string(set.Red), set.ERROR_INPUT, string(set.ResetColor))
			time.Sleep(2 * time.Second)
			GetStart()
		}
	}

}
