package handler

import (
	"net/http"
	"net/url"

	"github.com/miki-minaj/Kebele-Managment-System/form"
	"github.com/miki-minaj/Kebele-Managment-System/rtoken"
	"github.com/miki-minaj/Kebele-Managment-System/session"
	"golang.org/x/crypto/bcrypt"
)

func (uh *UserHandler) html(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		loginForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		uh.tmpl.ExecuteTemplate(w, "login.html", loginForm)
		return
	}

	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		loginForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		usr, errs := uh.userService.UserByEmail(r.FormValue("email"))
		if len(errs) > 0 {
			loginForm.VErrors.Add("generic", "Your email address or password is wrong")
			uh.tmpl.ExecuteTemplate(w, "login.html", loginForm)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(r.FormValue("password")))
		if err == bcrypt.ErrMismatchedHashAndPassword {
			loginForm.VErrors.Add("generic", "Your email address or password is wrong")
			uh.tmpl.ExecuteTemplate(w, "login.html", loginForm)
			return
		}

		uh.loggedInUser = usr
		claims := rtoken.Claims(usr.Email, uh.userSess.Expires)
		session.Create(claims, uh.userSess.UUID, uh.userSess.SigningKey, w)
		newSess, errs := uh.sessionService.StoreSession(uh.userSess)
		if len(errs) > 0 {
			loginForm.VErrors.Add("generic", "Failed to store session")
			uh.tmpl.ExecuteTemplate(w, "login.html", loginForm)
			return
		}
		uh.userSess = newSess
		roles, _ := uh.userService.UserRoles(usr)
		if uh.checkAdmin(roles) {
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
