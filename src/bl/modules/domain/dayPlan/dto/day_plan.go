package dto

type DayPlan struct {
	Id         int32
	IdDoc      int32
	PatientsId []int32
}

type DayPlans []*DayPlan

type CalculateDay struct {
	IdDoc      int32
	PatientsId int32
}

type ChangeDay struct {
	IdDoc      *int32
	PatientsId *int32
}
