package handler

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/jung-kurt/gofpdf"
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
func (ach *AdminCategoryHandler) Index(w http.ResponseWriter, r *http.Request) {

	//info := personalinfo{"Miki",20}
	ach.tmpl.ExecuteTemplate(w, "kebele.html", "hey")

}

// AdminCategoriesNew hanlde requests on route /admin/categories/new
func (ach *AdminCategoryHandler) AdminCategoriesNew(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		fmt.Println("pass")
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
		fmt.Println("here")
		mf, fh, err := r.FormFile("idphoto")

		if err != nil {
			fmt.Println("here")
			panic(err)
		}
		fmt.Println("he")
		defer mf.Close()

		ctg.Image = fh.Filename

		writeFile(&mf, fh.Filename)

		_, errs := ach.categorySrv.StoreCategory(ctg)

		if errs != nil {
			fmt.Println("error")
			panic(errs)
		} else {
			err := GeneratePdf(ctg.ID+".pdf", ctg.Name, ctg.AGE, ctg.Sex, ctg.Mothername, ctg.Phonenumber, ctg.Nationality, ctg.Relegion, ctg.Occupation, ctg.Emergencyname, ctg.Emergencyphone, ctg.Image)
			fmt.Println("pdf generated succesfully")
			if err != nil {
				panic(err)
			}
		}

		//http.Redirect(w, r, "/s", http.StatusSeeOther)
		ach.tmpl.ExecuteTemplate(w, "pdf.html", "asset/pdfs/"+ctg.ID+".pdf")

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
func (ach *AdminCategoryHandler) AdminSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello i got the problem here")
	if r.Method == http.MethodGet {
		fmt.Println("1")
		nameraw := r.URL.Query().Get("nname")
		fmt.Println("2")
		//name, errs := strconv.Atoi(nameraw)
		fmt.Println(nameraw)
		categories, errs := ach.categorySrv.Category(nameraw)
		if errs != nil {
			fmt.Println("come on bruh")
			panic(errs)
		}
		//fmt.Println(categories.Mothername)
		//http.Redirect(w, r, "REG.html", http.StatusSeeOther)
		ach.tmpl.ExecuteTemplate(w, "REG.html", categories)
		//fmt.Println(categories.Name)
		//http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		fmt.Println("hello i got the problem here")
		fmt.Println("3")
	}
	fmt.Println("4")
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

func writeFile(mf *multipart.File, fname string) {

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, "templates", fname)
	fmt.Println(path)
	fmt.Println("----------------------------------------------------------------")
	image, err := os.Create(path)

	if err != nil {
		fmt.Println("No way buety from the west")
		panic(err)
	}
	defer image.Close()
	io.Copy(image, *mf)
}
func GeneratePdf(filename string, n string, a string, s string, mn string, p string, na string, r string, o string, en string, ep string, img string) error {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	//pdf.SetFont("Arial", "B", 16)
	pdf.SetFont("Arial", "", 12)

	// CellFormat(width, height, text, border, position after, align, fill, link, linkStr)

	//pdf.CellFormat(100, 7, "ID CARD", "0", 0, "CM", false, 0, "")
	pdf.MultiCell(120, 4, "\t \t \n \n \n \n Name: "+n+" \n Age: "+a+" \n Sex: "+s+" \n Mother's Name: "+mn+"\n Phone no: "+p+" \n Nationality: "+na+"\nRelegion: "+r+" \n Ocuupation: "+o+" \n Emergency Contanct Name: "+en+" \n Emergency Contac Phone: "+ep, "0", "L", false)
	//ImageOptions(src, x, y, width, height, flow, options, link, linkStr)
	pdf.ImageOptions(
		"templates/"+img,
		90, 25,
		20, 20,
		false,
		gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
		0,
		"",
	)
	fmt.Println("waiting...")
	//return pdf.OutputFileAndClose(filename)
	return pdf.OutputFileAndClose("templates/asset/pdfs/" + filename)
}
