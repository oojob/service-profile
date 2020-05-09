package app

import (
	"github.com/oojob/service-profile/src/model"
)

func (ctx *Context) CreateProfile(profile *model.Profile) (string, error) {
	return ctx.Database.CreateProfile(profile)
}
