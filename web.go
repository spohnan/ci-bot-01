// Package bot is the top level container for ci-bot-01 code
package bot

import (
	"fmt"
	"log"
	"net/http"
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
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	log.Printf("%s", string(body))
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
