package set

var Headlogo string = `
++++++++++++++++++++++++++++++++++++++++++++++++++++
`
var Logo string = `
__    __  ________  __    __  ________  _______  
/  |  /  |/        |/  |  /  |/        |/       \ 
$$ |  $$ |$$$$$$$$/ $$ | /$$/ $$$$$$$$/ $$$$$$$  |
$$ |__$$ |$$ |__    $$ |/$$/  $$ |__    $$ |__$$ |
$$    $$ |$$    |   $$  $$<   $$    |   $$    $$/ 
$$$$$$$$ |$$$$$/    $$$$$  \  $$$$$/    $$$$$$$/  
      $$ |$$ |_____ $$ |$$  \ $$ |_____ $$ |      
      $$ |$$       |$$ | $$  |$$       |$$ |      
      $$/ $$$$$$$$/ $$/   $$/ $$$$$$$$/ $$/       
`
var Sublogo string = `
+++++++++++++++++++++++++++++++++++++++++++++++++++
+     Система учета рабочего времени  вер:1.0     +
+++++++++++++++++++++++++++++++++++++++++++++++++++
`

var Acces string = `
+++++++++++++++++++++++++++++++++++++++++++++++++++
+                 ДОСТУП РАЗРЕШЕН                 +
+                 для номера карты                +
+++++++++++++++++++++++++++++++++++++++++++++++++++
`
var AccesDenied string = `
+++++++++++++++++++++++++++++++++++++++++++++++++++
+                 ДОСТУП ЗАПРЕЩЕН                 +
+                 для номера карты                +
+++++++++++++++++++++++++++++++++++++++++++++++++++
`

var Descr string = `
+++++++++++++++++++++++++++++++++++++++++++++++++++
+                                                 +
+           ДЛЯ ФИКСАЦИИ РАБОЧЕГО ВРЕМЕНИ         +
+           КОСНИТЕСЬ КАРТОЙ СЧИТЫВАТЕЛЯ          +
+                                                 +   
+++++++++++++++++++++++++++++++++++++++++++++++++++
`
var DescrSettings string = `
+++++++++++++++++++ МЕНЮ ++++++++++++++++++++

-----------------Сотрудники------------------ 

1. 	Добавить нового сотрудника
2. 	Удалить сотрудника
3. 	Редактировать сотрудника
4. 	Вывести список сотрудников
5. 	Вывести журнал присутсвующих сейчас

---------------------------------------------

---------------Администраторы----------------

6. 	Добавить суперпользователя
7. 	Редактировать суперпользователя
8. 	Удалить суперпользователя
9. 	Показать список суперпользователей

---------------------------------------------

Выберите пункт меню...
`

var (
	ResetColor string = "\033[0m"
	Red        string = "\033[31m"
	Green      string = "\033[32m"
	Yellow     string = "\033[33m"
	Blue       string = "\033[34m"
	Purple     string = "\033[35m"
	Cyan       string = "\033[36m"
	White      string = "\033[37m"
)

var TokenFile = "/home/lastart/bot_token.json"
var ExcelFile = "./export.xlsx"

var (
	ERROR_SEND_BOT    string = Red + "ОШИБКА ОТПРАВКИ СООБЩЕНИЯ БОТУ: " + ResetColor
	ERROR_TOKEN       string = Red + "ОШИБКА ПОЛУЧЕНИЯ ТОКЕНА: " + ResetColor
	BAD_ADMIN_ACCESS  string = Red + "НЕВЕРНЫЙ КОД ДОСТУПА" + ResetColor
	ERROR_INPUT       string = Red + "ОШИБКА ВВОДА. НЕКОРРЕКТНЫЕ ДАННЫЕ" + ResetColor
	ERROR_INSERT_TODB string = Red + "ОШИБКА ЗАПИСИ В БАЗУ ДАННЫХ: " + ResetColor
)

// БОТ! Ошибки
var (
	BOT_ERROR_DELETE_USER string = "⛔️ Ошибка выполнения удаления из БД\nПопробуйте еще раз!"
	BOT_ERROR_ADDTODB     string = "⛔️ Ошибка выполнения записи в БД\nПопробуйте еще раз!"
	BOT_ADMIN_ACCESS_BAD  string = "⛔️ Доступ запрещен!\nНеверный логин или пароль!"
	BOT_ERROR_EDITUSER    string = "⛔️ Ошибка выполнения изменения записи в БД\nПопробуйте еще раз!"
)

// БOT! Уведомления
var (
	BOT_WARNING_DELETE_FINE           string = "✅Удаление прошло успешно!\nДля выхода из режима нажмите '↩️-Назад' или продолжайте удаление пользователей из БД!"
	BOT_WARNING_ADD_FINE              string = "✅Запись прошла успешно!\nДля выхода из режима нажмите '↩️-Назад' или продолжайте заполнение базы данных!"
	BOT_WARNING_ADMIN_ACCESS          string = "✅ Доступ открыт!"
	BOT_WARNING_EDITUSER_FINE         string = "✅Изменения прошли успешно!\nДля выхода из режима нажмите '↩️-Назад' или продолжайте вносить изменения в  базу данных!"
	BOT_WARNING_ARGUMENTS_NOT_ENOUGH  string = "⚠️Неверное количество аргументов для записи!\nПроверьте корректность и количество вводимых значений! \nОжидается "
	BOT_WRNING_ADMIN_INVATION         string = "🔐 АВТОРИЗАЦИЯ\nДля доступа в раздел настроек введите логин и пароль \nПример: Admin/qwerty123"
	BOT_WARNING_ADDUSER_INFO          string = "ℹ️Для добавления нового сотрудника, введите данные в следующей последовательности \nНомер карты/ФИО Сотрудника/Должность сотрудника/Зарплата сотрудника через '/'\nПример: 485548845/Иванов Иван Иванович/Слесарь/45000"
	BOT_WARNING_DELUSER_INFO          string = "ℹ️Для того чтобы удалить администратора введите ФИО удаляемого сотрудника \nПример: Иванов Иван Иванович"
	BOT_WARNING_ADDADMIN_INFO         string = "ℹ️Для добавления нового Администратора, введите данные в следующей последовательности \nЛогин/Пароль/Почта через '/'\nПример: Admin/qwerty123/admin@mail.ru"
	BOT_WARNING_DELADMIN_INFO         string = "ℹ️Для того чтобы удалить администратора введите логин \nПример: Admin"
	BOT_WARNING_EDITUSER_INFO         string = "ℹ️Для того чтобы редактировать сотрудника, введите № карты сотрудника которого нужно изменить, а следом значения на которые надо изменить данные.\nПример: 48858485/Петров Петр Петрович/Слесарь/56000"
	BOT_WARNING_EXPORT_USER_FILE_INFO string = "ℹ️ Укажите № Карты сотрудника и период за который нужно выгрузить журнал\nПример записи: 485545884/01.05.2022/13.05.2022"
	BOT_WARNING_EXPORT_FILE_INFO      string = "ℹ️ Укажите период за который нужно выгрузить журнал \nПример записи: 01.05.2022/13.05.2022"
	BOT_WARNING_EDITADMIN_INFO        string = "ℹ️Для того чтобы редактировать, введите № карты администратора которого нужно изменить, а следом значения на которые надо изменить данные.\nПример: 21/AdminNew/PasswordNew/newmail@example.ru"
)
