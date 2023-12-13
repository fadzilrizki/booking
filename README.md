# booking

This program is for technical test of cosmart.

## Running Program

```bash
go install

go run . 
# it will serve at port 8090
```

## APIs

POST insert books
using data from http://openlibrary.org/subjects/love.json

```bash
POST http://localhost:8090/books
```

GET books

```bash
GET http://localhost:8090/books/{genre}
```

POST appointments

```bash
POST http://localhost:8090/appointments

{
    "borrower": "Fadzil",
    "book_key": "/works/OL362427WX",
    "booking_time": "2023-12-22T10:00:00Z"
}
```

GET appointments

```bash
GET http://localhost:8090/appointments
```

## License

Fadzil Rizki