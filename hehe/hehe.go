package hehe
 

import(
	"net/http"
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"encoding/json"
)

func Hehe(){

	fmt.Println("hh")
}

type ResponseJson struct{
	Code int `json:"code"`
	Data []Pingpang `json:"data"`
}

type Pingpang struct{
	Who string `json:"who"`
	Place string `json:"place"`
	Holddays int `json:"holddays"`
	Ps string `json:"ps"`
}


func Query(w http.ResponseWriter, r *http.Request)  {
	db,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/hehe?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	var who string
	var place string
	var ps string
	var holddays int

	rows,err := db.Query("select * from pingpang")
	if err!=nil{
		fmt.Println(err)
	}


	
	
	var items []Pingpang
	for rows.Next(){
		rows.Scan(&who,&place,&holddays,&ps)
		fmt.Println(who,place,holddays,ps)
		pingpang := Pingpang{who,place,holddays,ps}
		items = append(items,pingpang)

	}

	responseJson := ResponseJson{0,items}

	js,err := json.Marshal(responseJson)
	if err !=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.Write(js)

	defer rows.Close()
}