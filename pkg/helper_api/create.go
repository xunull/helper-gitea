package helper_api

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/xunull/goc/commonx"
	"github.com/xunull/goc/git_cmd"
	"github.com/xunull/helper-gitea/pkg/gitea_api/gitea_req"
	"github.com/xunull/helper-gitea/pkg/gitea_api/gitea_resp"
	"github.com/xunull/helper-gitea/pkg/global"
)

func (a *api) CreateUserRepo(name string) *gitea_resp.Repository {
	repo, err := gapi.CreateUserRepo(name)
	commonx.CheckErrOrFatal(err)
	return repo
}

func (a *api) CreateOrgRepo(org string, name string) {
	data := gitea_req.CreateRepoOption{
		Name:     name,
		AutoInit: false,
		Private:  true,
	}
	_, err := gapi.CreateOrgRepo(org, data)
	commonx.CheckErrOrFatal(err)
	fmt.Printf("created success\n")
}

func (a *api) CreateAndAddRemote(name string) {
	_ = a.CreateUserRepo(name)

	api := &git_cmd.GitApi{
		Dir: global.WorkDir,
	}

	url := fmt.Sprintf("%s/%s/%s.git", global.Config.Host, global.Config.DefaultUser, name)
	resp, err := api.AddRemote("origin", url)
	if err != nil {
		log.Fatal().Err(err).Msg(resp)
	}

}
