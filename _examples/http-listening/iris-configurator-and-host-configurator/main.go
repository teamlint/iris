package main

import (
	"github.com/teamlint/iris/_examples/http-listening/iris-configurator-and-host-configurator/counter"

	"github.com/teamlint/iris"
)

func main() {
	app := iris.New()
	app.Configure(counter.Configurator)

	app.Run(iris.Addr(":8080"))
}
