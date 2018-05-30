package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Users []User

func addUser(newUser User) int {
	stmt, err := db.Prepare("INSERT user_table SET username=?, password=?, permission=1")
	fmt.Println(stmt, err)

	res, err := stmt.Exec(newUser.Usnm, newUser.Pswd)
	fmt.Println(res)
	fmt.Println(err)
	return 0
}

func searchUser(newUser User) int {
	// ret: 0 for all paired, 1 for username paired, 2 for nothing paired
	var username, password string
	err := db.QueryRow("SELECT username, password FROM user_table WHERE username=?", newUser.Usnm).Scan(&username, &password)
	fmt.Println(err)
	if err != nil {
		return 2
	}
	if newUser.Usnm == username && newUser.Pswd == password {
		return 0
	} else if newUser.Usnm == username {
		return 1
	}
	return 2
}

func updateUser(newUser User) int {
	stmt, err := db.Prepare("UPDATE user_table SET password=? WHERE username=?")
	fmt.Println(err)
	res, err := stmt.Exec(newUser.Pswd, newUser.Usnm)
	fmt.Println(res, err)
	return 0
}

func deleteUser(newUser User) int {
	//delete where .. and .. and ..
	stmt, err := db.Prepare(`DELETE FROM user_table WHERE username=?`)
	fmt.Println(err)
	res, err := stmt.Exec(newUser.Usnm)
	fmt.Println(res, err)
	return 0
}

func allUser(newUser User) int {
	var tmpusers []User
	Users = tmpusers
	rows, err := db.Query("SELECT username, password FROM user_table")
	if err != nil {
		return 1
	}
	defer rows.Close()
	for rows.Next() {
		mod := User{}
		rows.Scan(&mod.Usnm, &mod.Pswd)
		Users = append(Users, mod)
	}
	if err != nil {
		return 1
	}
	return 0
}

func mdb(opCode int, newUser *User) int {
	retCode := 0
	switch opCode { //new user for test, must be changed
	case 1:
		retCode = addUser(*newUser)
	case 2:
		retCode = searchUser(*newUser)
	case 3:
		retCode = updateUser(*newUser)
	case 4:
		retCode = deleteUser(*newUser)
	case 5:
		retCode = allUser(*newUser)
	}
	/*

	 */
	return retCode
}
