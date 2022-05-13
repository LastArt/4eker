package main

import (
	"kontroller/pkg"
	"kontroller/set"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// =================================================================================
// Кнопки для работы с ботом - Клавиатура
// =================================================================================
var nmShowVisiters = tgbotapi.NewReplyKeyboard( // Показывает журнал присутсвубщих
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("👁 Кто в цеху"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("📊 Отчеты"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🛠 Настройки"),
	),
)

var nmKeyJournal = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("📖 Журнал посещений по сотруднику за период"), // За период
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("💵 ЗП по сотруднику за период"), // За период
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("📆 Журнал посещения за период"), // За период
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🔙 Вернуться"),
	),
)
var nmKeySettings = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("👷🏽 Сотрудники"), //Открывает еще 3 кнопки nmKeyEmpl
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🦸🏻 Администраторы"), //Открывает еще 3 кнопки nmKeyAdmin
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🗄 База данных"), //Открывает еще 3 кнопки nmKeyDataBase
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🔙 Вернуться"),
	),
)

// ============= ВТОРОЙ УРОВЕНЬ КНОПОК nmKeySettings=========================
var nmKeyEmpl = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("➕ Добавить сотрудника"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("❌ Удалить сотрудника"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("✏️ Редактировать сотрудника"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🗂 Список сотрудников"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🔙 Вернуться"),
	),
)

var nmKeyAdmin = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("➕ Добавить администратора"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("❌ Удалить администратора"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("✏️ Редактировать администратора"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🗂 Список администраторов"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🔙 Вернуться"),
	),
)

var nmKeyDataBase = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🔌 Создать новое подключение"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🧹 Очистить базу данных"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("📦 Сделать бекап БД"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("⚙️ Посмотреть настройки подключения"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🔙 Вернуться"),
	),
)

// ============== ИНЛАЙН КНОПКИ ============================
var exitKey = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Выход", "exit"),
	),
)

//=============== КОНЕЦ БЛОКА С ИНЛАЙН КНОПКАМИ =============

func main() {
	bot, _ := tgbotapi.NewBotAPI(pkg.GetKey(set.TokenFile))

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	log.Printf("Authorized on account %s", bot.Self.UserName)
	//log.Printf("ChatID%s", ch)

	usr := new(pkg.User)
	supUser := new(pkg.SuperUser)
	jrnl := new(pkg.Journal)
	for update := range updates {

		if update.Message != nil {

			//msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введена команда:"+update.Message.Text)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			// SWITCH.
			msgGenaral := update.Message.Text
			log.Println("msgGeneral -> ", msgGenaral)
			switch update.Message.Text {
			case "/start":
				keys := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				keys.ReplyMarkup = nmShowVisiters
				bot.Send(keys)
			case "👁Кто в цеху":
				resOut := jrnl.WhoInPlaceForBot()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, resOut)
				bot.Send(msg)
			case "📊 Отчеты":
				keys := tgbotapi.NewMessage(update.Message.Chat.ID, "Включил --->"+update.Message.Text)
				keys.ReplyMarkup = nmKeyJournal
				bot.Send(keys)
			case "🛠 Настройки":
				var bln bool
				var str string
				msg.Text = "🔐 Для доступа в раздел настроек введите логин и пароль \nПример: Admin/qwerty123"
				bot.Send(msg)
				for upd := range updates {
					msgIn := upd.Message.Text
					if msgIn == "🔙 Вернуться" {
						break
					} else {
						res := pkg.NumberValuator(msgIn)
						if len(res) == 2 {
							bln, str = pkg.CheckAdminUser(res[0], res[1])
							if bln != true {
								msg := tgbotapi.NewMessage(update.Message.Chat.ID, str)
								bot.Send(msg)
							} else {
								keys := tgbotapi.NewMessage(update.Message.Chat.ID, str)
								keys.ReplyMarkup = nmKeySettings
								bot.Send(keys)
							}

						} else {
							msg.Text = "⚠️Неверное количество аргументов для записи!\nПроверьте корректность внесенной информации!\nТребуется 1 значение"
							bot.Send(msg)
						}
					}
				}

			case "👷🏽 Сотрудники":
				keys := tgbotapi.NewMessage(update.Message.Chat.ID, "Включил --->"+update.Message.Text)
				keys.ReplyMarkup = nmKeyEmpl
				bot.Send(keys)
			case "🦸🏻 Администраторы":
				keys := tgbotapi.NewMessage(update.Message.Chat.ID, "Включил --->"+update.Message.Text)
				keys.ReplyMarkup = nmKeyAdmin
				bot.Send(keys)
			case "🗄 База данных":
				keys := tgbotapi.NewMessage(update.Message.Chat.ID, "Включил --->"+update.Message.Text)
				keys.ReplyMarkup = nmKeyDataBase
				bot.Send(keys)
			case "🔙 Вернуться":
				keys := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				keys.ReplyMarkup = nmShowVisiters
				bot.Send(keys)
			// ТУТ ИДЕТ СЕРИЯ КЕЙСОВ ПО ОБРАБОТКЕ КОНЕЧНЫХ ПУНКТОВ МЕНЮ (Конечных кнопок!)
			case "➕ Добавить сотрудника": // ГОТОВО
				msg.Text = "Для добавления нового сотрудника, введите данные в следующей последовательности \nНомер карты/ФИО Сотрудника/Должность сотрудника/Зарплата сотрудника через '/'\nПример: 485548845/Иванов Иван Иванович/Слесарь/45000"
				bot.Send(msg)
				for upd := range updates {
					msgIn := upd.Message.Text
					if msgIn == "🔙 Вернуться" {
						break
					} else {
						res := pkg.NumberValuator(msgIn)
						var r string
						if len(res) == 4 {
							r = usr.AddInBot(res[0], res[1], res[2], res[3])
							msg.Text = r
							bot.Send(msg)
						} else {
							msg.Text = "⚠️Неверное количество аргументов для записи!\nПроверьте корректность внесенной информации!\nТребуется 4 значения"
							bot.Send(msg)
						}
					}
				}
			case "❌ Удалить сотрудника": // ГОТОВО
				msg.Text = "Для того чтобы удалить администратора введите ФИО удаляемого сотрудника \nПример: Иванов Иван Иванович"
				bot.Send(msg)
				for upd := range updates {
					msgIn := upd.Message.Text
					if msgIn == "🔙 Вернуться" {
						break
					} else {
						res := pkg.NumberValuator(msgIn)
						var r string
						if len(res) == 1 {
							r = usr.DeleteRowInBot(res[0])
							msg.Text = r
							bot.Send(msg)
						} else {
							msg.Text = "⚠️Неверное количество аргументов для записи!\nПроверьте корректность внесенной информации!\nТребуется 1 значение"
							bot.Send(msg)
						}
					}
				}
			case "✏️ Редактировать сотрудника": // TODO Доделать
			case "🗂 Список сотрудников": // ГОТОВО
				resOut := usr.ShowAllInBot()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, resOut)
				bot.Send(msg)
			case "➕ Добавить администратора": // ГОТОВО
				msg.Text = "Для добавления нового Администратора, введите данные в следующей последовательности \nЛогин/Пароль/Почта через '/'\nПример: Admin/qwerty123/admin@mail.ru"
				bot.Send(msg)
				for upd := range updates {
					msgIn := upd.Message.Text
					if msgIn == "🔙 Вернуться" {
						break
					} else {
						res := pkg.NumberValuator(msgIn)
						var r string
						if len(res) == 3 {
							r = supUser.AddInBot(res[0], res[1], res[2])
							msg.Text = r
							bot.Send(msg)
						} else {
							msg.Text = "⚠️Неверное количество аргументов для записи!\nПроверьте корректность внесенной информации!\nТребуется 4 значения"
							bot.Send(msg)
						}
					}
				}
			case "❌ Удалить администратора": // ГОТОВО
				msg.Text = "Для того чтобы удалить администратора введите логин \nПример: Admin"
				bot.Send(msg)
				for upd := range updates {
					msgIn := upd.Message.Text
					if msgIn == "🔙 Вернуться" {
						break
					} else {
						res := pkg.NumberValuator(msgIn)
						var r string
						if len(res) == 1 {
							r = supUser.DeleteRowInBot(res[0])
							msg.Text = r
							bot.Send(msg)
						} else {
							msg.Text = "⚠️Неверное количество аргументов для записи!\nПроверьте корректность внесенной информации!\nТребуется 1 значение"
							bot.Send(msg)
						}
					}
				}
			case "✏️ Редактировать администратора": // TODO Доделать
			case "🗂 Список администраторов": //ГОТОВО
				resOut := supUser.ShowAllInBot()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, resOut)
				bot.Send(msg)
			// * Подменю раздела "База данных"
			case "🔌 Создать новое подключение": // TODO Доделать
			case "🧹 Очистить базу данных": // TODO Доделать
			case "📦 Сделать бекап БД": // TODO Доделать
			case "⚙️Посмотреть настройки подключения": // TODO Доделать
			// * Подменю раздела "Отчеты"
			case "📖 Журнал посещений по сотруднику за период":
				msg.Text = "ℹ️ Укажите ФИО сотрудника и период за который нужно выгрузить журнал\nПример записи: 01.05.2022/13.05.2022/Иванов Иван Иванович"
				bot.Send(msg)
				for upd := range updates {
					msgIn := upd.Message.Text
					if msgIn == "🔙 Вернуться" {
						break
					} else {
						res := pkg.NumberValuator(msgIn)
						if len(res) == 3 {
							pkg.PresentJournalToEmpl(res[0], res[1], res[2])
							file := tgbotapi.FilePath(set.ExcelFile)
							msg := tgbotapi.NewDocument(update.Message.Chat.ID, file)
							bot.Send(msg)
						} else {
							msg.Text = "⚠️Неверное количество аргументов для записи!\nПроверьте корректность внесенной информации!\nТребуется 1 значение"
							bot.Send(msg)
						}
					}
				}
			case "💵 ЗП по сотруднику за период":
				msg.Text = "⚠️ Данный раздел находится в разработке!"
				bot.Send(msg)
			case "📆 Журнал посещения за период":
				msg.Text = "ℹ️ Укажите период за который нужно выгрузить журнал \nПример записи: 01.05.2022/13.05.2022"
				bot.Send(msg)
				for upd := range updates {
					msgIn := upd.Message.Text
					if msgIn == "🔙 Вернуться" {
						break
					} else {
						res := pkg.NumberValuator(msgIn)
						if len(res) == 2 {
							pkg.PresentJournalToDay(res[0], res[1])
							file := tgbotapi.FilePath(set.ExcelFile)
							msg := tgbotapi.NewDocument(update.Message.Chat.ID, file)
							bot.Send(msg)
						} else {
							msg.Text = "⚠️Неверное количество аргументов для записи!\nПроверьте корректность внесенной информации!\nТребуется 1 значение"
							bot.Send(msg)
						}
					}
				}
			}
			// Send the message.

			// if _, err := bot.Send(msg); err != nil {
			// 	panic(err)
			// }
		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}
			// msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			// if _, err := bot.Send(msg); err != nil {
			// 	panic(err)
			// }
			switch callback.Text {
			case "exit":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.Message.Text)
				bot.Send(msg)
			}
		}
	}
}
