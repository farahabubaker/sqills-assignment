package db

type Database interface {
	Create(data any, tableName string) error
	Get(data any, tableName string) (any, error)
	List(tableName string) (any, error)
}
