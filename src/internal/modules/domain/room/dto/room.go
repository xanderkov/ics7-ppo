package dto

type Room struct {
	Num            int
	Floor          int
	NumberBeds     int
	TypeRoom       string
	NumberPatients int
}

type Rooms []*Room

type CreateRoom struct {
	Floor          int
	NumberBeds     int
	TypeRoom       string
	NumberPatients int
}

type UpdateRoom struct {
	Floor          *int
	NumberBeds     *int
	TypeRoom       *string
	NumberPatients int
}
