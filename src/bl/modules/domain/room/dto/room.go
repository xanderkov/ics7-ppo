package dto

type Room struct {
	num            int32
	floor          int32
	numberBeds     int32
	typeRoom       string
	numberPatients int32
}

type Rooms []*Room

type CreateRoom struct {
	floor          int32
	numberBeds     int32
	typeRoom       string
	numberPatients int32
}

type UpdateRoom {
	floor          &int32
	numberBeds     &int32
	typeRoom       &string
	numberPatients int32
}