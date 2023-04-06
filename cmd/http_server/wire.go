//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"luwu/apps/hello"
	"luwu/apps/hello/access"
	"luwu/libs/router"
)

type application struct {
	helloApp *access.HelloApp
}

func (app *application) Register(engine *gin.Engine) error {
	router.RegisterSchema("rest", app.helloApp)
	return router.RegisterUrlPatterns(engine)
}

func InitializeApplication() (application, error) {
	wire.Build(
		hello.HelloSet,
		wire.Struct(new(application), "*"),
	)
	return application{}, nil
}

func InitApp(engine *gin.Engine) error {
	app, err := InitializeApplication()
	if err != nil {
		return err
	}
	return app.Register(engine)
}
