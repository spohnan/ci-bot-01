// Package bot is the top level container for ci-bot-01 code
package bot

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

const (
	HealthOK = "OK"
)

func initWebHandlers() {
	http.HandleFunc("/health", authWrapper(healthHandler))
	http.HandleFunc("/webhook", authWrapper(webHookHandler))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, HealthOK)
}

func webHookHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	log.Infof(c, "%s", string(body))
}

func authWrapper(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if IsRequestAuthorized(r.RemoteAddr) {
			h(w, r)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	}
}
