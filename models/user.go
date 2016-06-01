package models

import "github.com/google/go-github/github"

type User struct {
    Login           *string      `json:"login",omitempty`
    ID              int         `json:"id",omitempty`
    Avatar          *string      `json:"avatar_url",omitempty`
    Gravatar        *string      `json:"gravatar_id",omitempty`
    URL             *string      `json:"url",omitempty`
    Html_Url        *string      `json:"html_url",omitempty`
    Follower_Url    *string      `json:"followers_url",omitempty`
    Following_Url   *string      `json:"following_url",omitempty`
    Gist_Url        *string      `json:"gists_url",omitempty`
    Starred_Url     *string      `json:"starred_url",omitempty`
    Sub_Url         *string      `json:"subscriptions_url",omitempty`
    Orga_Url        *string      `json:"organizations_url",omitempty`
    Repo_Url        *string      `json:"repos_url",omitempty`
    Event_Url       *string      `json:"events_url",omitempty`
    Received_Event  *string      `json:"received_events_url",omitempty`
    Type            *string      `json:"type",omitempty`
    Site_Admin      bool        `json:"site_admin",omitempty`
}

func setUser(github_user github.User)(User){
	tmp := User{}
	tmp.Login = github_user.Login
	tmp.ID = *github_user.ID
	tmp.Avatar = github_user.AvatarURL
	tmp.Gravatar = github_user.GravatarID
	tmp.URL = github_user.URL
	tmp.Html_Url = github_user.HTMLURL
	tmp.Follower_Url = github_user.FollowersURL
	tmp.Following_Url = github_user.FollowingURL
	tmp.Gist_Url = github_user.GistsURL
	tmp.Starred_Url = github_user.StarredURL
	tmp.Sub_Url = github_user.SubscriptionsURL
	tmp.Orga_Url = github_user.OrganizationsURL
	tmp.Repo_Url = github_user.ReposURL
	tmp.Event_Url = github_user.EventsURL
	tmp.Received_Event = github_user.ReceivedEventsURL
	tmp.Type = github_user.Type
	tmp.Site_Admin = *github_user.SiteAdmin
	return tmp
}