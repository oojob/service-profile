package api

import (
	"github.com/oojob/service-profile/src/app"
)

// API api base struct
type API struct {
	App    *app.App
	Config *Config
}

// New new api instance
func New(a *app.App) (api *API, err error) {
	api = &API{App: a}

	api.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}

	return api, nil
}
