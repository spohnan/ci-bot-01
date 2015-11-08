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
	// TODO: Add check of source IP
	// https://help.github.com/articles/what-ip-addresses-does-github-use-that-i-should-whitelist/
	// https://groups.google.com/forum/embed/#!topic/golang-nuts/Usu-B5rcCJs
	fmt.Fprintf(w, "Method %s - URI %s - %s", r.Method, r.RequestURI, r.RemoteAddr)
}
