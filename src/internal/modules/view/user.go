package view

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Queue struct {
	front  int
	rear   int
	size   int
	QArray []string
}

func (q *Queue) initQueue(size int) {
	q.size = size
	q.QArray = make([]string, q.size)
	q.front = -1
	q.rear = -1
}

func (q *Queue) enqueue(value string) {
	if q.rear == q.size-1 {
		fmt.Println("Queue is Full")
		return
	} else {
		q.rear++
		q.QArray[q.rear] = value
	}
}

func (q *Queue) dequeue() string {
	var x string
	if q.front == q.rear {
		fmt.Println("Queue is Empty!")
	} else {
		q.front++
		x = q.QArray[q.front]
	}
	return x
}

type UsersMessage struct {
	NextMessages Queue
	ChatId       int64
	UserMessages []string
	Command      string
}

func addNextMessages(message string, user *UsersMessage, chatId int64) {
	msg := tgbotapi.NewMessage(chatId, message)
	user.NextMessages = append(user.NextMessages, msg.Text)
}
