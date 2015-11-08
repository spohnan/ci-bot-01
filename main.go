package ci_bot

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

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		//return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)

		fmt.Fprintf(w, "userip: %q is not IP:port", r.RemoteAddr)
	}

	userIP := net.ParseIP(ip)
	// if userIP == nil {
	// 	//return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
	// 	fmt.Fprintf(w, "userip: %q is not IP:port", r.RemoteAddr)
	// 	return
	// }

	fmt.Fprint(w, userIP)
}
