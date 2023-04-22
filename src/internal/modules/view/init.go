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

func singUp(bot *tgbotapi.BotAPI, chatId int64, update tgbotapi.Update, controller *controllers.Controller) string {
	var reply = ""

	msg := tgbotapi.NewMessage(chatId, "Введите свою фамилию")
	update
	msg = tgbotapi.NewMessage(chatId, "Введите свою специальность")

	msg = tgbotapi.NewMessage(chatId, "Введите свою роль")

	newDoctor := &auth_dto.NewDoctor{
		TokenId:    strconv.FormatInt(chatId, 10),
		Surname:    surname,
		Speciality: speciality,
		Role:       role,
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

	for update := range updates {
		if update.Message != nil {

			ChatId := update.Message.Chat.ID
			msg := tgbotapi.NewMessage(ChatId, update.Message.Text)

			switch update.Message.Text {
			case "help":
				msg.Text = "Type /singup"
			case "singup":
				msg.Text = singUp(bot, ChatId, updates, controller)
			case "status":
				msg.Text = "I'm ok."
			case "open":
				msg.ReplyMarkup = numericKeyboard
			case "close":
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			default:
				msg.Text = "Ты кринж"
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
