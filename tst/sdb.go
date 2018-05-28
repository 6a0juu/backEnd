package main

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func readForm(qStr []byte) Studss {
	var studf Studss
	err := json.Unmarshal(qStr, &studf)
	if err != nil {
		log.Println(err)
	}
	return studf
}

func addStud(qStr []byte) int {
	mod := Stud{}
	err := json.Unmarshal(qStr, mod)
	if err != nil {
		log.Println(err)
		return 1
	}
	stmt, err := db.Prepare("INSERT stud_table SET sid=?, name=?, tel=?, email=?")
	if err != nil {
		log.Println(err)
		return 2
	}
	res, err := stmt.Exec(mod.SID, mod.Name, mod.Email, mod.Tel, mod.SID)
	log.Println(res)
	if err != nil {
		log.Println(err)
		return 3
	}

	return 0
}

func delStud(qStr []byte) (int, []byte) {
	mod := Stud{}
	err := json.Unmarshal(qStr, mod)
	stmt, err := db.Prepare("DELETE FROM stud_table WHERE sid=?, name=?, tel=?, email=?")
	log.Println(stmt, err)
	res, err := stmt.Exec(mod.SID, mod.Name, mod.Email, mod.Tel, mod.SID)
	log.Println(res, err)

	return 0, qStr
}

func edtStud(qStr []byte) (int, []byte) {
	mod := Stud{}
	err := json.Unmarshal(qStr, mod)
	stmt, err := db.Prepare("UPDATE stud_table SET sid=?, name=?, tel=?, email=? WHERE sid = ?")
	log.Println(stmt, err)
	res, err := stmt.Exec(mod.SID, mod.Name, mod.Email, mod.Tel, mod.SID)
	log.Println(res, err)

	return 0, qStr
}

func serStud(qStr []byte) (int, []byte) {
	// ret: 0 for all paired, 1 for username paired, 2 for nothing paired
	mod := Stud{}
	nMod := Stud{}
	err := json.Unmarshal(qStr, mod)
	err = db.QueryRow("SELECT * FROM stud_table WHERE sid LIKE ?", mod.SID).Scan(&nMod.SID, &nMod.Name, &nMod.Tel, &nMod.Email)
	log.Println(err)
	retData, err := json.Marshal(nMod)
	if err != nil {
		return 2, retData
	} else if nMod.SID != "" && (nMod.SID == mod.SID) {
		return 0, retData
	}
	return 1, retData
}

func pAddStud(qStr []byte) (int, []byte) {
	return 1, qStr
}

func pDelStud(qStr []byte) (int, []byte) {
	return 1, qStr
}

func pSerStud(qStr []byte) (int, []byte) {

	mod := Stud{}
	err := json.Unmarshal(qStr, mod)
	rows, err := db.Query("SELECT * FROM stud_table WHERE sid LIKE ?, name LIKE ?, cast(tel as varchar(20)) LIKE ?, email LIKE ?", "%"+mod.SID+"%", "%"+mod.Name+"%", "%"+mod.Email+"%", "%"+string(mod.Tel)+"%", "%"+mod.SID+"%") //tel for char
	var xxxx []byte
	if err != nil {
		return 1, xxxx
	}
	defer rows.Close()
	var Studs []Stud
	it := 0
	for rows.Next() {
		rows.Scan(&mod.SID, &mod.Name, &mod.Tel, &mod.Email)
		Studs[it] = mod
	}
	retData, err := json.Marshal(Studs)
	if err != nil {
		return 1, retData
	}
	return 0, retData

	return 1, qStr
}

func retAll() (int, []byte) {
	mod := Stud{}
	rows, err := db.Query("SELECT * FROM stud_table WHERE sid=?, name=?, tel=?, email=?", "%", "%", "%", "%")
	var xxxx []byte
	if err != nil {
		return 1, xxxx
	}
	defer rows.Close()
	var Studs []Stud
	it := 0
	for rows.Next() {
		rows.Scan(&mod.SID, &mod.Name, &mod.Tel, &mod.Email)
		Studs[it] = mod
	}
	retData, err := json.Marshal(Studs)
	if err != nil {
		return 1, retData
	}
	return 0, retData
}

func sdb(op string, qStr []byte) (int, []byte) {
	studb, err := sql.Open("mysql", "root:bjwdttz@tcp(127.0.0.1:3306)/tst?charset=utf8")
	retCode := 0
	var retData []byte
	if err != nil {
		return 10, retData
	}
	switch op {
	case "add":
		retCode, retData = addStud(qStr)
	case "del":
		retCode, retData = delStud(qStr)
	case "edt":
		retCode, retData = edtStud(qStr)
	case "ser":
		retCode, retData = serStud(qStr)
	case "pAdd":
		retCode, retData = pAddStud(qStr)
	case "pDel":
		retCode, retData = pDelStud(qStr)
	case "pEdt":
		retCode, retData = pEdtStud(qStr)
	case "pSer":
		retCode, retData = pSerStud(qStr)
	}
	log.Println(retCode, retData)
	defer studb.Close()
	return retCode, retData
}
