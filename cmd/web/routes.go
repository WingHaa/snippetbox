package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	fileServer := http.FileServer(neuteredFs{http.Dir("./ui/static/")})

	mux := http.NewServeMux()
	mux.Handle("GET /static", http.NotFoundHandler())
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	// mux.HandleFunc("GET /snippet/create", app.snipppetCreate)
	mux.HandleFunc("POST /snippet/create", app.snipppetCreatePost)

	return app.recoverPanic(app.logRequest(commonHeaders(mux)))
}
