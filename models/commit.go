package models

// Commit represents a github project branch commit
type Commit struct {
    SHA string `json:"sha"`
    URL string `json:"url"`
}