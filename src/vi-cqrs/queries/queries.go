package queries

type GetConcertByIDQuery struct {
	ID int
}

type GetAvailableConcertsQuery struct {
	MinAvailableSeats int
}

type GetStudentTicketsQuery struct {
	StudentName string
}
