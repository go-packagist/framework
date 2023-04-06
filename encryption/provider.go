package encryption

import (
	"github.com/go-packagist/encryption"
	"github.com/go-packagist/framework/container"
	"github.com/go-packagist/framework/provider"
)

type Provider struct {
	*container.Container
	*provider.UnimplementedProvider
}

var _ provider.Provider = (*Provider)(nil)

func (p *Provider) Register() {
	p.Singleton("encrypter", func(c *container.Container) interface{} {
		return encryption.NewEncrypter("EAFBSPAXDCIOGRUVNERQGXPYGPNKYATM")
	})
}
