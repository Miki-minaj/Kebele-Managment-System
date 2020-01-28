package handler

import (
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
