package main

import (
	"4eker/pkg"
	"4eker/set"
	"flag"
	"fmt"
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
		tgbotapi.NewKeyboardButton("📆 Журнал посещения за период"), // За период
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("↩️-Назад"),
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
		tgbotapi.NewKeyboardButton("↩️-Назад"),
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
		tgbotapi.NewKeyboardButton("↩️-Назад"),
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
		tgbotapi.NewKeyboardButton("↩️-Назад"),
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
	var token string
	flag.StringVar(&token, "apikey", "---", "Токен телеграм бота")

	flag.Parse()
	if token == "---" {
		fmt.Println("Введите токен!")
	}

	bot, _ := tgbotapi.NewBotAPI(token)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	log.Printf("Запуск бота '%s' успешно выполнен!", bot.Self.UserName)
	log.Printf("ID чата")
	log.Printf("ChatID%s", updates)

	usr := new(pkg.User)
	supUser := new(pkg.SuperUser)
	jrnl := new(pkg.Journal)
	for update := range updates {
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msgGenaral := update.Message.Text
			log.Println("LOG -> ", msgGenaral)
			switch update.Message.Text {
			case "/startmenu": // ГОТОВО!
				keys := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				keys.ReplyMarkup = nmShowVisiters
				bot.Send(keys)
				break
			case "👁 Кто в цеху": // ГОТОВО!
				resOut := jrnl.WhoInPlaceForBot()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, resOut)
				bot.Send(msg)
			case "📊 Отчеты": // ГОТОВО!
				keys := tgbotapi.NewMessage(update.Message.Chat.ID, "Включил --->"+update.Message.Text)
				keys.ReplyMarkup = nmKeyJournal
				bot.Send(keys)
			case "🛠 Настройки": // ГОТОВО!
				var bln bool
				msg.Text = set.BOT_WRNING_ADMIN_INVATION
				bot.Send(msg)
				for upd := range updates {
					msgIn := upd.Message.Text
					if msgIn == "↩️-Назад" {
						break
					} else {
						res := pkg.NumberValuator(msgIn)
						if len(res) == 2 {
							bln = pkg.CheckAdminUser(res[0], res[1])
							if bln != true {
								msg := tgbotapi.NewMessage(update.Message.Chat.ID, set.BOT_ADMIN_ACCESS_BAD)
								bot.Send(msg)
								break
							} else {
								keys := tgbotapi.NewMessage(update.Message.Chat.ID, set.BOT_WARNING_ADMIN_ACCESS)
								keys.ReplyMarkup = nmKeySettings
								bot.Send(keys)
								break
							}
						} else {
							msg.Text = set.BOT_WARNING_ARGUMENTS_NOT_ENOUGH + "2 значения"
							bot.Send(msg)
						}
					}
				}
			case "👷🏽 Сотрудники": // ГОТОВО!
				keys := tgbotapi.NewMessage(update.Message.Chat.ID, "Меню --->"+update.Message.Text)
				keys.ReplyMarkup = nmKeyEmpl
				bot.Send(keys)
			case "🦸🏻 Администраторы": // ГОТОВО!
				keys := tgbotapi.NewMessage(update.Message.Chat.ID, "Меню --->"+update.Message.Text)
				keys.ReplyMarkup = nmKeyAdmin
				bot.Send(keys)
			case "↩️-Назад": // ГОТОВО!
				keys := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				keys.ReplyMarkup = nmShowVisiters
				bot.Send(keys)
			case "➕ Добавить сотрудника": // ГОТОВО
				msg.Text = set.BOT_WARNING_ADDUSER_INFO
				bot.Send(msg)
				for upd := range updates {
					msgIn := upd.Message.Text
					if msgIn == "↩️-Назад" {
						break
					} else {
						res := pkg.NumberValuator(msgIn)
						var r string
						if len(res) == 4 {
							r = usr.AddInBot(res[0], res[1], res[2], res[3])
							msg.Text = r
							bot.Send(msg)
							break
						} else {
							msg.Text = set.BOT_WARNING_ARGUMENTS_NOT_ENOUGH + "4 значения"
							bot.Send(msg)
						}
					}
				}
			case "❌ Удалить сотрудника": // ГОТОВО
				msg.Text = set.BOT_WARNING_DELUSER_INFO
				bot.Send(msg)
				for upd := range updates {
					msgIn := upd.Message.Text
					if msgIn == "↩️-Назад" {
						break
					} else {
						res := pkg.NumberValuator(msgIn)
						var r string
						if len(res) == 1 {
							r = usr.DeleteRowInBot(res[0])
							msg.Text = r
							bot.Send(msg)
						} else {
							msg.Text = set.BOT_WARNING_ARGUMENTS_NOT_ENOUGH + "1 значение"
							bot.Send(msg)
						}
					}
				}
			case "✏️ Редактировать сотрудника": // ГОТОВО
				msg.Text = set.BOT_WARNING_EDITUSER_INFO
				bot.Send(msg)
				preShow := usr.ShowAllInBot()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, preShow)
				bot.Send(msg)
				for upd := range updates {
					msgIn := upd.Message.Text
					if msgIn == "↩️-Назад" {
						break
					} else {
						res := pkg.NumberValuator(msgIn)
						var r, show string
						if len(res) == 4 {
							r, show = usr.EditFromBot(res[0], res[1], res[2], res[3])
							msg.Text = r + "\n" + show
							bot.Send(msg)
							break
						} else {
							msg.Text = set.BOT_WARNING_ARGUMENTS_NOT_ENOUGH + "4 значение"
							bot.Send(msg)
						}
					}
				}
			case "🗂 Список сотрудников": // ГОТОВО
				resOut := usr.ShowAllInBot()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, resOut)
				bot.Send(msg)
			case "➕ Добавить администратора": // ГОТОВО
				msg.Text = set.BOT_WARNING_ADDADMIN_INFO
				bot.Send(msg)
				for upd := range updates {
					msgIn := upd.Message.Text
					if msgIn == "↩️-Назад" {
						break
					} else {
						res := pkg.NumberValuator(msgIn)
						var r string
						if len(res) == 3 {
							r = supUser.AddInBot(res[0], res[1], res[2])
							msg.Text = r
							bot.Send(msg)
						} else {
							msg.Text = set.BOT_WARNING_ARGUMENTS_NOT_ENOUGH + "3 значения"
							bot.Send(msg)
						}
					}
				}
			case "❌ Удалить администратора": // ГОТОВО
				msg.Text = set.BOT_WARNING_DELADMIN_INFO
				bot.Send(msg)
				for upd := range updates {
					msgIn := upd.Message.Text
					if msgIn == "↩️-Назад" {
						break
					} else {
						res := pkg.NumberValuator(msgIn)
						var r string
						if len(res) == 1 {
							r = supUser.DeleteRowInBot(res[0])
							msg.Text = r
							bot.Send(msg)
						} else {
							msg.Text = set.BOT_WARNING_ARGUMENTS_NOT_ENOUGH + "1 значение"
							bot.Send(msg)
						}
					}
				}
			case "✏️ Редактировать администратора": // ГОТОВО!
				msg.Text = set.BOT_WARNING_EDITADMIN_INFO
				bot.Send(msg)
				preShow := supUser.ShowAllInBot()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, preShow)
				bot.Send(msg)
				for upd := range updates {
					msgIn := upd.Message.Text
					if msgIn == "↩️-Назад" {
						break
					} else {
						res := pkg.NumberValuator(msgIn)
						var r, show string
						if len(res) == 4 {
							r, show = supUser.EditFromBot(res[0], res[1], res[2], res[3])
							msg.Text = r + "\n" + show
							bot.Send(msg)
							break
						} else {
							msg.Text = set.BOT_WARNING_ARGUMENTS_NOT_ENOUGH + "4 значение"
							bot.Send(msg)
						}
					}
				}
			case "🗂 Список администраторов": //ГОТОВО
				resOut := supUser.ShowAllInBot()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, resOut)
				bot.Send(msg)
			case "📖 Журнал посещений по сотруднику за период": // ГОТОВО!
				msg.Text = set.BOT_WARNING_EXPORT_USER_FILE_INFO
				bot.Send(msg)
				for upd := range updates {
					msgIn := upd.Message.Text
					if msgIn == "↩️-Назад" {
						break
					} else {
						res := pkg.NumberValuator(msgIn)
						if len(res) == 3 {
							pkg.PresentJournalToEmpl(res[0], res[1], res[2])
							file := tgbotapi.FilePath(set.ExcelFile)
							msg := tgbotapi.NewDocument(update.Message.Chat.ID, file)
							bot.Send(msg)
							break
						} else {
							msg.Text = set.BOT_WARNING_ARGUMENTS_NOT_ENOUGH + "3 значения"
							bot.Send(msg)
						}
					}
				}
			case "📆 Журнал посещения за период": // ГОТОВО!
				msg.Text = set.BOT_WARNING_EXPORT_FILE_INFO
				bot.Send(msg)
				for upd := range updates {
					msgIn := upd.Message.Text
					if msgIn == "↩️-Назад" {
						break
					} else {
						res := pkg.NumberValuator(msgIn)
						if len(res) == 2 {
							pkg.PresentJournalToDay(res[0], res[1])
							file := tgbotapi.FilePath(set.ExcelFile)
							msg := tgbotapi.NewDocument(update.Message.Chat.ID, file)
							bot.Send(msg)
						} else {
							msg.Text = set.BOT_WARNING_ARGUMENTS_NOT_ENOUGH + "2 значение"
							bot.Send(msg)
						}
					}
				}
			}
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
