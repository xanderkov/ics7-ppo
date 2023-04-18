package view

import (
	"context"
	"fmt"
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

func singUp(chatId int64, controller *controllers.Controller) string {
	var reply = ""
	newDoctor := &auth_dto.NewDoctor{
		TokenId:    strconv.FormatInt(chatId, 10),
		Surname:    "Kovel",
		Speciality: "Психотерапевт",
		Role:       "Глав врач",
	}
	doctor, err := controller.SingUp(context.Background(), newDoctor)
	if err != nil {
		reply = "Already registrated"
	}

	fmt.Println(doctor)
	reply = doctor.Surname

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
			reply = singUp(ChatId, controller)
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
