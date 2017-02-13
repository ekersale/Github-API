package models

import "fmt"
import "strconv"
import mgithub "github.com/ekersale/Github-API/modules/github"
import "github.com/google/go-github/github"
import "strings"

// Project represents a github project with some additional info
type Project struct {
	ID              string          `json:"id"`
	Name            *string         `json:"name"`
	HTMLURL         *string         `json:"html_url"`
    RepoStatus      string          `json:"hciStatus"`
	Contributors    []Contributor   `json:"contributors"`
	Language        *string         `json:"language"`
	Languages       map[string]int  `json:"languages"`
	Branches        []Branch        `json:"branches"`
    LastCommit      string          `json:"lastCommit"`
	Badges          []string        `json:"badges"`
	OpenIssuesCount int             `json:"open_issues_count"`
    OpenIssues      []Issue         `json:"iopenIssues"`
    CloneURL        *string         `json:"clone_url"`
}

func getOpenIssues(url string, client *github.Client, issnb int) []Issue {
    issueState := "closed"
    split := strings.Split(url, "/")
    nbOpenIssues := 0
    for j := 0; j < issnb; j++ {
        issue, _, err := client.Issues.Get(split[4], split[5], j)
        if (err == nil && *issue.State == issueState) {
            nbOpenIssues++
        }
    }
    if (nbOpenIssues > 5) {
        nbOpenIssues = 5
    }
    var issues = make([]Issue, nbOpenIssues)
    i := 0
    for x := 0; x < issnb; x++ {
       issue, _, err := client.Issues.Get(split[4], split[5], x)
        if (err == nil) {
            if (*issue.State == issueState) {
                issues[i].ID = i;
                issues[i].URL = issue.URL
                issues[i].HTMLURL = issue.HTMLURL
                issues[i].Number = *issue.Number
                issues[i].State = issue.State
                issues[i].Title = issue.Title
                issues[i].Body = issue.Body
                if (issue.User != nil) {issues[i].User = setUser(*issue.User)}
                issues[i].Labels = setLabels(issue.Labels)
                if (issue.Assignee != nil) {issues[i].Assignee = setUser(*issue.Assignee)}
                if (issue.Milestone != nil) {issues[i].Milestone = setMilestone(*issue.Milestone)}
                issues[i].Comments = *issue.Comments
                if (issue.PullRequestLinks != nil) {issues[i].PullRequest = setPullRequest(*issue.PullRequestLinks)}
                issues[i].ClosedAt = *issue.ClosedAt
                issues[i].CreatedAt = *issue.CreatedAt
                issues[i].UpdatedAt = *issue.UpdatedAt
                i++
            }
       }
    }
    return issues
}

func getContributors(url string, client *github.Client) ([] Contributor, error) {
    split := strings.Split(url, "/")
    stats, _, err := client.Repositories.ListContributorsStats(split[4], split[5])
    if (err != nil) {
        return nil, err
    }
    var contrib = make([]Contributor, len(stats))
    var x = 0
    for _, Usr := range stats {
        contrib[x].Login = *Usr.Author.Login
        contrib[x].ID = *Usr.Author.ID
        contrib[x].AvatarURL = *Usr.Author.AvatarURL
        contrib[x].URL = *Usr.Author.URL
        contrib[x].HTMLURL = *Usr.Author.HTMLURL
        contrib[x].Type = *Usr.Author.Type
        contrib[x].Contributions = *Usr.Total
        x++
    }
    return contrib, nil
}

func getBranches(url string, client *github.Client) ([]Branch, error) {
    split := strings.Split(url, "/")
    branch, _, err := client.Repositories.ListBranches(split[4], split[5], nil)
    if (err != nil) {
        return nil, err
    }
    var branches = make([]Branch, len(branch))
    var x = 0
    for _, branch := range branch {
       branches[x].Name = *branch.Name
       branches[x].Commit.SHA = *branch.Commit.SHA
       branches[x].Commit.URL = *branch.Commit.URL
       x++
    }
    return branches, nil
}

func getLanguages(url string, client *github.Client) (map[string]int, error) {
    split := strings.Split(url, "/")
    lang, _, err :=   client.Repositories.ListLanguages(split[4], split[5]);
    if (err != nil) {
        return nil, err
    }
    return lang, nil
}


func getLastCommit(url string, client *github.Client) (string, error) {
    split := strings.Split(url, "/")
    commits, _, err := client.Repositories.ListCommits(split[4], split[5], nil)
   if (err != nil) {
        return "", err
    }
    var listCommits = make([]Commit, len(commits))
    var x = 0
    for _, commits := range commits {
       listCommits[x].SHA = *commits.SHA
       x++
    }
    return listCommits[0].SHA, nil
}

func getRepoStatus(url string, client *github.Client, ref string) (string, error) {
    split := strings.Split(url, "/")
    status, _, err := client.Repositories.ListStatuses(split[4], split[5], ref, nil)
   if (err != nil) {
        return "undefined", err
    }
    var x = 0
    for _, status := range status {
       return *status.State, nil
       x++
    }
    return "undefined", nil
}

func FromRepositories(repos []github.Repository, client *github.Client) ([]Project, error) {
    x := 0
    aproj := make([]Project, len(repos))
    var error error
    for _, repos := range repos {
        tmpproj := Project{}
        tmpproj.ID = strconv.Itoa(*repos.ID)
        tmpproj.Name = repos.FullName
        tmpproj.HTMLURL = repos.HTMLURL
        tmpproj.Contributors, error = getContributors(*repos.ContributorsURL, client)
        tmpproj.Language = repos.Language
        tmpproj.Languages, error = getLanguages(*repos.ContributorsURL, client)
        tmpproj.OpenIssuesCount = *repos.OpenIssuesCount
        tmpproj.Branches, error = getBranches(*repos.BranchesURL, client)
        tmpproj.OpenIssues = getOpenIssues(*repos.ContributorsURL, client, *repos.OpenIssuesCount)
        tmpproj.LastCommit, error = getLastCommit(*repos.ContributorsURL, client)
        tmpproj.RepoStatus, error = getRepoStatus(*repos.ContributorsURL, client, tmpproj.LastCommit)
        tmpproj.CloneURL = repos.CloneURL
        aproj[x] = tmpproj
        x++
    }
	return aproj, error
}

// GetProjects SHOULD returns the list of projects stored in the database
func GetProjects() ([]Project, error) {
    client := mgithub.NewClient()

    if (client == nil) {
        fmt.Println("client not connected")
        return nil, nil
    }
    repos, _, err := client.Repositories.List("", nil)
    if err != nil {
        fmt.Printf("error: %v\n", err)
        return nil, err
    }
	return FromRepositories(repos, client)
}

// UnitTestPassed passed
func UnitTestPassed() int {
    return 1
}

// UnitTestFailed failed
func UnitTestFailed() int {
    return 0
}
