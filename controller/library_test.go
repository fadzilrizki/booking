package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"booking/controller"
	model "booking/model"

	"github.com/gorilla/mux"
)

func TestInsertBook(t *testing.T) {
	requestBody := bytes.NewBufferString(`{"name": "fiction", "works": [{"key": "123", "title": "Book1"}]}`)
	request, err := http.NewRequest("POST", "/insert-book", requestBody)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()

	controller.InsertBook(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("InsertBook returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var result model.InsertBundleBook
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &result)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetBooksGenre(t *testing.T) {
	request, err := http.NewRequest("GET", "/get-books/fiction", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/get-books/{genre}", controller.GetBooksGenre).Methods("GET")

	router.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("GetBooksGenre returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestMakeAppointment(t *testing.T) {
	requestBody := bytes.NewBufferString(`{"borrower": "User1", "book_key": "123", "booking_time": null}`)
	request, err := http.NewRequest("POST", "/make-appointment", requestBody)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/make-appointment", controller.MakeAppointment).Methods("POST")

	router.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("MakeAppointment returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetAppointment(t *testing.T) {
	request, err := http.NewRequest("GET", "/get-appointment", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/get-appointment", controller.GetAppointment).Methods("GET")

	router.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("GetAppointment returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
