// Package bot is the top level container for ci-bot-01 code
package bot

import (
	"fmt"
	"net"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// TODO: Add check of source IP
	// https://help.github.com/articles/what-ip-addresses-does-github-use-that-i-should-whitelist/
	// https://groups.google.com/forum/embed/#!topic/golang-nuts/Usu-B5rcCJs

	fmt.Fprint(w, "<html><head/><body>Hello, world!<br /><br /></body></html>")
}
