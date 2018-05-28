package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func addUser(db *sql.DB, newUser user) {
	stmt, err := db.Prepare("INSERT userinfo SET username=?, password=?")
	fmt.Println(stmt)
	fmt.Println(err)
	res, err := stmt.Exec(newUser.username, "qwq", "qwq")
	fmt.Println(res)
	fmt.Println(err)
}

func searchUser(db *sql.DB, newUser user) int { //search needs all thing
	rows, err := db.Query("SELECT * FROM user_table WHERE username = '" + newUser.username + "' AND password = '" + newUser.password + "'")
	for rows.Next() {
		var username string
		var password string
		var permission string
		err = rows.Scan(&username, &password, &permission)

		fmt.Println(newUser.username)
		fmt.Println(newUser.password)
		fmt.Println(permission)
		if err != nil {
			log.Fatal(err)
			fmt.Println(2)

			return 2
			//return 1, "Database Connection Failed."
		} else if newUser.username == username && newUser.password == password {
			fmt.Println(0)
			return 0
		} else {
			fmt.Println(1)

			return 1
		}

	}
	return 3
}

func updateUser(db *sql.DB, newUser user) {
	stmt, err := db.Prepare("UPDATE userinfo SET username=?,departname=?,created=?")
	fmt.Println(stmt)
	fmt.Println(err)
	res, err := stmt.Exec("qwq", "qwq", "qwq")
	fmt.Println(res)
	fmt.Println(err)
}

func deleteUser(db *sql.DB, newUser user) {
	//delete where .. and .. and ..
}

func mdb() (int, string) {
	retCode := 0
	retData := "n"
	newUser := user{"1", "1", "1"}

	opCode := 2
	db, err := sql.Open("mysql", "root:dttz1998@tcp(127.0.0.1:3306)/tst?charset=utf8")
	if err != nil {
		log.Fatal(err)
		return 1, "Database Connection Failed."
	}
	switch opCode { //new user for test, must be changed
	case 1:
		addUser(db, newUser)
	case 2:
		retCode = searchUser(db, newUser)
	case 3:
		updateUser(db, newUser)
	case 4:
		deleteUser(db, newUser)
	}
	/*

	 */
	defer db.Close()
	return retCode, retData
}
