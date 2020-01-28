package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/miki-minaj/Kebele-Managment-System/handler"
	api "github.com/miki-minaj/Kebele-Managment-System/handler/api"
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
	//adminhandler := handler.NewAdminCategoryHandler(templ, categoryServ)
	navhandler := handler.NewAdminCategoryHandler(templ, categoryServ)

	apiinforepo := mrepim.NewCategoryRepositoryImpl(db)
	apiinfoserv := msrvim.NewCategoryService(apiinforepo)
	apiadmininfohandler := api.NewAdminCategoryHandler(apiinfoserv)

	router := httprouter.New()

	router.POST("/v1/admin/infos", apiadmininfohandler.AdminCategoriesNew)
	router.GET("/v1/admin/infos", apiadmininfohandler.AdminCategories)
	//http.HandleFunc("/f", insert)
	fs := http.FileServer(http.Dir("./templates/asset"))
	http.HandleFunc("/", navhandler.Index)
	http.HandleFunc("/f", navhandler.AdminCategoriesNew)
	http.HandleFunc("/s", navhandler.AdminSearch)
	http.HandleFunc("/sa", navhandler.AdminCategories)
	http.Handle("/asset/", http.StripPrefix("/asset/", fs))
	http.ListenAndServe(":8080", nil)

}
