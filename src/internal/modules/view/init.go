package view

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"hospital/internal/modules/db/ent"
	auth_dto "hospital/internal/modules/domain/auth/dto"
	"hospital/internal/modules/view/controllers"
	"log"
	"os"
	"strconv"
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
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
	surname := update.Message.Text

	msg = tgbotapi.NewMessage(chatId, "Введите свою специальность")
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
	speciality := update.Message.Text

	msg = tgbotapi.NewMessage(chatId, "Введите свою роль")
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
	role := update.Message.Text

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

func handleBot(client *ent.Client,
	controller *controllers.Controller,
	updates tgbotapi.UpdatesChannel,
	bot *tgbotapi.BotAPI) {

	for update := range updates {
		if update.Message == nil {
			continue
		}

		ChatId := update.Message.Chat.ID
		msg := tgbotapi.NewMessage(ChatId, update.Message.Text)
		text := update.Message.Text
		reply := ""

		if text == "/singup" {
			reply = singUp(bot, ChatId, update, controller)
		} else {
			reply = "Ты кринж"
		}

		msg = tgbotapi.NewMessage(ChatId, reply)

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}
}

func startBot(client *ent.Client, controller *controllers.Controller) {
	dotenv := goDotEnvVariable("TELEGRAM_APITOKEN")

	bot, err := tgbotapi.NewBotAPI(dotenv)
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	handleBot(client, controller, updates, bot)
}
