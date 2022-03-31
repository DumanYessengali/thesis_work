package main

import (
	"net/http"
	"thesis_work/pkg/forms"
)

func (app *application) signInForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) signIn(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	username := form.Get("username")
	password := form.Get("password")
	if username != "admin" || password != "123" {
		form.Errors.Add("generic", "Username or Password is incorrect")
		app.render(w, r, "login.page.tmpl", &templateData{Form: form})
		return

	}

	app.session.Put(r, "authenticatedUserID", 1)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	app.session.Remove(r, "authenticatedUserID")
	app.session.Put(r, "flash", "You've been logged out successfully!")
	http.Redirect(w, r, "/signin", http.StatusSeeOther)
}
