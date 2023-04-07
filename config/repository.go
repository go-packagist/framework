package config

import "github.com/spf13/viper"

type Repository struct {
	*viper.Viper
}

func NewRepository(v *viper.Viper) *Repository {
	return &Repository{
		Viper: v,
	}
}
