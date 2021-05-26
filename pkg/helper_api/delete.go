package helper_api

import (
	"github.com/xunull/goc/commonx"
	"github.com/xunull/helper-gitea/pkg/global"
)

func (a *api) DeleteUserRepo(name string) {
	err := gapi.DeleteUserRepo(global.Config.DefaultUser, name)
	commonx.CheckErrOrFatal(err)
}
