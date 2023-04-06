package hashing

import (
	"github.com/go-packagist/framework/foundation"
	"github.com/go-packagist/hashing"
)

func Facade() *hashing.Manager {
	if h, ok := foundation.GetInstance().MustMake("hash").(*hashing.Manager); ok {
		return h
	}

	panic("hashing manager not found")
}
