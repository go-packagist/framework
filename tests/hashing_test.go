package tests

import (
	"github.com/go-packagist/framework/foundation"
	"github.com/go-packagist/framework/hashing"
	h "github.com/go-packagist/hashing"
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
		app.MustMake("hash").(*h.Manager).Check("123456",
			app.MustMake("hash").(*h.Manager).MustMake("123456")))
}
