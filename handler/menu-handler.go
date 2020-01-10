package handler

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/miki-minaj/Kebele-Managment-System/entity"
	"github.com/miki-minaj/Kebele-Managment-System/menu"
)

type AdminCategoryHandler struct {
	tmpl        *template.Template
	categorySrv menu.CategoryRepository
}

// NewAdminCategoryHandler initializes and returns new AdminCateogryHandler
func NewAdminCategoryHandler(t *template.Template, cs menu.CategoryRepository) *AdminCategoryHandler {
	return &AdminCategoryHandler{tmpl: t, categorySrv: cs}
}

// AdminCategoriesNew hanlde requests on route /admin/categories/new
func (ach *AdminCategoryHandler) AdminCategoriesNew(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		fmt.Println("here")
		ctg := &entity.Category{}
		ctg.Name = r.FormValue("fname")
		ctg.ID = r.FormValue("iid")
		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()
		ctg.Image = fh.Filename

		writeFile(&mf, fh.Filename)
		//ctg.Description = r.FormValue("description")

		_, errs := ach.categorySrv.StoreCategory(ctg)

		if errs != nil {
			fmt.Println("got here")
			panic(errs)
		}

		http.Redirect(w, r, "REG.html", http.StatusSeeOther)

	} else {
		fmt.Println("no here")
		ach.tmpl.ExecuteTemplate(w, "REG.html", nil)

	}
}

// AdminCategories handle requests on route /admin/categories
func (ach *AdminCategoryHandler) AdminCategories(w http.ResponseWriter, r *http.Request) {
	categories, errs := ach.categorySrv.Categories()
	if errs != nil {
		panic(errs)
	}
	ach.tmpl.ExecuteTemplate(w, "REG.html", categories)
}

func writeFile(mf *multipart.File, fname string) {

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, "src", "github.com", "miki-minaj", fname)
	image, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	defer image.Close()
	io.Copy(image, *mf)
}
