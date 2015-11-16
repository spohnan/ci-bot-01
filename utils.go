// Package bot is the top level container for ci-bot-01 code
package bot

import (
	"net/http"
)

func getBodyContent(r *http.Request) []byte {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	return body
}

func mapContainsKey(m map[string]interface{}, k string) bool {
	_, ok := m[k]
	return ok
}
