package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/miki-minaj/Kebele-Managment-System/entity"
	"github.com/miki-minaj/Kebele-Managment-System/menu"
)

type AdminCategoryHandler struct {
	categorySrv menu.CategoryRepository
}

// NewAdminCategoryHandler initializes and returns new AdminCateogryHandler
func NewAdminCategoryHandler(cs menu.CategoryRepository) *AdminCategoryHandler {
	return &AdminCategoryHandler{categorySrv: cs}
}

// AdminCategoriesNew hanlde requests on route /admin/categories/new
func (ach *AdminCategoryHandler) AdminCategoriesNew(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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

		// mf, fh, err := r.FormFile("catimg")
		// if err != nil {
		// 	panic(err)
		// }
		// defer mf.Close()
		// ctg.Image = fh.Filename

		// writeFile(&mf, fh.Filename)
		//ctg.Description = r.FormValue("description")

		_, errs := ach.categorySrv.StoreCategory(ctg)

		if errs != nil {
			fmt.Println("got here")
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		p := fmt.Sprintf("/v1/admin/comments/%d", ctg.ID)
		w.Header().Set("Location", p)
		w.WriteHeader(http.StatusCreated)
		return
		//http.Redirect(w, r, "REG.html", http.StatusSeeOther)

	} else {
		fmt.Println("no here")
		//ach.tmpl.ExecuteTemplate(w, "REG.html", nil)

	}
	// ctg.AGE = r.FormValue("age")
	// ctg.Occupation = r.FormValue("occu")
	// ctg.Relegion = r.FormValue("relegion")
	// ctg.Nationality = r.FormValue("nationality")
}

// AdminCategories handle requests on route /admin/categories
func (ach *AdminCategoryHandler) AdminCategories(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	categories, errs := ach.categorySrv.Categories()
	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		//panic(errs)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(categories, "", "\t\t")
	//ach.tmpl.ExecuteTemplate(w, "REG.html", categories)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
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
