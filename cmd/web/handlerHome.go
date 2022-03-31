package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	flash := app.session.PopString(r, "flash")
	app.render(w, r, "home.page.tmpl", &templateData{
		Flash: flash,
	})
}

func (app *application) redirect(w http.ResponseWriter, r *http.Request) {
	flash := app.session.PopString(r, "flash")
	app.render(w, r, "dimash.page.tmpl", &templateData{
		Flash: flash,
	})
}
