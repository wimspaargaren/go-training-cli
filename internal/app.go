// Package app is the main package that constructs the application.
package app

import (
	"github.com/wimspaargaren/go-training-cli/internal/cli"
)

// Run runs the application.
func Run() error {
	executor, err := cli.NewExecutor()
	if err != nil {
		return err
	}
	return executor.Run()
}
