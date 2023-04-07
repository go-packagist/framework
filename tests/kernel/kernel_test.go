package kernel

import (
	"github.com/go-packagist/framework/foundation"
	"github.com/go-packagist/framework/foundation/bootstraps"
	"github.com/go-packagist/framework/foundation/kernel"
	"github.com/go-packagist/framework/hashing"
	"testing"
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

	// noop
}

func TestNullKernel(t *testing.T) {
	app := foundation.NewApplication()
	app.Register(&hashing.Provider{
		Container: app.Container,
	})

	k := kernel.NewNullKernel(app)

	k.Handle()
}
