// Package repository is the data layer definition of the application.
package repository

// TODO represents a todo item.
type TODO struct {
	ID          int
	Title       string
	Description string
}

// TODOStore is the interface that wraps the basic TODO methods.
type TODOStore interface {
	// Create creates a new TODO item.
	Create(todo TODO) (TODO, error)
	// Get returns a TODO item by id.
	Get(id int) (TODO, error)
	// Update updates a TODO item.
	Update(todo TODO) (TODO, error)
	// Delete deletes a TODO item by id.
	Delete(id int) error
	// List returns a list of TODO items.
	List() ([]TODO, error)
}

// NewPostgresStore creates a new postgres store.
func NewPostgresStore() TODOStore {
	// FIXME: Implement
	return nil
}
