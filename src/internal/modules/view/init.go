package view

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hospital/internal/modules/config"
	auth_dto "hospital/internal/modules/domain/auth/dto"
	"hospital/internal/modules/view/controllers"
	"strconv"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Зарегестрироваться"),
		tgbotapi.NewKeyboardButton("Помощь"),
		tgbotapi.NewKeyboardButton("Просмотреть данные о себе"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Удалить пациента"),
		tgbotapi.NewKeyboardButton("Добавить пациента"),
		tgbotapi.NewKeyboardButton("Посмотреть своих пациентов"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Найти палату по номеру"),
		tgbotapi.NewKeyboardButton("Вывести все палаты"),
	),
)

func singUp(chatId int64, user *UsersMessage) string {
	var msg string
	addNextMessages("Введите свою фамилию", user, chatId)
	addNextMessages("Введите свою специальность", user, chatId)
	addNextMessages("Введите свою роль", user, chatId)
	msg, user.NextMessages = printNewMessage(user.NextMessages)
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

func handleUsers(
	Users []UsersMessage, chatId int64,
	userMessage string,
	controller *controllers.Controller) (string, []UsersMessage) {

	var msg string
	for i, u := range Users {
		if u.ChatId == chatId {
			Users[i].UserMessages = append(Users[i].UserMessages, userMessage)
			if len(u.NextMessages) > 0 {
				msg, Users[i].NextMessages = printNewMessage(u.NextMessages)
			} else {
				msg = endSingUp(&Users[i], chatId, controller)
				Users = Users[:i+copy(Users[i:], Users[i+1:])]
			}
		}
	}
	return msg, Users
}

func getInfoAboutDoctor(id int64, controller *controllers.Controller) string {

	doctor, err := controller.DoctorToken(context.Background(), strconv.FormatInt(id, 10))
	if err != nil {
		msg := "Вы не зарегистрированы"
		return msg
	}
	msg := fmt.Sprintf("Фамилия: %s \nСпециальность: %s \nРоль: %s \n",
		doctor.Surname, doctor.Speciality, doctor.Role)
	return msg
}

func deletePatient(id int64) string {
	msg := fmt.Sprintf("Пустышка")
	return msg
}

func addPatient(id int64) string {
	msg := fmt.Sprintf("Пустышка")
	return msg
}

func getInfoAboutPatients(id int64) string {
	msg := fmt.Sprintf("Пустышка")
	return msg
}

func printByNum(id int64) string {
	msg := fmt.Sprintf("Пустышка")
	return msg
}

func printAllRooms(id int64) string {
	msg := fmt.Sprintf("Пустышка")
	return msg
}

func handleBot(
	controller *controllers.Controller,
	updates tgbotapi.UpdatesChannel,
	bot *tgbotapi.BotAPI) {

	var Users []UsersMessage

	for update := range updates {
		if update.Message != nil {

			ChatId := update.Message.Chat.ID
			msg := tgbotapi.NewMessage(ChatId, update.Message.Text)

			if len(Users) > 0 {
				msg.Text, Users = handleUsers(Users, ChatId, update.Message.Text, controller)
			} else {
				switch update.Message.Text {
				case "Помощь":
					msg.Text = "Напишите: Зарегестрироваться или open"
				case "Зарегестрироваться":
					Users = append(Users, UsersMessage{ChatId: ChatId, Command: update.Message.Text})
					msg.Text = singUp(ChatId, &Users[len(Users)-1])
				case "Просмотреть данные о себе":
					msg.Text = getInfoAboutDoctor(ChatId, controller)
				case "Удалить пациента":
					msg.Text = deletePatient(ChatId)
				case "Добавить пациента":
					msg.Text = addPatient(ChatId)
				case "Посмотреть своих пациентов":
					msg.Text = getInfoAboutPatients(ChatId)
				case "Найти палату по номеру":
					msg.Text = printByNum(ChatId)
				case "Вывести все палаты":
					msg.Text = printAllRooms(ChatId)
				case "open":
					msg.ReplyMarkup = numericKeyboard
				case "close":
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				default:
					msg.Text = "Команда не найдена, напишите Помощь"
				}
			}

			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}

func startBot(controller *controllers.Controller, cfg config.Config) {
	dotenv := cfg.TelegramToken

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
