package models

// Contributor represents a github project contributor
type Contributor struct {
    Login         string `json:"login"`
    ID            int    `json:"id"`
    AvatarURL     string `json:"avatar_url"`
    URL           string `json:"url"`
    HTMLURL       string `json:"html_url"`
    Type          string `json:"type"`
    Contributions int    `json:"contributions"`
}