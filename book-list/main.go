package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"database/sql"
	"github.com/alexanchek/go-rest-api/book-list/models"
	"os"

	"github.com/lib/pq"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"time"
	"log"
)


var books[] models.Book
var db *sql.DB


func init() {
	// load envinronment from env
	gotenv.Load()
}



func main() {
	currentTime := time.Now()
	fmt.Println("Server is running. Current time is -",currentTime)

	pgURL, err:= pq.ParseURL(os.Getenv("POSTRESSQL_URL"))
	log.Fatal(err)

	db, err = sql.Open("postgres", pgURL)
	if err!=nil {
		fmt.Println("An issue with opening")
		log.Fatal(err)
	}

	err = db.Ping()
	if err!=nil {
		fmt.Println("An issue with a ping with database")
		log.Println(pgURL)
	}
	router:=mux.NewRouter()

	// handlers to deal with requests
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	// listening a port
	log.Fatal(http.ListenAndServe(":8000", router))

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	books = []models.Book{}

	rows,err := db.Query("SELECT * FROM books")
	if err!=nil {
		fmt.Println("An issue with a connection to  database")
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err:= rows.Scan(&book.ID,&book.Title, &book.Author,  &book.Year)
		if err!=nil {
			fmt.Println("An issue with a request")
			log.Fatal(err)
		}

		books = append(books, book)
	}
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {

	var book models.Book
	params:=mux.Vars(r)

	row := db.QueryRow("SELECT * FROM books WHERE id=$1", params["id"])

	
	err := row.Scan(&book.ID, &book.Title, &book.Author,  &book.Year)
	if err!=nil {
		fmt.Println("There's no record")

	}
	
	json.NewEncoder(w).Encode(book)
	
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	var bookID int

	json.NewDecoder(r.Body).Decode(&book)
	err := db.QueryRow("insert into books (title, author, year) values($1,$2,$3) RETURNING id;",
	 book.Title, book.Author, book.Year).Scan(&bookID)
	 if err!=nil {
		fmt.Println("impossible to create")
	}

	json.NewEncoder(w).Encode(bookID)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	json.NewDecoder(r.Body).Decode(&book)
	result, err := db.Exec("UPDATE books SET title=$1, author=$2, year=$3 WHERE id=$4 RETURNING id;",
	 book.Title, book.Author, book.Year, book.ID)
	 if err!=nil {
		fmt.Println("impossible to update")
	 }

	 rowsUpdated, err := result.RowsAffected()
	 if err!=nil {
		fmt.Println("impossible to show off rows were updated")
	 }

	 json.NewEncoder(w).Encode(rowsUpdated)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	params:= mux.Vars(r)
	result, err :=db.Exec("DELETE from books where id=$1", params["id"])
	if err!=nil {
		fmt.Println("impossible to delete")
	 }

	 rowsDeleted, err := result.RowsAffected()
	 if err!=nil {
		fmt.Println("impossible to delete")
	 }

	json.NewEncoder(w).Encode(rowsDeleted)

}