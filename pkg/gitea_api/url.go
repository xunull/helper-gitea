package gitea_api

import "fmt"

const (
	ListUserRepoUrl   = "/user/repos"
	CreateUserRepoUrl = "/user/repos"
	GetOrgUrl         = "/orgs"
)

func GetOrgRepoUrl(org string) string {
	return fmt.Sprintf("/orgs/%s/repos", org)
}

func GetRepoUrl(parent, name string) string {
	return fmt.Sprintf("/%s/%s", parent, name)
}
