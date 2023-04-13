package dto

type Room struct {
	Num            int32
	Floor          int32
	NumberBeds     int32
	TypeRoom       string
	NumberPatients int32
}

type Rooms []*Room

type CreateRoom struct {
	Floor          int32
	NumberBeds     int32
	TypeRoom       string
	NumberPatients int32
}

type UpdateRoom struct {
	Floor          *int32
	NumberBeds     *int32
	TypeRoom       *string
	NumberPatients int32
}
