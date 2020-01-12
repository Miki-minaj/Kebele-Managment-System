package handler

import (
	"fmt"
	"net/http"
	"strconv"
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
		ctg.Mothername = r.FormValue("mother'sname")
		ctg.AGE = r.FormValue("age")
		ctg.Occupation = r.FormValue("occu")
		ctg.Relegion = r.FormValue("relegion")
		ctg.Nationality = r.FormValue("nationality")
		ctg.Phonenumber = r.FormValue("phonenum")
		ctg.Emergencyname = r.FormValue("emergencyn")
		ctg.Emergencyphone = r.FormValue("emergencyp")

		

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
	// ctg.AGE = r.FormValue("age")
	// ctg.Occupation = r.FormValue("occu")
	// ctg.Relegion = r.FormValue("relegion")
	// ctg.Nationality = r.FormValue("nationality")
}

// AdminCategories handle requests on route /admin/categories
func (ach *AdminCategoryHandler) AdminCategories(w http.ResponseWriter, r *http.Request) {
	categories, errs := ach.categorySrv.Categories()
	if errs != nil {
		panic(errs)
	}
	ach.tmpl.ExecuteTemplate(w, "REG.html", categories)
}

// AdminCategoriesDelete handle requests on route /admin/categories/delete
func (ach *AdminCategoryHandler) AdminCategoriesDelete(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		err = ach.categorySrv.DeleteCategory(id)

		if err != nil {
			panic(err)
		}

	}

	http.Redirect(w, r, "REG.html", http.StatusSeeOther)
}

// func writeFile(mf *multipart.File, fname string) {

// 	wd, err := os.Getwd()

// 	if err != nil {
// 		panic(err)
// 	}

// 	path := filepath.Join(wd, "src", "github.com", "miki-minaj", fname)
// 	image, err := os.Create(path)

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer image.Close()
// 	io.Copy(image, *mf)
// }
