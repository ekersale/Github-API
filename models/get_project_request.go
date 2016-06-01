package models

type RequestGetProjects struct {
	Name            *string        	`json:"name"`
	LastCommit		string			`json:"lastCommit"`
	CiStatus		string			`json:"ciStatus"`
	Contributors    []Contributor  	`json:"contributors"`
	Languages       map[string]int 	`json:"languages"`
	Branches        []Branch       	`json:"branches"`
	OpenIssues 		[]Issue 		`json:"open_issues"`
	Badges          []string       	`json:"badges"`
}

func GetProjectsFromRequest(projects []Project) ([]RequestGetProjects) {
	i := 0
	myprojects := make([]RequestGetProjects, len(projects))
	for i < len(projects) {
		tmp := RequestGetProjects{}
		tmp.Name = projects[i].Name
		tmp.LastCommit = projects[i].LastCommit
		tmp.CiStatus = projects[i].RepoStatus
		tmp.Contributors = projects[i].Contributors
		tmp.Languages = projects[i].Languages
		tmp.Branches = projects[i].Branches
		tmp.OpenIssues = projects[i].OpenIssues
		tmp.Badges = projects[i].Badges
		myprojects[i] = tmp
		i++
	}
	return myprojects
}
