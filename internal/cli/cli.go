// Package cli provides the command line interface for the application.
package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/wimspaargaren/go-training-cli/internal/todo"
)

// NewExecutor creates a new executor.
func NewExecutor() (*Executor, error) {
	executor := &Executor{}
	err := executor.initialise()
	if err != nil {
		return nil, err
	}

	return executor, nil
}

// Executor is the main executor of the application.
type Executor struct {
	rootCommand *cobra.Command
}

// Run runs the CLI application.
func (e *Executor) Run() error {
	return e.rootCommand.Execute()
}

func (e *Executor) initialise() error {
	rootCommand := e.initRootCommand()
	createCommand, err := e.initCreateCommand()
	if err != nil {
		return fmt.Errorf("unable to create create command: %w", err)
	}
	rootCommand.AddCommand(createCommand)
	e.rootCommand = rootCommand
	return nil
}

func (e *Executor) initRootCommand() *cobra.Command {
	rootCommand := &cobra.Command{
		Use:   "todo",
		Short: "todo app is used to maintain a todo list",
	}
	rootCommand.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	return rootCommand
}

func (e *Executor) initCreateCommand() (*cobra.Command, error) {
	createCommand := &cobra.Command{
		Use:       "create",
		ValidArgs: []string{"title", "description"},
		Short:     "creates a todo item",
		RunE:      executeCommandMiddleware(e.RunCreateCommand),
	}
	createCommand.PersistentFlags().String("title", "", "title of the todo item")
	createCommand.PersistentFlags().String("description", "", "description of the todo item")
	err := createCommand.MarkPersistentFlagRequired("title")
	if err != nil {
		return nil, err
	}
	err = createCommand.MarkPersistentFlagRequired("description")
	if err != nil {
		return nil, err
	}
	return createCommand, nil
}

// ExecutionFunc is a function that executes a command.
type ExecutionFunc func(service todo.Service, cmd *cobra.Command, args []string) error

func executeCommandMiddleware(executionFunc ExecutionFunc) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		todoService, err := serviceForGlobalOpts(cmd)
		if err != nil {
			return err
		}
		return executionFunc(todoService, cmd, args)
	}
}

// RunCreateCommand runs the create command.
func (e *Executor) RunCreateCommand(todoService todo.Service, cmd *cobra.Command, _ []string) error {
	title, err := cmd.Flags().GetString("title")
	if err != nil {
		return err
	}
	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return err
	}
	return todoService.Create(
		&todo.TODO{
			Title:       title,
			Description: description,
		},
	)
}

type globalOpts struct {
	verbose bool
}

func serviceForGlobalOpts(cmd *cobra.Command) (todo.Service, error) {
	globalOpts, err := checkGlobalFlags(cmd)
	if err != nil {
		return nil, err
	}
	opts := []todo.Opts{}
	if globalOpts.verbose {
		opts = append(opts, todo.WithVerbose())
	}
	return todo.NewService(opts...), nil
}

func checkGlobalFlags(cmd *cobra.Command) (*globalOpts, error) {
	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		return nil, err
	}
	return &globalOpts{
		verbose: verbose,
	}, nil
}
