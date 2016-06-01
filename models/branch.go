package models

// Branch represents a github project branch
type Branch struct {
    Name   string
    Commit Commit
}
