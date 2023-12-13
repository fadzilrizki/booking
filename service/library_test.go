package service_test

import (
	model "booking/model"
	service "booking/service"
	"reflect"
	"testing"
	"time"
)

func TestInsertBook(t *testing.T) {
	type args struct {
		genre string
		books []model.Book
	}
	tests := []struct {
		name string
		args args
		want map[string][]model.Book
	}{
		{
			name: "Insert books into an empty genre",
			args: args{
				genre: "fiction",
				books: []model.Book{{Title: "Book1"}, {Title: "Book2"}},
			},
			want: map[string][]model.Book{"fiction": {{Title: "Book1"}, {Title: "Book2"}}},
		},
		{
			name: "Insert more books into an existing genre",
			args: args{
				genre: "fiction",
				books: []model.Book{{Title: "Book3"}},
			},
			want: map[string][]model.Book{"fiction": {{Title: "Book1"}, {Title: "Book2"}, {Title: "Book3"}}},
		},
	}

	model.Shelf = make(map[string][]model.Book)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service.InsertBook(tt.args.genre, tt.args.books)

			if !reflect.DeepEqual(model.Shelf, tt.want) {
				t.Errorf("InsertBook failed. Expected %v, got %v", tt.want, model.Shelf)
			}
		})
	}
}

func TestGetBook(t *testing.T) {
	type args struct {
		genre string
	}
	tests := []struct {
		name string
		args args
		want []model.Book
	}{
		{
			name: "Get books from an existing genre",
			args: args{
				genre: "fiction",
			},
			want: []model.Book{{Title: "Book1"}, {Title: "Book2"}},
		},
		{
			name: "Get books from a non-existing genre",
			args: args{
				genre: "fantasy",
			},
			want: []model.Book{},
		},
	}

	model.Shelf = map[string][]model.Book{
		"fiction": {{Title: "Book1"}, {Title: "Book2"}},
		"history": {{Title: "Book3"}, {Title: "Book4"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.GetBook(tt.args.genre); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeAppointment(t *testing.T) {
	t1, _ := time.Parse("2023-12-22T10:00:00Z", "2023-12-22T10:00:00Z")
	t2, _ := time.Parse("2023-12-22T10:00:00Z", "2023-12-23T10:00:00Z")

	type args struct {
		appointment model.Appointment
	}
	tests := []struct {
		name string
		args args
		want []model.Appointment
	}{
		{
			name: "Make an appointment",
			args: args{
				appointment: model.Appointment{
					Borrower: "User 1",
					BookKey:  "123",
					Time:     &t1,
				},
			},
			want: []model.Appointment{{
				Borrower: "User 1",
				BookKey:  "123",
				Time:     &t1,
			}},
		},
		{
			name: "Make another appointment",
			args: args{
				appointment: model.Appointment{
					Borrower: "User 2",
					BookKey:  "456",
					Time:     &t2,
				},
			},
			want: []model.Appointment{{
				Borrower: "User 1",
				BookKey:  "123",
				Time:     &t1,
			}, {
				Borrower: "User 2",
				BookKey:  "456",
				Time:     &t2,
			}},
		},
	}

	model.Appointments = nil

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service.MakeAppointment(tt.args.appointment)

			appointments := service.GetAppointment()

			if !reflect.DeepEqual(appointments, tt.want) {
				t.Errorf("MakeAppointment failed. Expected %v, got %v", tt.want, model.Appointments)
			}
		})
	}
}
