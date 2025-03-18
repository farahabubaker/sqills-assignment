package db

// TODO: fill out interface a bit better
type Database interface {
	// Initialize()
	Create(data any, tableName string) error
	// Update(id int, data any, tableName string) error
	// Delete(id int) error
	Get(data any, tableName string) (any, error)
	List(tableName string) (any, error)
}
