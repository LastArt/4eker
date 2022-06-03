package pkg

import (
	"4eker/set"
	"fmt"
)

func GetSettings() {
	fmt.Println(string(set.Yellow), set.DescrSettings, string(set.ResetColor))
	var choice string
	fmt.Scanln(&choice)

	switch choice {
	case "1": // Добавить нового сотрудника
		var usr = new(User)
		//var do string
		usr.Add()
		usr.ShowAll()
	L2:
		Question()
		goto L2
	case "2": // Удалить сотрудника
		var usrDel = new(User)
		usrDel.ShowAll()
	L3:
		fmt.Println(set.Yellow, "Введите № карты сотрудника которого нужно удалить:\n", set.ResetColor)
		fio := NewScan()
		usrDel.DeleteRow(fio)
		fmt.Println(set.Yellow, "\nПроверьте удален ли сотрудник из базы, если нет, повторите попытку:\n", set.ResetColor)
		usrDel.ShowAll()
		Question()
		goto L3
	case "3": // Редактировать сотрудника
		var usr = new(User)
		usr.ShowAll()
		fmt.Println(set.Yellow, "Введите № Карты сотрудника которого нужно редактировать:\n", set.ResetColor)

	case "4": // Вывести список сотрудников
		var usr = new(User)
		usr.ShowAll()
		GetSettings()
	case "5": // Вывести журнал присутсвующих на текущий час на местах
		var jrnl = new(Journal)
		jrnl.WhoInPlace()

	case "6": // Подключение к БД (делаем новое подключение к БД через создание таблиц в том числе)
		fmt.Println("Подключение к БД")
	case "7": // Очистка записей из таблиц
		fmt.Println("Очистка записей БД")
	case "8": // Бэкап БД
		fmt.Println("Бэкап БД")
	case "9":
		var supuser = new(SuperUser)
		supuser.Add()
		supuser.ShowAll()
	case "10":
		fmt.Println("Редактировать суперпользователя")
	case "11":
		var supuser = new(SuperUser)
		supuser.ShowAll()
		fmt.Println(set.Yellow, "Введите Логин сотрудника которого нужно удалить:\n", set.ResetColor)
		login := NewScan()
		supuser.DeleteRow(login)
		fmt.Println(set.Yellow, "\nПроверьте удален ли суперпользователь из базы, если нет, повторите попытку:\n", set.ResetColor)
		supuser.ShowAll()
	case "12":
		// Показать всех суперпользователей
		var spusr = new(SuperUser)
		spusr.ShowAll()
	case "13":
		fmt.Println("Настройка пути для файла tokena")
	}
}
