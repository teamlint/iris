// Package main shows how you can add middleware to an mvc Application, simply
// by using its `Router` which is a sub router(an iris.Party) of the main iris app.
package main

import (
	"fmt"
	"time"

	"github.com/teamlint/iris"
	"github.com/teamlint/iris/cache"
	"github.com/teamlint/iris/mvc"
)

var cacheHandler = cache.Handler(10 * time.Second)

var testMid = func(ctx iris.Context) {
	ctx.Application().Logger().Info("test middleware")
	ctx.Next()
}

func main() {
	app := iris.New()
	app.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Info("global middle")
		ctx.ViewData("M", "global middle")
		ctx.Next()
	})
	mvc.Configure(app, configure)

	// http://localhost:8080
	// http://localhost:8080/other
	//
	// refresh every 10 seconds and you'll see different time output.
	app.Run(iris.Addr(":8083"))
}

func configure(m *mvc.Application) {
	m.Router.Use(testMid)
	m.Router.Use(cacheHandler)
	foo := m.Party("/foo")
	foo.Handle(&fooController{})
	m.Handle(&exampleController{
		timeFormat: "Mon, Jan 02 2006 15:04:05",
	})
}

type fooController struct {
}

func (c *fooController) Get() string {
	return fmt.Sprintf("foo controller @ %v ", time.Now())
}

type exampleController struct {
	Context    iris.Context
	timeFormat string
}

func (c *exampleController) Get() string {
	now := time.Now().Format(c.timeFormat)
	return c.Context.GetViewData()["M"].(string) + "last time executed without cache: " + now
}

func (c *exampleController) GetOther() string {
	now := time.Now().Format(c.timeFormat)
	return "/other: " + now
}
