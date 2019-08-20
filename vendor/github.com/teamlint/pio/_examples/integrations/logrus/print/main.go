package main

import (
	"github.com/teamlint/pio"
	_ "github.com/teamlint/pio/_examples/integrations/logrus"
)

func main() {
	pio.Print("This is an info message that will be printed to the logrus' printer")
}
