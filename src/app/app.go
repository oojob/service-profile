package app

import (
	"context"

	"github.com/oojob/service-profile/src/db"
	"github.com/sirupsen/logrus"
)

// App Application
type App struct {
	Config   *Config
	Database *db.Database
}

// NewContext return new application context
func (a *App) NewContext() *Context {
	return &Context{
		Logger:   logrus.StandardLogger(),
		Database: a.Database,
	}
}

// New Application new instance
func New() (app *App, err error) {
	app = &App{}

	app.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}

	dbConfig, err := db.InitConfig()
	if err != nil {
		return nil, err
	}

	app.Database, err = db.New(dbConfig)
	if err != nil {
		return nil, err
	}

	return app, err
}

// Close close the database
func (a *App) Close() error {
	client := a.Database.Client()
	return client.Disconnect(context.Background())
}
