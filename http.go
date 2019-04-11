package main

import (
    //"encoding/json"
    "fmt"
    //"io/ioutil"
    "log"
    "net/http"
	//"strings"
	"strconv"
	"helloworld/hehe"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	
)




func pingpangshare(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var who string
	var place string
	var ps string
	var holddays int

	who = r.Form["who"][0]
	place = r.Form["place"][0]
	ps = r.Form["ps"][0]
	holddays, _ = strconv.Atoi(r.Form["holddays"][0])

	fmt.Println(who)
	fmt.Println(place)
	fmt.Println(ps)
	fmt.Println(holddays)

	insert(who,place,ps,holddays)

    fmt.Println(" it works ...")
}




func insert(who string,place string,ps string,holddays int){

	//"root:root@tcp(localhost:3306)/daqiuma?charset=utf8mb4"

	db,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/hehe")
	defer db.Close()
	check(err)

	//ret,err :=db.Exec("INSERT INTO pingpang (who,place,holddays,ps) VALUES ('who','place',2,'ps')")
	stmt,err := db.Prepare("INSERT pingpang (who,place,holddays,ps) VALUES (?,?,?,?)")
	check(err)

	res,err := stmt.Exec(who,place,holddays,ps)

	id,err := res.LastInsertId()
	check(err)

	fmt.Println(id)
	stmt.Close()

}

func getpingpang(w http.ResponseWriter,r *http.Request){
	db,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/hehe")
	defer db.Close()
	check(err)

	//ret,err :=db.Exec("INSERT INTO pingpang (who,place,holddays,ps) VALUES ('who','place',2,'ps')")
	stmt,err := db.Prepare("select * from pingpang")
	check(err)

	defer stmt.Close()

	rows, err := stmt.Query()
	defer rows.Close()

	if err != nil {
		fmt.Printf("query data error: %v\n", err)
		return
	}
	for rows.Next() {

		var pingpang pingPang

		rows.Scan(&pingpang.who, &pingpang.place, &pingpang.ps)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}
		if !pingpang.who.Valid {
			pingpang.who.String = ""
		}
		if !pingpang.place.Valid {
			pingpang.place.String = ""
		}
		if !pingpang.ps.Valid {
			pingpang.who.String = ""
		}
		
		
		fmt.Println("get data, who: ", pingpang.who, " place: ", pingpang.place, " ps: ", pingpang.ps)
	}

	err = rows.Err()
	if err != nil {
		fmt.Printf(err.Error())
	}
}

type pingPang struct {
	who sql.NullString
	place sql.NullString
	ps sql.NullString
	
}



func check(err error){
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
}



func main() {

	hehe.Hehe()
    
	http.HandleFunc("/pingpangshare/", pingpangshare)
	http.HandleFunc("/getpingpang/",hehe.Query)
	

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}



