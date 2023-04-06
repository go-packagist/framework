package foundation

import (
	"github.com/go-packagist/framework/provider"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApplication_Register(t *testing.T) {
	app := NewApplication()

	app.Register(&provider.NullProvider{
		Container: app.Container,
	})

	assert.Nil(t, app.MustMake("null"))
}

func TestApplication_Bootstrap(t *testing.T) {
	app := NewApplication()

	app.Register(&provider.NullProvider{
		Container: app.Container,
	})

	app.Boot()
	app.Boot()

	app.Register(&provider.NullProvider{
		Container: app.Container,
	}) // booted and registered
}
