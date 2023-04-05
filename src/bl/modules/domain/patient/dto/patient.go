package dto

type Patient struct {
	id             int32
	surname        string
	name 		   string
	patronymic 	   string
	height 		   int32
	weight 		   float32
	roomNumber     int32
	degreeOfDanger int32
}

type Patients []*Patient

type CreatePatient struct {
	surname        string
	name 		   string
	patronymic 	   string
	height 		   int32
	weight 		   float32
	roomNumber     int32
	degreeOfDanger int32
}

type UpdatePatient {
	surname        string
	name 		   string
	patronymic 	   string
	height 		   int32
	weight 		   float32
	roomNumber     int32
	degreeOfDanger int32
}