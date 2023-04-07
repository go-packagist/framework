package bootstraps

import (
	"github.com/go-packagist/framework/foundation"
)

type Providers struct {
}

var _ Bootstrapper = (*Providers)(nil)

func (c *Providers) Bootstrap(app *foundation.Application) {
	app.Boot()
}
