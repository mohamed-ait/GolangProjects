package controllers

import (
	"GolangProjects/BOOKSTORE/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"GolangProjects/BOOKSTORE/pkg/utils"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	BookId := vars["bookId"]
	ID, err := strconv.ParseInt(BookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing !")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(r, newBook) //parse to database format
	b := newBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func main() {
}
