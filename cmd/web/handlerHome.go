package main

import (
	"errors"
	"net/http"
	"thesis_work/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	app.student.GetStudentId()
	syllabus, err := app.student.GetNameSyllabusWithStudent()
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	flash := app.session.PopString(r, "flash")
	app.render(w, r, "home.page.tmpl", &templateData{
		Flash:    flash,
		Syllabus: syllabus,
	})
}
