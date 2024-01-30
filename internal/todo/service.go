// Package todo is the domain layer of the application.
package todo

import (
	"github.com/sirupsen/logrus"

	"github.com/wimspaargaren/go-training-cli/internal/repository"
)

// TODO represents a todo item.
type TODO struct {
	ID          int
	Title       string
	Description string
}

// Service is the interface that wraps the basic TODO methods.
type Service interface {
	Create(todo *TODO) error
}

// Options is a struct that contains options for the todo service.
type Options struct {
	Verbose    bool
	Repository repository.TODOStore
}

func defaultOptions() Options {
	return Options{
		Verbose:    false,
		Repository: repository.NewPostgresStore(),
	}
}

// Opts is a function that sets options such as verbose mode.
type Opts func(*Options)

// WithVerbose sets verbose mode.
func WithVerbose() Opts {
	return func(o *Options) {
		o.Verbose = true
	}
}

// NewService creates a new todo service.
func NewService(opts ...Opts) Service {
	options := defaultOptions()
	for _, o := range opts {
		o(&options)
	}
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	if options.Verbose {
		logger.SetLevel(logrus.DebugLevel)
	}
	return &service{
		logger:     logger,
		repository: options.Repository,
	}
}

type service struct {
	logger *logrus.Logger

	repository repository.TODOStore
}

func (s *service) Create(todo *TODO) error {
	// FIXME: implement
	s.logger.Info("create todo")
	s.logger.Debug("create todo debug", todo)
	return nil
}
