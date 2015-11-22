// Package bot is the top level container for ci-bot-01 code
package bot

import (
	"net/http"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"google.golang.org/appengine"
)

func rejectIssue(r *http.Request, ie github.IssueActivityEvent) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("CI_BOT_GITHUB_KEY")})
	tc := oauth2.NewClient(appengine.NewContext(r), ts)
	client := github.NewClient(tc)
	client.Issues.AddLabelsToIssue("spohnan", *ie.Repo.Name, *ie.Issue.Number, []string{"wontfix"})
}
