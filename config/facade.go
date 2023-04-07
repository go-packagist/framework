package config

import "github.com/go-packagist/framework/foundation"

func Facade() *Repository {
	if c, ok := foundation.GetInstance().MustMake("config").(*Repository); ok {
		return c
	}

	panic("config repository not found")
}
