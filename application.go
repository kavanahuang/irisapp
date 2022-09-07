/*
 Author: Kernel.Huang
 Mail: kernelman79@gmail.com
 File: app
 Date: 8/16/22 2:02 AM
*/
package irisapp

import (
	"github.com/kataras/iris/v12"
	"github.com/kavanahuang/config"
	"github.com/kavanahuang/logs"
)

type apps struct {
	Apply *iris.Application
}

var App = new(apps)

// Initialize app
func (app *apps) New() *apps {
	newApp := iris.New()
	Loader(newApp)
	app.Apply = newApp

	return app
}

// Run a app
func (app *apps) Run(newApp *iris.Application) {
	// Begin listening for the application
	port := config.Toml.NewToml("config", "app.toml").Zone("listen").Get("port").AtStr()
	err := newApp.Run(iris.Addr(port), iris.WithoutServerError(iris.ErrServerClosed))
	logs.Error("App service is down, closed port: "+port, err)
}
