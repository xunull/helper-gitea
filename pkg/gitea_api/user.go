package gitea_api

import (
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/xunull/helper-gitea/pkg/gitea_api/gitea_req"
	"github.com/xunull/helper-gitea/pkg/gitea_api/gitea_resp"
	"strconv"
)

func (a *GiteaApi) ListUserRepo() ([]*gitea_resp.Repository, error) {
	var ret []*gitea_resp.Repository
	for page := 1; ; page++ {

		if resp, err := a.Client.GetWithQueryMap(ListUserRepoUrl,
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
			log.Error().Err(err).Msgf("GetUserRepo Failed")
			return nil, err
		}
	}
	return ret, nil
}

func (a *GiteaApi) CreateUserRepo(name string) (*gitea_resp.Repository, error) {
	data := gitea_req.CreateRepoOption{
		Name:     name,
		AutoInit: false,
		Private:  true,
	}

	if resp, err := a.Client.PostJson(data).Post(CreateUserRepoUrl); err == nil {
		if resp.StatusCode() == 201 {
			res, err := a.Client.Json(resp, &gitea_resp.Repository{})
			if err != nil {
				repo := res.(*gitea_resp.Repository)
				return repo, nil
			} else {
				return nil, err
			}
		} else if resp.StatusCode() == 403 {
			return nil, errors.New("CreateUserRepo 403")
		} else {
			return nil, errors.New("CreateUserRepo Failed: status code:" + strconv.Itoa(resp.StatusCode()))
		}
	} else {
		log.Error().Err(err).Msgf("CreateUserRepo: %s %s Failed", data.Name)
		return nil, err
	}
}
