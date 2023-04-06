package encryption

import (
	"github.com/go-packagist/encryption"
	"github.com/go-packagist/framework/foundation"
)

func Facade() *encryption.Encrypter {
	if e, ok := foundation.GetInstance().MustMake("encrypter").(*encryption.Encrypter); ok {
		return e
	}

	panic("encryption encrypter not found")
}
