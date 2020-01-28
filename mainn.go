package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	//mrepim "github.com/miki-minaj/web-prog-go-sample-master/comment/repository"
	"github.com/miki-minaj/Kebele-Managment-System/entity"
	"github.com/miki-minaj/Kebele-Managment-System/handler"

	api "github.com/miki-minaj/Kebele-Managment-System/handler/api"
	//msrvim "github.com/miki-minaj/web-prog-go-sample-master/menu/service"
	pim "github.com/miki-minaj/Kebele-Managment-System/menu/repository"
	vim "github.com/miki-minaj/Kebele-Managment-System/menu/service"
	"github.com/miki-minaj/Kebele-Managment-System/rtoken"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	urepimp "github.com/miki-minaj/Kebele-Managment-System/user/repository"
	usrvimp "github.com/miki-minaj/Kebele-Managment-System/user/service"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "miki"
	dbname   = "test"
)

func createTables(dbconn *gorm.DB) []error {
	errs := dbconn.CreateTable(&entity.User{}, &entity.Role{}, &entity.Session{}).GetErrors()
	if errs != nil {
		return errs
	}
	return nil
}

var templ = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	var db *sql.DB
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	csrfSignKey := []byte(rtoken.GenerateRandomID(32))
	tmpl := template.Must(template.ParseGlob("templatess/*.html"))

	dbconn, err := gorm.Open("postgres", "postgres://postgres:miki@localhost/test?sslmode=disable")
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		//panic(err)
	}
	if err != nil {
		panic(err)
	}
	createTables(dbconn)
	defer dbconn.Close()

	sessionRepo := urepimp.NewSessionGormRepo(dbconn)
	sessionSrv := usrvimp.NewSessionService(sessionRepo)

	userRepo := urepimp.NewUserGormRepo(dbconn)
	userServ := usrvimp.NewUserService(userRepo)

	roleRepo := urepimp.NewRoleGormRepo(dbconn)
	roleServ := usrvimp.NewRoleService(roleRepo)

	sess := configSess()
	uh := handler.NewUserHandler(tmpl, userServ, sessionSrv, roleServ, sess, csrfSignKey)

	fs := http.FileServer(http.Dir("./templates/asset"))
	http.Handle("/asset/", http.StripPrefix("/asset/", fs))

	// http.HandleFunc("/", mh.Index)
	// http.HandleFunc("/about", mh.About)
	// http.HandleFunc("/contact", mh.Contact)
	// http.HandleFunc("/menu", mh.Menu)
	// http.Handle("/admin", uh.Authenticated(uh.Authorized(http.HandlerFunc(mh.Admin))))

	catagoryRepo := pim.NewCategoryRepositoryImpl(db)
	categoryServ := vim.NewCategoryService(catagoryRepo)
	//adminhandler := handler.NewAdminCategoryHandler(templ, categoryServ)
	navhandler := handler.NewAdminCategoryHandler(templ, categoryServ, csrfSignKey)
	// navhandler := handler.NewAdminCategoryHandler(templ, categoryServ)

	apiinforepo := pim.NewCategoryRepositoryImpl(db)
	apiinfoserv := vim.NewCategoryService(apiinforepo)
	apiadmininfohandler := api.NewAdminCategoryHandler(apiinfoserv)

	router := httprouter.New()
	router.POST("/v1/admin/infos", apiadmininfohandler.AdminCategoriesNew)
	router.GET("/v1/admin/infos", apiadmininfohandler.AdminCategories)

	http.Handle("/end", uh.Authenticated(uh.Authorized(http.HandlerFunc(navhandler.Index))))
	http.Handle("/reg", uh.Authenticated(uh.Authorized(http.HandlerFunc(navhandler.AdminCategoriesNew))))
	http.Handle("/s", uh.Authenticated(uh.Authorized(http.HandlerFunc(navhandler.AdminSearch))))
	http.Handle("/all", uh.Authenticated(uh.Authorized(http.HandlerFunc(navhandler.AdminCategories))))

	http.HandleFunc("/login", uh.Login)
	//http.HandleFunc("/f", navhandler.Index)
	http.Handle("/logout", uh.Authenticated(http.HandlerFunc(uh.Logout)))
	http.HandleFunc("/signup", uh.Signup)

	http.ListenAndServe(":8181", nil)
}

func configSess() *entity.Session {
	tokenExpires := time.Now().Add(time.Minute * 30).Unix()
	sessionID := rtoken.GenerateRandomID(32)
	signingString, err := rtoken.GenerateRandomString(32)
	if err != nil {
		panic(err)
	}
	signingKey := []byte(signingString)

	return &entity.Session{
		Expires:    tokenExpires,
		SigningKey: signingKey,
		UUID:       sessionID,
	}
}
