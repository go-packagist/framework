package conf

import "time"

type App struct {
	Name     string         `json:"name"`
	Env      string         `json:"env"`
	Debug    bool           `json:"debug"`
	Timezone *time.Location `json:"timezone"`
	Locale   string         `json:"locale"`
	Key      string         `json:"key"`
}
