package dbs

type DB[T any] interface {
	Connect(path string) (T, error)
	UpdateOne(key []byte, value []byte) error
	UpdateMany(keyValuePairs any) error
}
