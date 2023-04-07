package bootstraps

import (
	"github.com/go-packagist/framework/config"
	"github.com/go-packagist/framework/foundation"
	"github.com/spf13/viper"
)

type Configuration struct {
}

var _ Bootstrapper = (*Configuration)(nil)

func (c *Configuration) Bootstrap(app *foundation.Application) {
	app.Instance("config", config.NewRepository(viper.New()))
}
