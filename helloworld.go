package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DbWorker struct {
	Dsn      string
	Db       *sql.DB
	UserInfo userTB
}
type userTB struct {
	age  sql.NullInt64
	name sql.NullString
	sex  sql.NullInt64
}

func hehe(){
	fmt.Println("other")
}

func main() {
	var err error
	dbw := DbWorker{
		Dsn: "root:root@tcp(localhost:3306)/daqiuma?charset=utf8mb4",
	}
	dbw.Db, err = sql.Open("mysql", dbw.Dsn)
	if err != nil {
		panic(err)
		return
	}
	defer dbw.Db.Close()

	dbw.queryData()
}


func (dbw *DbWorker) QueryDataPre() {
	dbw.UserInfo = userTB{}
}
func (dbw *DbWorker) queryData() {
	stmt, _ := dbw.Db.Prepare(`SELECT * From user`)
	defer stmt.Close()

	dbw.QueryDataPre()

	rows, err := stmt.Query()
	defer rows.Close()
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	for rows.Next() {
		rows.Scan(&dbw.UserInfo.age, &dbw.UserInfo.name, &dbw.UserInfo.sex)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}
		if !dbw.UserInfo.name.Valid {
			dbw.UserInfo.name.String = ""
		}
		if !dbw.UserInfo.age.Valid {
			dbw.UserInfo.age.Int64 = 0
		}
		if !dbw.UserInfo.sex.Valid{
			dbw.UserInfo.sex.Int64 = 0
		}
		fmt.Println("get data, id: ", dbw.UserInfo.age.Int64, " name: ", dbw.UserInfo.name.String, " age: ", int(dbw.UserInfo.sex.Int64))
	}

	err = rows.Err()
	if err != nil {
		fmt.Printf(err.Error())
	}
}