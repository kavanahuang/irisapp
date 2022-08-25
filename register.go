/**
 * Created by Kernel.Huang.
 * User: kernel@live.com
 * Date: 2019/8/21
 * Time: 15:03
 */

package irisapp

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kavanahuang/config"
	"github.com/kavanahuang/log"
	"github.com/kavanahuang/system"
)

type injection struct{}

var Inject = new(injection)

// Application entry.
func Loader(app *iris.Application) *iris.Application {
	Inject.LogService()
	Inject.DebugService(app)
	Inject.ResourceService(app)
	Inject.ViewService(app)
	Inject.MiddlewareService(app)

	return app
}

// Injection logs service
func (i *injection) LogService() {
	err := log.BootLogger()
	if err != nil {
		fmt.Println(err)
	}
}

// Handle static resource dir.
func (i *injection) ResourceService(app *iris.Application) {
	cfg := config.Toml.NewToml("config", "app.toml")
	path := cfg.Read("resource.path").ToStr()
	dir := cfg.Read("resource.dir").ToStr()
	dirPath := system.GetDirPath(dir)
	app.HandleDir(path, dirPath)
}

// Injection view service
func (i *injection) ViewService(app *iris.Application) {
	cfg := config.Toml.NewToml("config", "app.toml")
	dir := cfg.Read("view.dir").ToStr()
	types := cfg.Read("view.types").ToStr()
	dirPath := system.GetDirPath(dir)
	tmpl := iris.HTML(dirPath, types)
	app.RegisterView(tmpl)
}

// Open logger service
func (i *injection) DebugService(app *iris.Application) {
	debugger := config.Toml.NewToml("config", "logs.toml").Zone("log").Get("debug").AtBool()
	if debugger {
		app.Logger().SetLevel("debug")
	}
}

// Injection MiddlewareService service
func (i *injection) MiddlewareService(app *iris.Application) {
	app.Use(recover.New())
	app.Use(logger.New())
}
