package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	controller "booking/controller"
)

/**
* Main function to handle all of the routes...
 */
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{genre}", controller.GetBooksGenre).Methods("GET")
	r.HandleFunc("/books", controller.InsertBook).Methods("POST")

	r.HandleFunc("/appointments", controller.GetAppointment).Methods("GET")
	r.HandleFunc("/appointments", controller.MakeAppointment).Methods("POST")

	log.Println("Server listening on :8090")
	log.Fatal(http.ListenAndServe(":8090", r))
}
