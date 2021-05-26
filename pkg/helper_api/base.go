package helper_api

import (
	"fmt"
	"github.com/xunull/helper-gitea/pkg/gitea_api"
	"github.com/xunull/helper-gitea/pkg/global"
)

var (
	Api  *api
	gapi *gitea_api.GiteaApi
)

type api struct {
}

func InitDefaultApi() {
	fmt.Printf("%+v\n",global.Config)
	gapi = gitea_api.NewGiteaApi(global.Config.Token, global.Config.ApiHost)
	if global.Config.Proxy != "" {
		gapi.Client.SetProxy(global.Config.Proxy)
	}
}
