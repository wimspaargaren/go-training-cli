//go:build mocks

package mocks

//go:generate mockery --output . --filename db_mocks/todo_repository_mock.go --dir ../repository --name TODOStore
