// Package bot is the top level container for ci-bot-01 code
package bot

import (
	"net/http"
)

func getBodyContent(r *http.Request) string {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	return string(body)
}

func mapContainsKey(m map[string]interface{}, k string) bool {
	_, ok := m[k]
	return ok
}
