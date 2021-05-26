package helper_api

import (
	"fmt"
	"github.com/xunull/goc/commonx"
	"github.com/xunull/helper-gitea/pkg/global"
	"strings"
)

func (a *api) ListOrgRepos(name string) {
	repos, err := gapi.ListOrganizationRepos(name)
	commonx.CheckErrOrFatal(err)
	for _, repo := range repos {
		fmt.Printf("%s\n", repo.Name)
	}
}

func (a *api) ListOrg() {
	orgs, err := gapi.ListOrganization()
	commonx.CheckErrOrFatal(err)
	for _, repo := range orgs {
		fmt.Printf("%s\n", repo.Username)
	}
}

func (a *api) ListUserRepo() {
	repos, err := gapi.ListUserRepo()
	commonx.CheckErrOrFatal(err)
	for _, repo := range repos {
		if global.CommonFlag.UserListRepoOnlyUser {
			temp := strings.Split(repo.FullName, "/")
			if len(temp) > 0 {
				if temp[0] == global.Config.DefaultUser {
					fmt.Printf("%s\n", repo.Name)
				}
			}
		} else {
			fmt.Printf("%s\n", repo.FullName)
		}

	}
}
