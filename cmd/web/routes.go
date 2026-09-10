package main

import (
	"net/http"
	"snippetbox/ui"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	dinamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	mux.Handle("GET /{$}", dinamic.ThenFunc(app.home))
	mux.Handle("GET /snippet/view/{id}", dinamic.ThenFunc(app.snippetView))
	mux.Handle("GET /user/signup", dinamic.ThenFunc(app.userSignup))
	mux.Handle("POST /user/signup", dinamic.ThenFunc(app.userSignupPost))
	mux.Handle("GET /user/login", dinamic.ThenFunc(app.userLogin))
	mux.Handle("POST /user/login", dinamic.ThenFunc(app.userLoginPost))
	mux.HandleFunc("GET /ping", app.ping)
	protected := dinamic.Append(app.requireAuthentication)

	mux.Handle("POST /user/logout", protected.ThenFunc(app.userLogoutPost))
	mux.Handle("GET /snippet/create", protected.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", protected.ThenFunc(app.snippetCreatePost))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)
	return standard.Then(mux)
}
