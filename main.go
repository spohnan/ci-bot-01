// Package bot is the top level container for ci-bot-01 code
package bot

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", mainHandler)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {

	if IsRequestAuthorized(r.RemoteAddr) {
		fmt.Fprintf(w, "Method %s - URI %s - %s", r.Method, r.RequestURI, r.RemoteAddr)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

}
