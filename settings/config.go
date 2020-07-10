package settings

import "github.com/kelseyhightower/envconfig"

type Config struct {
	AuthSign string `envconfig:"auth_sign" default:"secret"`
}

var (
	Current = new(Config)
	prefix  = "requiem"
)

func init() {
	_ = envconfig.Process(prefix, Current)
}
