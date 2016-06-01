package github

import "github.com/google/go-github/github"
import "os"
import "golang.org/x/oauth2"


// NewClient returns a configured "github.com/google/go-github/github" Client
// with a token
func NewClient() *github.Client {
	// TODO: enable github authenticated client

	// TODO: GITHUB_TOKEN should be taken from an environment variable
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	return github.NewClient(tc)
}
