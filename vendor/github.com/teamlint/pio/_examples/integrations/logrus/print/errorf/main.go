package main

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/teamlint/pio"
)

func init() {
	// take an output from a print function
	output := pio.OutputFrom.Printf(logrus.Errorf)
	// register a new printer with name "logrus"
	// which will be able to read text and print as string.
	pio.Register("logrus", output).Marshal(pio.Text)
}

func main() {
	for i := 1; i <= 5; i++ {
		<-time.After(time.Second)
		pio.Print(fmt.Sprintf("[%d] This is an error message that will be printed to the logrus' printer", i))
	}
}
