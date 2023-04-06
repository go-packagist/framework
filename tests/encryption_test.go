package tests

import (
	"github.com/go-packagist/framework/encryption"
	"github.com/go-packagist/framework/foundation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryption(t *testing.T) {
	app := foundation.NewApplication()

	foundation.Facade().Register(&encryption.Provider{
		Container: app.Container,
	})

	ciphertext, _ := encryption.Facade().Encrypt("123456")
	plaintext, _ := encryption.Facade().Decrypt(ciphertext)

	assert.Equal(t, "123456", plaintext)
}
