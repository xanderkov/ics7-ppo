package view

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	auth_dto "hospital/internal/modules/domain/auth/dto"
	"hospital/internal/modules/view/controllers"
	"log"
	"os"
	"strconv"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("singup", "singup"),
		tgbotapi.NewInlineKeyboardButtonData("help", "help"),
		tgbotapi.NewInlineKeyboardButtonData("status", "status"),
	),
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func singUp(chatId int64, user *UsersMessage) string {
	addNextMessages("Введите свою фамилию", user, chatId)
	addNextMessages("Введите свою специальность", user, chatId)
	addNextMessages("Введите свою роль", user, chatId)
	msg := printNewMessage(user.NextMessages)
	return msg
}

func printNewMessage(nextMessages []string) string {
	msg := nextMessages[0]
	nextMessages = remove(nextMessages, 0)
	return msg
}

func handleUsers(Users []UsersMessage, chatId int64, userMessage string, controller *controllers.Controller) string {
	var msg string
	for _, u := range Users {
		if u.ChatId == chatId {
			u.UserMessages = append(u.UserMessages, userMessage)
			if len(u.NextMessages) > 0 {
				msg = printNewMessage(u.NextMessages)
			} else {
				msg = endSingUp(&u, chatId, controller)
			}
		}
	}
	return msg
}

func endSingUp(user *UsersMessage, chatId int64, controller *controllers.Controller) string {
	var reply string

	newDoctor := &auth_dto.NewDoctor{
		TokenId:    strconv.FormatInt(chatId, 10),
		Surname:    user.UserMessages[0],
		Speciality: user.UserMessages[1],
		Role:       user.UserMessages[2],
	}
	_, err := controller.SingUp(context.Background(), newDoctor)
	if err != nil {
		reply = "Уже зарегистрированы"
		return reply
	}
	reply = "Зарегистрирован"

	return reply
}

func handleBot(
	controller *controllers.Controller,
	updates tgbotapi.UpdatesChannel,
	bot *tgbotapi.BotAPI) {

	// ToDO:Сделать массив юзеров (это должна быть структура)
	// добавить обработку этого массива и в зависимости от поля запросы делать выводы.

	var Users []UsersMessage

	for update := range updates {
		if update.Message != nil {

			ChatId := update.Message.Chat.ID
			msg := tgbotapi.NewMessage(ChatId, update.Message.Text)

			if len(Users) > 0 {
				msg.Text = handleUsers(Users, ChatId, update.Message.Text, controller)
			} else {
				switch update.Message.Text {
				case "help":
					msg.Text = "Type /singup"
				case "singup":
					Users = append(Users, UsersMessage{ChatId: ChatId, Command: update.Message.Text})
					msg.Text = singUp(ChatId, &Users[len(Users)-1])
				case "status":
					msg.Text = "I'm ok."
				case "open":
					msg.ReplyMarkup = numericKeyboard
				case "close":
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				default:
					msg.Text = "Ты кринж"
				}
			}

			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		} else if update.CallbackQuery != nil {
			// Respond to the callback query, telling Telegram to show the user
			// a message with the data received.
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			// And finally, send a message containing the data received.
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}

func startBot(controller *controllers.Controller) {
	dotenv := goDotEnvVariable("TELEGRAM_APITOKEN")

	bot, err := tgbotapi.NewBotAPI(dotenv)
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	handleBot(controller, updates, bot)
}
