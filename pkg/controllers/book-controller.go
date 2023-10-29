package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/23piyush/go-bookstore/pkg/utils"
	"github.com/23piyush/go-bookstore/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks:=models.GetAllBooks()
	res, _ :=json.Marshal(newBooks) // convert whatever we received from database to json
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK) // 200
	w.Write(res) // our res is in json
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"] // bookId is in string
	ID, err:= strconv.ParseInt(bookId,0,0) // bookId to int
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _:= models.GetBookById(ID) // GetBookById() returns db as second variable
	// In golang, if we don't want to use a return variable, we use _
	res, _ := json.Marshal(bookDetails) // convert to json
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK) // 200
	w.Write(res) // res has json of details of book which we found by id
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{} // received json from user
	utils.ParseBody(r, CreateBook) // parsed from json to something golang understand
	b:= CreateBook.CreateBook()  // sent to database
	res, _ := json.Marshal(b) 
	w.WriteHeader(http.StatusOK) // converted back to json and sent back to user ( or postman)
	w.Write(res)
}

// r is a pointer to request received from user
func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book) // book details received from database needs to be converted to json before sending back to user (or postman)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook) // ParseBody does unmarshelling to convert json to a format which golang understands
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0) // convert string to int
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db:=models.GetBookById(ID)
	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails) // convert to json before sending back to user (or postman)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}