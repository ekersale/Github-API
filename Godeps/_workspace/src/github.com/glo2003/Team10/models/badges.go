package models

// Badges is the badge container
type Badges map[string]Badge

// Badge is the object that is used to generated badges from repositories data
type Badge struct {
	ImgURI      string `json:"img_uri"`
	Description string `json:"description"`
}

// AllBadges is the list of all badges supported by the software
var AllBadges = Badges{}
