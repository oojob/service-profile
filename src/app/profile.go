package app

import (
	profile "github.com/oojob/protorepo-profile-go"
	"github.com/oojob/service-profile/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateProfile :- creates a profile
func (ctx *Context) CreateProfile(profile *model.Profile) (string, error) {
	return ctx.Database.CreateProfile(profile)
}

// UpdateProfile :- updates a profile
func (ctx *Context) UpdateProfile(profile *model.Profile) (string, error) {
	return ctx.Database.UpdateProfile(profile)
}

// Auth :- help's us to authenticate
func (ctx *Context) Auth(in *profile.AuthRequest) (*profile.AuthResponse, error) {
	return ctx.Database.Auth(in)
}

// VerifyToken help's us to verify token
func (ctx *Context) VerifyToken(tokenString string) (*profile.AccessDetails, error) {
	return ctx.Database.VerifyToken(tokenString)
}

// ReadProfile : -read a single profile
func (ctx *Context) ReadProfile(id *primitive.ObjectID) (*model.Profile, error) {
	return ctx.Database.ReadProfile(id)
}

// ValidateUsername :- validates the given username for availability
func (ctx *Context) ValidateUsername(username string) (bool, error) {
	return ctx.Database.ValidateUsername(username)
}

// ValidateEmail :- validates the given email for availability
func (ctx *Context) ValidateEmail(email string) (bool, error) {
	return ctx.Database.ValidateEmail(email)
}
