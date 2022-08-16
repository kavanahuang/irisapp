/*
 Author: Kernel.Huang
 Mail: kernelman79@gmail.com
 File: main
 Date: 8/16/22 2:18 AM
*/
package main

import (
	"github.com/kavanahuang/irisapp"
	"github.com/kavanahuang/irisapp/router"
)

// Iris app call example.
func main() {
	newApp := irisapp.App.New()
	router.AppRoute.NewRouter(newApp.Apply)
	newApp.Run(newApp.Apply)
}
