package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

type info struct {
	id   string
	name string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "test"
)

var templ = template.Must(template.ParseFiles("newtest.html"))

func index(w http.ResponseWriter, r *http.Request) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	rows, err := db.Query("SELECT id,name FROM infos")
	infos := make([]info, 0)
	for rows.Next() {
		in := info{}
		//var id int
		//var name string

		err = rows.Scan(&in.id, &in.name)

		infos = append(infos, in)

		//panic(err)
		fmt.Println("id | name ")
		fmt.Println(in.id, in.name)
	}

	//w.Write([]byte("<h1>Hello World!</h1>"))
	//info := personalinfo{"Miki",20}
	templ.Execute(w, infos)

}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	http.ListenAndServe(":8080", mux)

}
