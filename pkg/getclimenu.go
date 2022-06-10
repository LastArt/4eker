package pkg

import (
	"4eker/set"
	"fmt"
)

func GetCliMenu() {
	fmt.Println(string(set.Yellow), set.DescrSettings, string(set.ResetColor))
	var choice string
	fmt.Scanln(&choice)

	switch choice {
	case "1": // Добавить нового сотрудника
		var usr = new(User)
		//var do string
		usr.Add()
		usr.ShowAll()
		Question()
	case "2": // Удалить сотрудника
		var usrDel = new(User)
		usrDel.ShowAll()
		fmt.Println(set.Yellow, "Введите № карты сотрудника которого нужно удалить:\n", set.ResetColor)
		fio := NewScan()
		usrDel.DeleteRow(fio)
		fmt.Println(set.Yellow, "\nПроверьте удален ли сотрудник из базы, если нет, повторите попытку:\n", set.ResetColor)
		usrDel.ShowAll()
		Question()
	case "3": // Редактировать сотрудника
		var usrEdit = new(User)
		usrEdit.ShowAll()
		fmt.Println(set.Yellow, "Введите № Карты сотрудника которого нужно редактировать:\n", set.ResetColor)
		cardnum := NewScan()
		usrEdit.Edit(cardnum)
		usrEdit.ShowAll()
		Question()
	case "4": // Вывести список сотрудников
		var usr = new(User)
		usr.ShowAll()
		GetCliMenu()
	case "5": // Вывести журнал присутсвующих на текущий день (а лучше час)
		var jrnl = new(Journal)
		jrnl.WhoInPlace()
		GetCliMenu()
	case "6": // Добавить суперпользователя
		var supuser = new(SuperUser)
		supuser.Add()
		supuser.ShowAll()
		Question()
	case "7": // Редактировать суперпользователя
		var supuser = new(SuperUser)
		supuser.ShowAll()
		fmt.Println(set.Yellow, "Введите Id суперпользователя которого нужно редактировать:\n", set.ResetColor)
		id := NewScan()
		supuser.Edit(id)
		supuser.ShowAll()
		Question()
	case "8": // Удалить суперпользователя
		var sup = new(SuperUser)
		sup.ShowAll()
		fmt.Println(set.Yellow, "Введите Id суперпользователя которого нужно удалить:\n", set.ResetColor)
		id := NewScan()
		sup.DeleteRow(id)
		fmt.Println(set.Yellow, "\nПроверьте удален ли сотрудник из базы, если нет, повторите попытку:\n", set.ResetColor)
		sup.ShowAll()
		Question()
	case "9": // Показать суперпользователей
		var supuser = new(SuperUser)
		supuser.ShowAll()
	default:
		fmt.Println("Команда не распознана. Попробуйте еще раз или нажмите В-[Выход] для возврата в учет рабочего времени")
		x := NewScan()
		if x == "В" || x == "в" || x == "Выход" || x == "выход" {
			GetStart()
		} else {
			GetCliMenu()
		}

	}
}
