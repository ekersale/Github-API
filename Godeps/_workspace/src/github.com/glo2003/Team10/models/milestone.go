package models

import "github.com/google/go-github/github"
import "time"

type Milestone struct {
    URL             *string      `json:"url",omitempty`
    Html_Url        *string      `json:"html_url",omitempty`
    Labels_Url      *string      `json:"labels_url",omitempty`
    ID              int         `json:"id",omitempty`
    Number          int         `json:"number",omitempty`
    State           *string      `json:"state",omitempty`
    Title           *string      `json:"title",omitempty`
    Description     *string      `json:"description",omitempty`
    Creator         User        `json:"creator",omitempty`
    Open_Issues     int         `json:"open_issues",omitempty`
    Closed_Issues   int         `json:"closed_issues",omitempty`
    Created_At      *time.Time   `json:"created_at",omitempty`
    Updated_At      *time.Time   `json:"updated_at",omitempty`
    Closed_At       *time.Time   `json:"closed_at",omitempty`
    Due_On          *time.Time   `json:"due_on",omitempty`
}

func setMilestone(github_milestone github.Milestone)(Milestone){
    tmp := Milestone{}
    tmp.URL = github_milestone.URL
    tmp.Html_Url = github_milestone.URL
    tmp.Number = *github_milestone.Number
    tmp.State = github_milestone.State
    tmp.Title = github_milestone.Title
    tmp.Description = github_milestone.Description
    tmp.Creator = setUser(*github_milestone.Creator)
    tmp.Open_Issues = *github_milestone.OpenIssues
    tmp.Closed_Issues = *github_milestone.ClosedIssues
    tmp.Created_At = github_milestone.CreatedAt
    tmp.Updated_At = github_milestone.UpdatedAt
    tmp.Due_On = github_milestone.DueOn
    return tmp
}
