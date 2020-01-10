package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

type info struct {
	Name string
	ID   string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "test"
)

//checkyg
var templ = template.Must(template.ParseFiles("kebele.html", "REG.html"))
var db *sql.DB

func init() {
	var err error
	fmt.Println("This will get called on main initialization")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		//panic(err)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		//panic(err)
	}

	fmt.Println("Successfully connected!")
}
func index(w http.ResponseWriter, r *http.Request) {
	// var err error

	// rows, err := db.Query("SELECT id,name FROM infos")
	// if err != nil {
	// 	//fmt.Println("error")
	// }
	// xc := []info{}
	// for rows.Next() {
	// 	in := info{}
	// 	//var id int
	// 	//var name string

	// 	rows.Scan(&in.ID, &in.Name)

	// 	xc = append(xc, in)

	// 	//panic(err)
	// 	//fmt.Println("id | name ")
	// 	//fmt.Println(in.ID, in.Name)
	// }

	//w.Write([]byte("<h1>Hello World!</h1>"))
	//info := personalinfo{"Miki",20}
	templ.ExecuteTemplate(w, "kebele.html", "hey")

}

func insert(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	id := r.FormValue("iid")
	name := r.FormValue("fname")
	fmt.Println(name)
	fmt.Println(id)
	sqlStatement := `
	INSERT INTO infos (id,name)
	VALUES (` + id + `,'` + name + `')`

	db.Query(sqlStatement)
	if err != nil {
		log.Fatal("no open , error")
	}
	//templ.ExecuteTemplate(w, "REG.html", "h")
	//fmt.Println("This will get called on main initialization")
}

func show(w http.ResponseWriter, r *http.Request) {
	var err error

	rows, err := db.Query("SELECT id,name FROM infos")
	if err != nil {
		//fmt.Println("error")
	}
	xc := []info{}
	for rows.Next() {
		in := info{}
		//var id int
		//var name string

		rows.Scan(&in.ID, &in.Name)

		xc = append(xc, in)

		//panic(err)
		//fmt.Println("id | name ")
		//fmt.Println(in.ID, in.Name)
	}
	templ.ExecuteTemplate(w, "REG.html", xc)
}

func main() {
	//defer db.Close()
	//mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("asset"))
	http.HandleFunc("/", index)
	http.HandleFunc("/f", insert)
	http.HandleFunc("/r", show)
	http.Handle("/asset/", http.StripPrefix("/asset/", fs))
	http.ListenAndServe(":8080", nil)

}
