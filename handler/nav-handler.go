package handler

import (
	"net/http"
	"text/template"

	"github.com/miki-minaj/Kebele-Managment-System/menu"
)

// MenuHandler handles menu related requests
type MenuHandler struct {
	tmpl        *template.Template
	categorySrv menu.CategoryService
}

// NewMenuHandler initializes and returns new MenuHandler
func NewMenuHandler(t *template.Template, cs menu.CategoryRepository) *MenuHandler {
	return &MenuHandler{tmpl: t, categorySrv: cs}
}

// NewMenuHandler initializes and returns new MenuHandler
func (mh *MenuHandler) Index(w http.ResponseWriter, r *http.Request) {

	//info := personalinfo{"Miki",20}
	mh.tmpl.ExecuteTemplate(w, "kebele.html", "hey")

}
