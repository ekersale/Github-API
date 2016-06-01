package models

import "github.com/google/go-github/github"

type Label struct {
    URL             *string      `json:"url,omitempty"`
    Name            *string      `json:"name,omitempty"`
    Color           *string      `json:"color,omitempty"`
}

func setLabels(github_labels []github.Label)([]Label) {
	var labels = make([]Label, len(github_labels))
    i := 0
    for i < len(github_labels) {
        labels[i].URL = github_labels[i].URL
        labels[i].Name = github_labels[i].Name
        labels[i].Color = github_labels[i].Color
        i++;
    }
    return labels
}

