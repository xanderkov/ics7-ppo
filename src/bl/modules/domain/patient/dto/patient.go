package dto

type Patient struct {
	Id             int32
	Surname        string
	Name           string
	Patronymic     string
	Height         int32
	Weight         float32
	RoomNumber     int32
	DegreeOfDanger int32
}

type Patients []*Patient

type CreatePatient struct {
	Surname        string
	Name           string
	Patronymic     string
	Height         int32
	Weight         float32
	RoomNumber     int32
	DegreeOfDanger int32
}

type UpdatePatient struct {
	Surname        string
	Name           string
	Patronymic     string
	Height         int32
	Weight         float32
	RoomNumber     int32
	DegreeOfDanger int32
}
