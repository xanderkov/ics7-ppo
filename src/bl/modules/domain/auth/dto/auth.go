package dto

// Сервис для аунтефикации в телеграмм

type Auth struct {
	Id      int32
	TokenId int32
}

type Users []*Auth

type NewUser struct {
	TokenId    int32
	Surname    string
	Speciality string
	Role       string
}
