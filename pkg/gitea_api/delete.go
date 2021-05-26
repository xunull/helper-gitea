package gitea_api

import (
	"errors"
	"github.com/rs/zerolog/log"
	"strconv"
)

func (a *GiteaApi) DeleteUserRepo(user, name string) error {

	if resp, err := a.Client.Delete(GetRepoUrl(user, name)); err == nil {
		if resp.StatusCode() == 204 {
			return nil
		} else if resp.StatusCode() == 403 {
			return errors.New("DeleteUserRepo 403")
		} else {
			return errors.New("DeleteUserRepo Failed: status code:" + strconv.Itoa(resp.StatusCode()))
		}
	} else {
		log.Error().Err(err).Msgf("v: %s %s Failed", user, name)
		return err
	}
}
