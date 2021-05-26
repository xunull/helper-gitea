package gitea_api

import (
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/xunull/helper-gitea/pkg/gitea_api/gitea_req"
	"github.com/xunull/helper-gitea/pkg/gitea_api/gitea_resp"
	"strconv"
)

func (a *GiteaApi) IsRepoExist(parent, name string) (bool, error) {
	return a.Client.CheckUrlExist(GetRepoUrl(parent, name))
}

func (a *GiteaApi) CreateOrgRepo(org string, data gitea_req.CreateRepoOption) (*gitea_resp.Repository, error) {
	if resp, err := a.Client.PostJson(data).Post(GetOrgRepoUrl(org)); err == nil {
		if resp.StatusCode() == 201 {
			res, err := a.Client.Json(resp, &gitea_resp.Repository{})
			if err == nil {
				repo := res.(*gitea_resp.Repository)
				return repo, nil
			} else {
				return nil, err
			}
		} else if resp.StatusCode() == 403 {
			return nil, errors.New("CreateOrgRepo 403")
		} else {
			return nil, errors.New("CreateOrgRepo Failed: status code:" + strconv.Itoa(resp.StatusCode()))
		}
	} else {
		log.Error().Err(err).Msgf("CreateOrgRepo: %s %s Failed", org, data.Name)
		return nil, err
	}
}

func (a *GiteaApi) GetOrgOneRepo(org string, name string) (*gitea_resp.Repository, error) {

	repos, err := a.getOrganizationRepos(org)
	if err != nil {
		return nil, err
	} else {
		for _, repo := range repos {
			if repo.Name == name {
				return repo, nil
			}
		}
		return nil, nil
	}
}
