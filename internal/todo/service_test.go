package todo

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	mocks "github.com/wimspaargaren/go-training-cli/internal/mocks/db_mocks"
	"github.com/wimspaargaren/go-training-cli/internal/repository"
)

func TestCreateTODOSuccess(t *testing.T) {
	// when
	todoStoreMock := mocks.NewTODOStore(t)
	todoStoreMock.On("Create", mock.Anything).Return(repository.TODO{
		ID:          42,
		Title:       "test",
		Description: "test",
	}, nil)

	// given
	todoService := service{
		repository: todoStoreMock,
		logger:     logrus.New(),
	}
	todoItem := &TODO{
		Title:       "test",
		Description: "test",
	}
	err := todoService.Create(
		todoItem,
	)

	// then
	assert.NoError(t, err)
	assert.Equal(t, "test", todoItem.Title)
	assert.Equal(t, "test", todoItem.Description)
	assert.Equal(t, 42, todoItem.ID)
}

func TestCreateTODOError(t *testing.T) {
	// when
	todoStoreMock := mocks.NewTODOStore(t)
	todoStoreMock.On("Create", mock.Anything).Return(repository.TODO{}, assert.AnError)

	// given
	todoService := service{
		repository: todoStoreMock,
		logger:     logrus.New(),
	}
	todoItem := &TODO{
		Title:       "test",
		Description: "test",
	}
	err := todoService.Create(
		todoItem,
	)

	// then
	assert.Error(t, err)
	assert.Equal(t, "test", todoItem.Title)
	assert.Equal(t, "test", todoItem.Description)
	assert.Equal(t, 0, todoItem.ID)
}
