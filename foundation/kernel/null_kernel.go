package kernel

import (
	"github.com/go-packagist/framework/foundation"
	"github.com/go-packagist/framework/foundation/bootstraps"
)

var bootstrappers = []bootstraps.Bootstrapper{
	&bootstraps.Configuration{},
	&bootstraps.Providers{},
}

type NullKernel struct {
	app *foundation.Application
}

func NewNullKernel(app *foundation.Application) *NullKernel {
	return &NullKernel{
		app: app,
	}
}

func (k *NullKernel) Bootstrap() {
	for _, bootstrapper := range bootstrappers {
		bootstrapper.Bootstrap(k.app)
	}
}

func (k *NullKernel) Handle() {
	k.Bootstrap()
}
