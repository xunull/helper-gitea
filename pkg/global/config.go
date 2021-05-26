package global

var Config *AppConfig

type AppConfig struct {
	Token       string
	DefaultUser string
	DefaultOrg  string
	Host        string
	ApiHost     string
	Proxy       string
}

func InitConfig(config *AppConfig) {
	Config = config
}
