package gitea_req

type CreateRepoOption struct {
	Name     string `json:"name"`
	AutoInit bool   `json:"auto_init"`
	Private  bool   `json:"private"`
}
