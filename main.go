package main

import (
	"database/sql"
	"fmt"
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
	password = "1234"
	dbname   = "test"
)

//hena
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
	templ := template.Must(template.ParseGlob("templates/*.html"))
	catagoryRepo := mrepim.NewCategoryRepositoryImpl(db)

	categoryServ := msrvim.NewCategoryService(catagoryRepo)
	adminhandler := handler.NewAdminCategoryHandler(templ, categoryServ)

	//http.HandleFunc("/f", insert)
}
