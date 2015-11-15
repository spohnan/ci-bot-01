// Package bot is the top level container for ci-bot-01 code
package bot

import (
	"encoding/json"
	"net/http"
)

type WebHookType int

const (
	webHookTypePing WebHookType = iota
	webHookTypeIssue
	webHookTypeCommits
)

// WebHookRouter accepts a request sent by the GitHub webhooks API server, identifies the
// type of webhook and then sends to the proper intenral handler
func WebHookRouter(r *http.Request) {

	webHookInfo := parseWebHookInfo(getBodyContent(r))

	if isIssueAction(webHookInfo) {

	}
}

func isCommitsAction(wh map[string]interface{}) bool {
	return mapContainsKey(wh, "commits")
}

func isIssueAction(wh map[string]interface{}) bool {
	return mapContainsKey(wh, "issue")
}

func isPingAction(wh map[string]interface{}) bool {
	return mapContainsKey(wh, "zen")
}

func parseWebHookInfo(content string) map[string]interface{} {
	var webhookJSON interface{}
	err := json.Unmarshal([]byte(content), &webhookJSON)
	if err != nil {
		return nil
	}
	wh := webhookJSON.(map[string]interface{})
	return wh
}
