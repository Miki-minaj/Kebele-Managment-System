package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
	"github.com/miki-minaj/Kebele-Managment-System/handler"
	mrepim "github.com/miki-minaj/Kebele-Managment-System/menu/repository"
	msrvim "github.com/miki-minaj/Kebele-Managment-System/menu/service"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "miki"
	dbname   = "test"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))

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
func main() {
	var db *sql.DB
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		//panic(err)
	}

	catagoryRepo := mrepim.NewCategoryRepositoryImpl(db)

	categoryServ := msrvim.NewCategoryService(catagoryRepo)
	adminhandler := handler.NewAdminCategoryHandler(templ, categoryServ)

	//http.HandleFunc("/f", insert)
	fs := http.FileServer(http.Dir("./templates/asset"))
	http.HandleFunc("/", index)
	http.HandleFunc("/f", adminhandler.AdminCategoriesNew)
	http.HandleFunc("/s", adminhandler.AdminCategories)
	http.Handle("/asset/", http.StripPrefix("/asset/", fs))
	http.ListenAndServe(":8080", nil)
}
