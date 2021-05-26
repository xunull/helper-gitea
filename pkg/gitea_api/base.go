package gitea_api

import (
	"fmt"
	"github.com/xunull/goc/resty_client"
)

type GiteaApi struct {
	Token  string
	Url    string
	Client *resty_client.RestyClient
}

func NewGiteaApi(token string, url string, ops ...resty_client.RestyOption) *GiteaApi {
	api := GiteaApi{
		Token: token,
		Url:   url,
	}
	ops = append(ops, resty_client.WithLowerTokenHeader(token), resty_client.WithHostUrl(url))
	api.Client = resty_client.NewClient(ops...)
	return &api
}

func RepoWithUsername(username, repo string) string {
	return fmt.Sprintf("-%s-%s", username, repo)
}
