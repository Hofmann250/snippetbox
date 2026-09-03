package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	dinamic := alice.New(app.sessionManager.LoadAndSave)

	mux.Handle("GET /{$}", dinamic.ThenFunc(app.home))
	mux.Handle("GET /snippet/view/{id}", dinamic.ThenFunc(app.snippetView))
	mux.Handle("GET /snippet/create", dinamic.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", dinamic.ThenFunc(app.snippetCreatePost))

	mux.Handle("GET /user/signup", dinamic.ThenFunc(app.userSignup))
	mux.Handle("POST /user/signup", dinamic.ThenFunc(app.userSignupPost))
	mux.Handle("GET /user/login", dinamic.ThenFunc(app.userLogin))
	mux.Handle("POST /user/login", dinamic.ThenFunc(app.userLoginPost))
	mux.Handle("POST /user/logout", dinamic.ThenFunc(app.userLogoutPost))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)
	return standard.Then(mux)
}
