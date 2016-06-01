package models

import "time"

// Issues represents a github project issue
type Issue struct {
    ID              int             `json:"id",omitempty`
    URL             *string          `json:"url",omitempty`
    RepoURL         *string          `json:"repository_url",omitempty`
    LabelURL        *string          `json:"labels_url",omitempty`
    CommentURL      *string          `json:"comments_url",omitempty`
    EventURL        *string          `json:"events_url",omitempty`
    HTMLURL         *string          `json:"html_url",omitempty`
    Number          int             `json:"number",omitempty`
    State           *string          `json:"state",omitempty`
    Title           *string          `json:"title",omitempty`
    Body            *string          `json:"body",omitempty`
    User            User            `json:"user",omitempty`
    Labels          []Label         `json:"labels",omitempty`
    Assignee        User            `json:"assignee",omitempty`
    Milestone       Milestone       `json:"milestone",omitempty`
    Locked          bool            `json:"locked",omitempty`
    Comments        int             `json:"comments",omitempty`
    PullRequest     PullRequest     `json:"pull_request",omitempty`
    ClosedAt        time.Time       `json:"closed_at",omitempty`
    CreatedAt       time.Time       `json:"created_at",omitempty`
    UpdatedAt       time.Time       `json:"updated_at",omitempty`
}
