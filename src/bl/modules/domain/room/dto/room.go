package dto

type Room struct {
	num            int32
	floor          int32
	numberBeds     int32
	typeRoom       string
	numberPatients int32
}

type Room []*Room
