// Package bot is the top level container for ci-bot-01 code
package bot

import (
	"encoding/json"
	"net/http"

	"github.com/google/go-github/github"
)

// WebHookType represents the subset of GitHub web hooks activities on which we wish to take action
type WebHookType int

const (
	webHookTypePing WebHookType = iota
	webHookTypeIssue
	webHookTypeCommits
)

// WebHookRouter accepts a request sent by the GitHub webhooks API server, identifies the
// type of webhook and then sends to the proper intenral handler
func WebHookRouter(r *http.Request) {

	bodyContent := getBodyContent(r)
	webHookInfo := parseWebHookInfo(bodyContent)

	if isIssueAction(webHookInfo) {
		var ievent github.IssueActivityEvent
		json.Unmarshal(bodyContent, &ievent)
		rejectIssue(r, ievent)
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

func parseWebHookInfo(content []byte) map[string]interface{} {
	var webhookJSON interface{}
	err := json.Unmarshal(content, &webhookJSON)
	if err != nil {
		return nil
	}
	wh := webhookJSON.(map[string]interface{})
	return wh
}
