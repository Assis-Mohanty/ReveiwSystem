package dto

type GetReveiwDTO struct {
	userId  int64
	hotelId   int64
	bookingId int64
	comment string
	rating   float64
}