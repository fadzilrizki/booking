package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "booking/model"
	"booking/service"

	"github.com/gorilla/mux"
)

func InsertBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var insertBook model.InsertBundleBook
	_ = json.NewDecoder(r.Body).Decode(&insertBook)
	fmt.Println(insertBook)

	service.InsertBook(insertBook.Genre, insertBook.Books)

	json.NewEncoder(w).Encode(insertBook)
}

func GetBooksGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	genre := params["genre"]

	books := service.GetBook(genre)

	json.NewEncoder(w).Encode(books)
}

func MakeAppointment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var appointment model.Appointment
	_ = json.NewDecoder(r.Body).Decode(&appointment)
	fmt.Println(appointment)

	service.MakeAppointment(appointment)

	json.NewEncoder(w).Encode(appointment)
}

func GetAppointment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	appointments := service.GetAppointment()

	json.NewEncoder(w).Encode(appointments)
}
