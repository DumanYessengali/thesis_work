package main

import (
	"net"
	"net/http"
	"strconv"
	"thesis_work/pkg/forms"
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

func (app *application) ipCheckerForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "ip-checker.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) ipChecker(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	ip := form.Get("ip")

	if net.ParseIP(ip) == nil {
		app.render(w, r, "error.page.tmpl", &templateData{Form: form})
		return
	}
	http.Redirect(w, r, "/ip-info", http.StatusSeeOther)
	return
}

func (app *application) ipInform(w http.ResponseWriter, r *http.Request) {
	_, err := strconv.Atoi(r.URL.Query().Get("ip"))
	if err != nil {
		app.notFound(w)
		return
	}

	flash := app.session.PopString(r, "flash")
	app.render(w, r, "ip-info.page.tmpl", &templateData{
		Flash: flash,
	})
}

func (app *application) ipInfo(w http.ResponseWriter, r *http.Request) {

}
