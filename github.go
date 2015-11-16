// Package bot is the top level container for ci-bot-01 code
package bot

import (
	"net/http"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func rejectIssue(r *http.Request, ie github.IssueActivityEvent) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("CI_BOT_GITHUB_KEY")})
	tc := oauth2.NewClient(appengine.NewContext(r), ts)
	client := github.NewClient(tc)

	log.Infof(c, "Name %s Num %d", *ie.Repo.Name, *ie.Issue.Number)
	_, _, err := client.Issues.AddLabelsToIssue("spohnan", *ie.Repo.Name, *ie.Issue.Number, []string{"wontfix"})
	if err != nil {
		log.Infof(c, "%s", err)
	}
}
