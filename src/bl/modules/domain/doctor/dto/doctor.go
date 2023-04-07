package dto

type Doctor struct {
	id             int32
	tokenId        int32
	surname        string
	speciality	   string
	role 		   string
}

type Doctors []*Doctor

type CreateDoctor struct {
	surname        string
	tokenId        int32
	speciality	   string
	role 		   string
}

type UpdateDoctor {
	surname        *string
	speciality	   *string
	role 		   *string
}