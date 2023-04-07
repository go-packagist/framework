package hashing

import (
	"github.com/go-packagist/framework/config"
	"github.com/go-packagist/framework/container"
	"github.com/go-packagist/framework/provider"
	"github.com/go-packagist/hashing"
)

type Provider struct {
	*container.Container
	*provider.UnimplementedProvider
}

var _ provider.Provider = (*Provider)(nil)

func (h *Provider) Register() {
	h.Singleton("hash", func(c *container.Container) interface{} {
		return hashing.NewManager(
			config.Facade().Get("hash").(*hashing.Config),
		)
	})
}
