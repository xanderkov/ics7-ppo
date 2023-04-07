package dto

// Сервис для аунтефикации в телеграмм

type Auth struct {
	id             int32
	tokenId		   int32 
}

type Users []*Auth

type NewUser struct {
	tokenId        int32
	surname        string
	speciality     string
	role           string
}
