package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	dynamicMiddleware := alice.New(app.session.Enable)

	mux := pat.New()

	mux.Get("/signin", dynamicMiddleware.Append().ThenFunc(app.signInForm))
	mux.Post("/signin", dynamicMiddleware.ThenFunc(app.signIn))
	mux.Post("/logout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.logoutUser))

	mux.Get("/", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.home))
	mux.Get("/redirect", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.redirect))

	mux.Get("/ip-checker", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.ipCheckerForm))
	mux.Post("/ip-checker", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.ipChecker))
	mux.Get("/ip-info", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.ipInfo))

	mux.Get("/grok", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.grokInfo))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
