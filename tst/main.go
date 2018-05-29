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
	router.OPTIONS("/api/user", SignUp)
	router.PUT("/api/user", UserUpdate)
	router.DELETE("/api/user", UserDelete)
	router.POST("/api/item", ItemAdd)
	router.OPTIONS("/api/item", ItemAdd)
	router.OPTIONS("/api/itemdel", ItemDelete)
	router.PUT("/api/itemdel", ItemDelete)
	router.PUT("/api/item", ItemUpdate)
	/*
		router.POST("/api/form", ItemAdd)
		router.DELETE("/api/form", ItemDelete)
		router.PUT("/api/form", ItemUpdate)
	*/
	router.GET("/api/all", RetAll)
	router.OPTIONS("/api/all", RetAll)
	router.GET("/api/form", MultiSearch)

	log.Fatal(http.ListenAndServe(":19845", router))
}
