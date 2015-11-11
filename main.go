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

	// Restrict access to the web app based on remote IP
	// Whitelist is set in app.yaml using the CI_BOT_IP_WHITELIST env variable
	if IsRequestAuthorized(r.RemoteAddr) {
		fmt.Fprintf(w, "Method %s - URI %s - %s", r.Method, r.RequestURI, r.RemoteAddr)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

}
