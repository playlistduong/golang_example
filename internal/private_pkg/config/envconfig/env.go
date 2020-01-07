package envconfig

import (
	"github.com/kelseyhightower/envconfig"
)

var (
	envPrefix = "myapp"
)

func Load(t interface{}) {
	envconfig.Process(envPrefix, t)
}
