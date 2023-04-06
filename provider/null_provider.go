package provider

import (
	"fmt"
	"github.com/go-packagist/framework/container"
)

type NullProvider struct {
	*container.Container
}

var _ Provider = (*NullProvider)(nil)

func (p *NullProvider) Register() {
	p.Singleton("null", func(c *container.Container) interface{} {
		return nil
	})
}

func (p *NullProvider) Boot() {
	fmt.Println("null provider boot")
}
