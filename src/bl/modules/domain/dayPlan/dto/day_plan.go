package dto

type DayPlan struct {
	id             int32
	idDoc          int32
	patientsId	   []int32
}

type DayPlans []*DayPlan

type CalculateDay struct {
	idDoc	       int32
	patientsId 	   int32
}

type ChangeDay {
	idDoc        *int32
	patientsId	 *int32
}