package service

import (
	model "booking/model"
)

func InsertBook(genre string, books []model.Book) {
	if model.Shelf == nil {
		model.Shelf = make(map[string][]model.Book)
	}

	val, ok := model.Shelf[genre]
	if ok {
		val = append(val, books...)
		model.Shelf[genre] = val
	} else {
		model.Shelf[genre] = books
	}
}

func GetBook(genre string) []model.Book {

	val, ok := model.Shelf[genre]
	if ok {
		return val
	}

	return []model.Book{}
}

func MakeAppointment(appointment model.Appointment) {
	model.Appointments = append(model.Appointments, appointment)
}

func GetAppointment() []model.Appointment {
	return model.Appointments
}
