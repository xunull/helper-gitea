package gitea_api

import (
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/xunull/helper-gitea/pkg/gitea_api/gitea_resp"
	"strconv"
)

func (a *GiteaApi) ListOrganization() ([]*gitea_resp.Organization, error) {
	var ret []*gitea_resp.Organization
	for page := 1; ; page++ {

		if resp, err := a.Client.GetWithQueryMap(GetOrgUrl, map[string]string{
			"page":  strconv.Itoa(page),
			"limit": "1000",
		}); err == nil {
			if resp.IsSuccess() {
				var repos []*gitea_resp.Organization
				res, err := a.Client.Json(resp, &repos)
				if err == nil {
					repos := res.(*[]*gitea_resp.Organization)
					if len(*repos) == 0 {
						break
					} else {
						ret = append(ret, *repos...)
					}
				} else {
					return nil, err
				}
			} else {
				return nil, errors.New("GetRepo Failed: status code:" + strconv.Itoa(resp.StatusCode()))
			}
		} else {
			log.Error().Err(err).Msgf("GetOrg Failed, %s", err)
			return nil, err
		}
	}
	return ret, nil
}

func (a *GiteaApi) getOrganizationRepos(name string) ([]*gitea_resp.Repository, error) {
	var ret []*gitea_resp.Repository
	for page := 1; ; page++ {

		if resp, err := a.Client.GetWithQueryMap(GetOrgRepoUrl(name),
			map[string]string{
				"page":  strconv.Itoa(page),
				"limit": "1000",
			}); err == nil {
			if resp.IsSuccess() {

				var repos []*gitea_resp.Repository
				res, err := a.Client.Json(resp, &repos)
				if err == nil {
					repos := res.(*[]*gitea_resp.Repository)
					if len(*repos) == 0 {
						break
					} else {
						ret = append(ret, *repos...)
					}
				} else {
					return nil, err
				}
			} else {
				return nil, errors.New("GetRepo Failed: status code:" + strconv.Itoa(resp.StatusCode()))
			}
		} else {
			log.Error().Err(err).Msgf("GetOrgRepo: %s Repos Failed", name)
			return nil, err
		}
	}
	return ret, nil
}

func (a *GiteaApi) ListOrganizationRepos(name string) ([]*gitea_resp.Repository, error) {
	repos, err := a.getOrganizationRepos(name)
	if err != nil {
		return nil, err
	} else {
		return repos, nil
	}
}
