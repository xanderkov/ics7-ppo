package dto

type Doctor struct {
	Id         int32
	TokenId    int32
	Surname    string
	Speciality string
	Role       string
}

type Doctors []*Doctor

type CreateDoctor struct {
	Surname    string
	TokenId    int32
	Speciality string
	Role       string
}

type UpdateDoctor struct {
	Surname    *string
	Speciality *string
	Role       *string
}
