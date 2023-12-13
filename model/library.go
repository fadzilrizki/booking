package model

import "time"

type Appointment struct {
	Borrower string     `json:"borrower"`
	BookKey  string     `json:"book_key"`
	Time     *time.Time `json:"booking_time"`
}

var Appointments []Appointment

var Shelf map[string][]Book
