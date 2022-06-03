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

	fmt.Println(string(set.Yellow), set.Headlogo, string(set.ResetColor))
	fmt.Println(string(set.Blue), set.Logo)
	fmt.Println(string(set.Yellow), set.Sublogo, string(set.ResetColor))
	fmt.Println(set.Descr)
L1:
	emplID = NewScan()

	isNumbers := reg.MatchString(emplID)
	if emplID == "--admin" {
		fmt.Println(string(set.Yellow), "Введите логин администратора", string(set.ResetColor))
		adminLog = NewScan()
		fmt.Println(string(set.Yellow), "Введите пароль администратора", string(set.ResetColor))
		adminPass = NewScan()
		check = CheckAdminUser(adminLog, adminPass) //! Проверить на предмет обновления принципа проверки!!!
		if check != true {
			fmt.Println(set.Red + set.AccesDenied + set.ResetColor)
		} else {
			GetSettings()
		}
		// Запускаем отдельную функцию с кейсами по настройки и управлению GetSettings()
		time.Sleep(3 * time.Second)
		fmt.Println(set.Descr)
		goto L1 // Оставляю рекурсивно, но возможно перейду на goto
	} else {
		if isNumbers {
			var usr = new(User)
			//res, _ = strconv.Atoi(dataTime)
			datetime := TmFormat()
			dt := datetime.Format("02.01.2006")
			tm := datetime.Format("15:04")
			fmt.Printf("dtRES : - %s\ntmRES : - %s", dt, tm)
			whoEntered := usr.CheckInTimeValidation(emplID)
			usr.AdCheckinToDb(whoEntered[1], datetime)
			Check(whoEntered[1], whoEntered[0], whoEntered[2], dt, tm)
			goto L1
		} else {
			fmt.Println(string(set.Red), set.ERROR_INPUT, string(set.ResetColor))
			goto L1
		}
	}
	//return res
}
