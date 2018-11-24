package main

import (
	"github.com/teamlint/iris/_benchmarks/iris-mvc-templates/controllers"

	"github.com/teamlint/iris"
	"github.com/teamlint/iris/context"
	"github.com/teamlint/iris/mvc"
)

const (
	// publicDir is the exactly the same path that .NET Core is using for its templates,
	// in order to reduce the size in the repository.
	// Change the "C\\mygopath" to your own GOPATH.
	// publicDir = "C:\\mygopath\\src\\github.com\\teamlint\\iris\\_benchmarks\\netcore-mvc-templates\\wwwroot"
	publicDir = "/home/teamlint/mygopath/src/github.com/teamlint/iris/_benchmarks/netcore-mvc-templates/wwwroot"
)

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html").Layout("shared/layout.html"))
	app.StaticWeb("/public", publicDir)
	app.OnAnyErrorCode(onError)

	mvc.New(app).Handle(new(controllers.HomeController))

	app.Run(iris.Addr(":5000"))
}

type err struct {
	Title string
	Code  int
}

func onError(ctx context.Context) {
	ctx.ViewData("", err{"Error", ctx.GetStatusCode()})
	ctx.View("shared/error.html")
}
