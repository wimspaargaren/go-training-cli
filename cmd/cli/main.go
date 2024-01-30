// Package main is the entry point of the CLI application.
package main

import (
	"os"

	"github.com/sirupsen/logrus"

	app "github.com/wimspaargaren/go-training-cli/internal"
)

func main() {
	err := app.Run()
	if err != nil {
		logrus.WithError(err).Errorf("unable to run the application")
		os.Exit(1)
	}
}
