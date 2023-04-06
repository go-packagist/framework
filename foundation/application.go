package foundation

import (
	"fmt"
	"github.com/go-packagist/framework/container"
	"github.com/go-packagist/framework/provider"
	"sync"
)

var app *Application

type Application struct {
	*container.Container
	providers []provider.Provider
	rw        sync.RWMutex
	booted    bool
}

func NewApplication() *Application {
	app := &Application{
		Container: container.NewContainer(),
		providers: make([]provider.Provider, 10),
	}

	app.bootstrapContainer()

	return app
}

func SetInstance(a *Application) {
	app = a
}

func (a *Application) bootstrapContainer() {
	SetInstance(app)

	a.Instance("app", app)
}

func (a *Application) Register(provider provider.Provider) {
	a.rw.Lock()
	defer a.rw.Unlock()

	provider.Register()

	a.providers = append(a.providers, provider)

	if a.booted {
		provider.Boot()
	}
}

func (a *Application) Boot() {
	if a.booted {
		return
	}

	fmt.Println(a.providers)

	for _, p := range a.providers {
		if p != nil {
			p.Boot()
		}
	}

	a.booted = true
}
