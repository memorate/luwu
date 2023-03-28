//go:build wireinject
// +build wireinject

package main

import (
	"gin-practice/apps/hello"
	"gin-practice/apps/hello/access"
	"gin-practice/libs/router"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
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
