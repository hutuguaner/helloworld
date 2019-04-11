package main

import(
	"fmt"
	"database/sql"
)

func Query()  {
	db,err := sql.Open("mysql","root:root@/daqiuma?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}

	defer db.close()

	var who string
	var place string
	var ps string
	var holddays int

	rows,err := db.Query("select * from pingpang")
	if err!=nil{
		fmt.Println(err)
	}

	for rows.Next(){
		rows.Scan(&who,&place,&ps,&holddays)
		fmt.Println(who,place,ps,holddays)
	}

	defer rows.close()
}