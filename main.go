package main

import (
	"net/http"
	"net/url"

	"encoding/base64"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	var (
		redirectLink = "https://google.com/"
	)

	r.Get("/ref_{base64}", func(w http.ResponseWriter, r *http.Request) {

		data, err := base64.StdEncoding.DecodeString(chi.URLParam(r, "base64"))
		targetLink := string(data)
		//fmt.Printf("%q\n", err, string(data))

		if err != nil {
			http.Redirect(w, r, redirectLink, http.StatusMovedPermanently)
		}

		if !IsUrl(targetLink) {
			http.Redirect(w, r, redirectLink, http.StatusMovedPermanently)
		}

		http.Redirect(w, r, targetLink, http.StatusMovedPermanently)
	})

	http.ListenAndServe(":80", r)
}

//env GOOS=windows GOARCH=amd64 go build rdir
//env GOOS=linux GOARCH=amd64 go build -v ./rdir main.go

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
