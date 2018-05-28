package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

var db *sql.DB
var err error

func main() {

	db, err = sql.Open("mysql", "root:bjwdttz@tcp(127.0.0.1:3306)/tst?charset=utf8")
	if err != nil {
		log.Fatal(err)
		return
	}
	router := httprouter.New()
	router.GET("/", Index)

	router.POST("/api/login", SignIn)
	router.OPTIONS("/api/login", SignIn)
	router.POST("/api/user", SignUp)
	router.PUT("/api/user", UserUpdate)
	router.DELETE("/api/user", UserDelete)
	router.POST("/api/item", ItemAdd)
	router.DELETE("/api/item", ItemDelete)
	router.PUT("/api/item", ItemUpdate)
	/*
		router.POST("/api/item", ItemAdd)
		router.DELETE("/api/item", ItemDelete)
		router.PUT("/api/item", ItemUpdate)
	*/
	router.GET("/api/form", MultiSearch)

	log.Fatal(http.ListenAndServe(":19845", router))
}

// Create a couple of sample Book entries
/*
	bookstore["123"] = &Book{
		ISDN:   "123",
		Title:  "Silence of the Lambs",
		Author: "Thomas Harris",
		Pages:  367,
	}
	bookstore["124"] = &Book{
		ISDN:   "124",
		Title:  "To Kill a Mocking Bird",
		Author: "Harper Lee",
		Pages:  320,
	}


	router.GET("/books", BookIndex)
	router.GET("/books/:isdn", BookShow)
	router.GET("/books/:isdn")
*/
