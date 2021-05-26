package gitea_resp

type Repository struct {
	Id       int64
	Name     string
	FullName string `json:"full_name"`
	Private  bool
	SshUrl   string `json:"ssh_url"`
	CloneUrl string `json:"clone_url"`
}

type Organization struct {
	Id          int64
	Description string
	AvatarUrl   string `json:"avatar_url"`
	FullName    string `json:"full_name"`
	Location    string
	Username    string
	Visibility  string
	Website     string
}
