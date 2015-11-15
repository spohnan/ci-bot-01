// Package bot is the top level container for ci-bot-01 code
package bot

import (
	"io/ioutil"
	"testing"
)

type testResult struct {
	isCommitType bool
	isIssueType  bool
	isPingType   bool
}

type webhookInfo struct {
	whType   WebHookType
	fileName string
	whMap    map[string]interface{}
	testResult
}

// Store all the needed information to test event type detection
var whTestData = []webhookInfo{
	{webHookTypeCommits, "commit.json", parseWebHookInfo(getWebhookExample("commit.json")), testResult{isCommitType: true}},
	{webHookTypeIssue, "issue_closed.json", parseWebHookInfo(getWebhookExample("issue_closed.json")), testResult{isIssueType: true}},
	{webHookTypeIssue, "issue_opened.json", parseWebHookInfo(getWebhookExample("issue_opened.json")), testResult{isIssueType: true}},
	{webHookTypePing, "ping.json", parseWebHookInfo(getWebhookExample("ping.json")), testResult{isPingType: true}},
}

// Test all available example event captures against all types of web hook
// types to ensure that each type is identified correctly
func TestEventTypeDetection(t *testing.T) {
	for _, whi := range whTestData {

		if isPingAction(whi.whMap) && !whi.testResult.isPingType {
			t.Fatalf("%s was incorrectly identified as event type %v", whi.fileName, whi.whType)
		}
		if isCommitsAction(whi.whMap) && !whi.testResult.isCommitType {
			t.Fatalf("%s was incorrectly identified as event type %v", whi.fileName, whi.whType)
		}
		if isIssueAction(whi.whMap) && !whi.testResult.isIssueType {
			t.Fatalf("%s was incorrectly identified as event type %v", whi.fileName, whi.whType)
		}

	}
}

// A helper function to retrieve the contents of test data files
func getWebhookExample(filename string) string {
	contentBytes, err := ioutil.ReadFile("./.hook-examples/" + filename)
	if err != nil {
		return ""
	}
	return string(contentBytes)
}
