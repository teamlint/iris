package main

import (
	"time"

	"github.com/teamlint/iris/_examples/structuring/login-mvc-single-responsibility-package/user"

	"github.com/teamlint/iris"
	"github.com/teamlint/iris/mvc"
	"github.com/teamlint/iris/sessions"
)

func main() {
	app := iris.New()
	// You got full debug messages, useful when using MVC and you want to make
	// sure that your code is aligned with the Iris' MVC Architecture.
	app.Logger().SetLevel("debug")

	tmp := iris.HTML("./views", ".html").Layout("shared/layout.html")
	tmp.Reload(true)
	app.RegisterView(tmp)

<<<<<<< HEAD
	app.StaticWeb("/public", "./public")
	app.OnErrorCode(400, func(ctx iris.Context) {
		ctx.Writef("错误截获,错误码:%v, 错误消息:%v", ctx.Values().Get("Title"), ctx.Values().Get("Message"))
	})
=======
	app.HandleDir("/public", "./public")
>>>>>>> upstream/master

	mvc.Configure(app, configureMVC)

	// http://localhost:8080/user/register
	// http://localhost:8080/user/login
	// http://localhost:8080/user/me
	// http://localhost:8080/user/logout
	// http://localhost:8080/user/1
	app.Run(iris.Addr(":8082"), configure)
}

func configureMVC(app *mvc.Application) {
	manager := sessions.New(sessions.Config{
		Cookie:  "sessioncookiename",
		Expires: 24 * time.Hour,
	})

	userApp := app.Party("/user")
	userApp.Register(
		user.NewDataSource(),
		manager.Start,
	)
	userApp.Handle(new(user.Controller))
}

func configure(app *iris.Application) {
	app.Configure(
		iris.WithoutServerError(iris.ErrServerClosed),
	)
}
