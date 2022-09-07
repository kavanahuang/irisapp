/**
 * Created by Kernel.Huang.
 * User: kernel@live.com
 * Date: 2019/8/26
 * Time: 14:13
 */

package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kavanahuang/status"
)

type IndexController struct {
	data interface{}
}

func (index *IndexController) IndexAction(ctx iris.Context) error {
	status.OK.Data = "index"
	_, err := ctx.JSON(status.OK)
	return err
}

func (index *IndexController) HomeAction(ctx iris.Context) error {
	status.OK.Data = ctx.Params().Get("name")
	_, err := ctx.JSON(status.OK)
	return err
}
