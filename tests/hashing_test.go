package tests

import (
	"github.com/go-packagist/framework/foundation"
	"github.com/go-packagist/framework/hashing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashing(t *testing.T) {
	app := foundation.NewApplication()

	app.Register(&hashing.Provider{
		Container: app.Container,
	})

	app.Boot()

	assert.True(t,
		hashing.Facade().Check("123456",
			hashing.Facade().MustMake("123456")))
}
