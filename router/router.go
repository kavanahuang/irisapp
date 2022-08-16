/**
 * Created by Kernel.Huang.
 * User: kernel@live.com
 * Date: 2019/8/21
 * Time: 16:03
 */

package router

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/kavanahuang/irisapp/controller"
)

type router struct {
	ctx             iris.Context
	indexController *controller.IndexController
}

var AppRoute = new(router)

// Initialize routing
func (r *router) NewRouter(app *iris.Application) *iris.Application {
	index := hero.Handler(r.indexController.IndexAction)
	home := hero.Handler(r.indexController.HomeAction)

	app.Handle("GET", "index/{name:string}", index)
	app.Handle("GET", "home/{name:string}", home)

	return app
}
