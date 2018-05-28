package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tstToken string = "123456"

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to !\n")
}

type Stud struct {
	// The main identifier for the Book. This will be unique.
	SID    string `json:"sid"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Tel    int    `json:"tel"`
	OriSID string `json:"orisid"`
}

type Studss struct {
	// May not used
	Studs []Stud `json:"studs"`
}

type StudIn struct {
	// The main identifier for the Book. This will be unique.
	SID   string `json:"sid"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Tel   int    `json:"tel"`
}

type User struct {
	// The main identifier for the Book. This will be unique.
	Usnm string `json:"usnm"`
	Pswd string `json:"pswd"`
}

type Idtf struct {
	Token string `json:"token"`
}

type JsonRes struct {
	// Reserved field to add some meta information to the API response
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}

type JsonErrRes struct {
	Error *ApiErr `json:"error"`
}

type ApiErr struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
}

func SignIn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	}
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	tmpUser := &User{}
	if err := populateModelFromHandler(w, r, params, tmpUser); err != nil {
		writeErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		return
	}
	fmt.Println(tmpUser.Usnm, tmpUser.Pswd)
	retCode, retData := mdb(2, tmpUser)
	fmt.Println(retCode, retData)
	if retCode != 0 {
		writeErrorResponse(w, http.StatusNotAcceptable, "Not Acceptable")
		return
	}
	if err := json.NewEncoder(w).Encode(&JsonRes{Meta: &Idtf{tstToken}}); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}

func SignUp(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	tmpUser := &User{}
	if err := populateModelFromHandler(w, r, params, tmpUser); err != nil {
		writeErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		return
	}
	retCode, retData := mdb(2, tmpUser)
	fmt.Println(retCode, retData)
	if retCode != 2 {
		writeErrorResponse(w, http.StatusConflict, "Conflict")
		return
	}
	retCode, retData = mdb(1, tmpUser)
	fmt.Println(retCode, retData)
	if retCode != 0 {
		writeErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Add("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UserUpdate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Update user, pending
	tmpUser := &User{}
	if err := populateModelFromHandler(w, r, params, tmpUser); err != nil {
		writeErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		return
	}
	retCode, retData := mdb(2, tmpUser)
	fmt.Println(retCode, retData)
	if retCode != 0 {
		writeErrorResponse(w, http.StatusNotAcceptable, "Not Acceptable")
		return
	}
	retCode, retData = mdb(3, tmpUser)
	fmt.Println(retCode, retData)
	if retCode != 0 {
		writeErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	//if searchUser(tmpUser.Usnm) {OK} else {gg}
	//if updateUser(tmpUser.Usnm, tmpUser.Pswd) {OK} else {gg}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Add("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UserDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	tmpUser := &User{}
	if err := populateModelFromHandler(w, r, params, tmpUser); err != nil {
		writeErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		return
	}

	retCode, retData := mdb(2, tmpUser)
	fmt.Println(retCode, retData)
	if retCode != 0 {
		writeErrorResponse(w, http.StatusNotAcceptable, "Not Acceptable")
		return
	}
	retCode, retData = mdb(4, tmpUser)
	fmt.Println(retCode, retData)
	if retCode != 0 {
		writeErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	//if searchUser(tmpUser.Usnm, tmpUser.Pswd) {OK} else {gg}
	//if deleteUser(tmpUser.Usnm, tmpUser.Pswd) {OK} else {gg}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Add("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func ItemAdd(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("afwafw")
}

func ItemDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("afwafw")
}

func ItemUpdate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("afwafw")
}

func MultiSearch(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("afwafw")
}

// Writes the response as a standard JSON response with StatusOK
func writeOKResponse(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Add("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&JsonRes{Data: m}); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
	}
}

// Writes the error response as a Standard API JSON response with a response code
func writeErrorResponse(w http.ResponseWriter, errorCode int, errorMsg string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Add("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(errorCode)
	json.NewEncoder(w).
		Encode(&JsonErrRes{Error: &ApiErr{Status: errorCode, Title: errorMsg}})
}

//Populates a model from the params in the Handler
func populateModelFromHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params, model interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return err
	}
	if err := r.Body.Close(); err != nil {
		return err
	}
	if err := json.Unmarshal(body, model); err != nil {
		return err
	}
	return nil
}

func populateStrFromHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params, body []byte) error {
	var err error
	body, err = ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return err
	}
	if err := r.Body.Close(); err != nil {
		return err
	}
	return nil
}

/*
type Book struct {
	// The main identifier for the Book. This will be unique.
	ISDN   string `json:"isdn"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}
type JsonResponse struct {
	// Reserved field to add some meta information to the API response
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}
type JsonErrorResponse struct {
	Error *ApiError `json:"error"`
}
type ApiError struct {
	Status int16  `json:"status"`
	Title  string `json:"title"`
}

// A map to store the books with the ISDN as the key
// This acts as the storage in lieu of an actual database
var bookstore = make(map[string]*Book)
// Handler for the books index action
// GET /books
func BookIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	books := []*Book{}
	for _, book := range bookstore {
		books = append(books, book)
	}
	response := &JsonResponse{Data: &books}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

// Handler for the books Show action
// GET /books/:isdn
func BookShow(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	isdn := params.ByName("isdn")
	fmt.Println(isdn)
	book, ok := bookstore[isdn]
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if !ok {
		// No book with the isdn in the url has been found
		w.WriteHeader(http.StatusNotFound)
		response := JsonErrorResponse{Error: &ApiError{Status: 404, Title: "Record Not Found"}}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
	}
	response := JsonResponse{Data: book}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
*/
