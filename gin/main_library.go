package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Author struct {
	Author_id     int    `json:"author_id"`
	Name          string `json:"name"`
	Surname       string `json:"surname"`
	Patronymic    string `json:"patronymic"`
	Date_of_birth string `json:"date_of_birth"`
	Email         string `json:"email"`
	Country       string `json:"country"`
}

type Genre struct {
	Genre_id    int    `json:"genre_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Book struct {
	Book_id                int    `json:"book_id"`
	Title                  string `json:"title"`
	Author_id              int    `json:"author_id"`
	Genre_id               int    `json:"genre_id"`
	Pages_amount           int    `json:"pages_amount"`
	Popularity_coefficient int    `json:"popularity_coefficient"`
}

type User struct {
	User_id    int    `json:"user_id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age"`
	Email      string `json:"email"`
	Country    string `json:"country"`
	Address    string `json:"address"`
}

type Library struct {
	Library_id int    `json:"library_id"`
	Country    string `json:"country"`
	City       string `json:"city"`
	Address    string `json:"address"`
	Capacity   int    `json:"capacity"`
	Email      string `json:"email"`
}

type Loan struct {
	Loan_id     int    `json:"loan_id"`
	Library_id  int    `json:"library_id"`
	Book_id     int    `json:"book_id"`
	User_id     int    `json:"user_id"`
	Loan_date   string `json:"loan_date"`
	Return_date string `json:"return_date"`
}

var authors []Author
var genres []Genre
var users []User
var books []Book
var libraries []Library
var loans []Loan

func main() {
	// Инициализируем тестовыми данными
	libraries = []Library{{Library_id: 1, Country: "Russia", City: "Moscow", Address: "12121212", Capacity: 100, Email: "moscow_lib@gmail.com"}}
	loans = []Loan{{Loan_id: 1, Library_id: 1, Book_id: 2, User_id: 1, Loan_date: "01-01-24", Return_date: "04-01-24"}}
	books = []Book{{Book_id: 1, Title: "На западном фронте без перемен", Author_id: 1, Genre_id: 1, Pages_amount: 200, Popularity_coefficient: 52}}
	users = []User{{User_id: 1, Name: "Alex", Surname: "Soldatov", Patronymic: "Konstantinovich", Age: 18, Email: "alalalalal", Country: "Russia", Address: "Moscow"}}
	genres = []Genre{{Genre_id: 1, Name: "Novel", Description: "123"}}
	authors = []Author{{Author_id: 1, Name: "Erich Maria", Surname: "Remark", Patronymic: "", Date_of_birth: "22-06-1898", Email: "", Country: "German Empire"}}

	http.HandleFunc("/authors", autHandler)
	http.HandleFunc("/authors/", autByIdHandler)

	http.HandleFunc("/genres", genHandler)
	http.HandleFunc("/genres/", genByIdHandler)

	http.HandleFunc("/users", userHandler)
	http.HandleFunc("/users/", userByIdHandler)

	http.HandleFunc("/books", bookHandler)
	http.HandleFunc("/books/", bookByIdHandler)

	http.HandleFunc("/libraries", libHandler)
	http.HandleFunc("/libraries/", libByIdHandler)

	http.HandleFunc("/loans", loanHandler)
	http.HandleFunc("/loans/", loanByIdHandler)

	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func autHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(authors)
	case "POST":
		var newAuthor Author
		err := json.NewDecoder(r.Body).Decode(&newAuthor)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if newAuthor.Author_id < 0 {
			http.Error(w, "ID field cannot be < 0", http.StatusBadRequest)
			return
		}
		authors = append(authors, newAuthor)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newAuthor)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func genHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(genres)
	case "POST":
		var newGenre Genre
		err := json.NewDecoder(r.Body).Decode(&newGenre)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if newGenre.Genre_id < 0 {
			http.Error(w, "ID field cannot be < 0", http.StatusBadRequest)
			return
		}
		genres = append(genres, newGenre)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newGenre)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	case "POST":
		var newUser User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if newUser.User_id < 0 {
			http.Error(w, "ID field cannot be < 0", http.StatusBadRequest)
			return
		}
		users = append(users, newUser)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
	case "POST":
		var newBook Book
		err := json.NewDecoder(r.Body).Decode(&newBook)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if newBook.Book_id < 0 || newBook.Author_id < 0 || newBook.Genre_id < 0 {
			http.Error(w, "ID field cannot be < 0", http.StatusBadRequest)
			return
		}
		books = append(books, newBook)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newBook)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func libHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(libraries)
	case "POST":
		var newLibrary Library
		err := json.NewDecoder(r.Body).Decode(&newLibrary)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if newLibrary.Library_id < 0 {
			http.Error(w, "ID field cannot be < 0", http.StatusBadRequest)
			return
		}
		libraries = append(libraries, newLibrary)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newLibrary)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func loanHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(loans)
	case "POST":
		var newLoan Loan
		err := json.NewDecoder(r.Body).Decode(&newLoan)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if newLoan.Library_id < 0 || newLoan.Book_id < 0 || newLoan.Loan_id < 0 || newLoan.User_id < 0 {
			http.Error(w, "ID field cannot be < 0", http.StatusBadRequest)
			return
		}
		loans = append(loans, newLoan)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newLoan)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func autByIdHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Author not found", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(authors[id])
	case "DELETE":
		f := true
		for i := 0; i < len(authors); i++ {
			fmt.Println(authors[i].Author_id, id)
			if authors[i].Author_id == id {
				f = false
				authors = append(authors[:i], authors[i+1:]...)
			}
		}
		if f {
			http.Error(w, "Author not found", http.StatusNotFound)
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func genByIdHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Genre not found", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		f := true
		for i := 0; i < len(loans); i++ {
			fmt.Println(genres[i].Genre_id, id)
			if genres[i].Genre_id == id {
				f = false
				json.NewEncoder(w).Encode(genres[i])
			}
		}
		if f {
			http.Error(w, "Genre not found", http.StatusNotFound)
		}
		w.WriteHeader(http.StatusNoContent)

	case "DELETE":
		f := true
		for i := 0; i < len(loans); i++ {
			fmt.Println(genres[i].Genre_id, id)
			if genres[i].Genre_id == id {
				f = false
				genres = append(genres[:i], genres[i+1:]...)
			}
		}
		if f {
			http.Error(w, "Genre not found", http.StatusNotFound)
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func userByIdHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users[id])
	case "DELETE":
		f := true
		for i := 0; i < len(users); i++ {
			fmt.Println(users[i].User_id, id)
			if users[i].User_id == id {
				f = false
				users = append(users[:i], users[i+1:]...)
			}
		}
		if f {
			http.Error(w, "User not found", http.StatusNotFound)
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func bookByIdHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books[id])
	case "DELETE":
		f := true
		for i := 0; i < len(books); i++ {
			fmt.Println(books[i].Book_id, id)
			if books[i].Book_id == id {
				f = false
				books = append(books[:i], books[i+1:]...)
			}
		}
		if f {
			http.Error(w, "Book not found", http.StatusNotFound)
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func libByIdHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Library not found", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(libraries[id])
	case "DELETE":
		f := true
		for i := 0; i < len(libraries); i++ {
			fmt.Println(libraries[i].Library_id)
			if libraries[i].Library_id == id {
				f = false
				libraries = append(libraries[:i], libraries[i+1:]...)
			}
		}
		if f {
			http.Error(w, "Library not found", http.StatusNotFound)
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func loanByIdHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		http.Error(w, "Loan not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(loans[id])
	case "DELETE":
		f := true
		for i := 0; i < len(loans); i++ {
			fmt.Println(loans[i].Loan_id, id)
			if loans[i].Loan_id == id {
				f = false
				loans = append(loans[:i], loans[i+1:]...)
			}
		}
		if f {
			http.Error(w, "Loan not found", http.StatusNotFound)
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
