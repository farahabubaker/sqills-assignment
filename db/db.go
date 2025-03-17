package db

// TODO: fill out interface a bit better
type Database interface {
	Create(data any, tableName string) error
	// Update(id int, data any, tableName string) error
	// Delete(id int) error
	// Get(id int) error
	// List() error
}
