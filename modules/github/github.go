package github

import "github.com/google/go-github/github"
import "os"
import "golang.org/x/oauth2"

//APIKey is github credetials
var APIKey string
// NewClient returns a configured "github.com/google/go-github/github" Client
// with a token
func NewClient() *github.Client {
	// TODO: enable github authenticated client

	// TODO: GITHUB_TOKEN should be taken from an environment variable
    if (APIKey == "") {
         APIKey = os.Getenv("GITHUB_TOKEN")
    }
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: APIKey},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	return github.NewClient(tc)
}

//SetAPIKey used to set github api key
func SetAPIKey(key string) int {
    APIKey = key
    return 0;
}