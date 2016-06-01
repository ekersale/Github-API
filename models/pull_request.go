package models

import "github.com/google/go-github/github"

type PullRequest struct {
    URL             *string      `json:"url",omitempty`
    Html_Url        *string      `json:"html_url",omitempty`
    Diff_Url        *string      `json:"diff_url",omitempty`
    Patch_Url       *string      `json;"diff_url",omitempty`
}

func setPullRequest(github_pullRequest github.PullRequestLinks)(PullRequest){
	tmp := PullRequest{}
    tmp.URL = github_pullRequest.URL
    tmp.Html_Url = github_pullRequest.HTMLURL
    tmp.Diff_Url = github_pullRequest.DiffURL
    tmp.Patch_Url = github_pullRequest.PatchURL
	return tmp
}