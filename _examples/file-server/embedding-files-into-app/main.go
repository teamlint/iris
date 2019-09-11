package main

import (
	"github.com/teamlint/iris"
)

// Follow these steps first:
// $ go get -u github.com/shuLhan/go-bindata/...
// $ go-bindata ./assets/...
// $ go build
// $ ./embedding-files-into-app
// "physical" files are not used, you can delete the "assets" folder and run the example.
//
// See `file-server/embedding-gziped-files-into-app` example as well.
func newApp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")

<<<<<<< HEAD
	app.StaticEmbedded("/static", "./assets_", Asset, AssetNames)
=======
	app.HandleDir("/static", "./assets", iris.DirOptions{
		Asset:      Asset,
		AssetInfo:  AssetInfo,
		AssetNames: AssetNames,
		ShowList:   true,
	})
>>>>>>> upstream/master

	return app
}

func main() {
	app := newApp()

	// http://localhost:8080/static/css/bootstrap.min.css
	// http://localhost:8080/static/js/jquery-2.1.1.js
	// http://localhost:8080/static/favicon.ico
	app.Run(iris.Addr(":8080"))
}
