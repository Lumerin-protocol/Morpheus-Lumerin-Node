package interfaces

type KeyValueStorage interface {
	Get(key string) (string, error)
	Insert(key string, value string) error
	Upsert(key string, value string) error
	Delete(key string) error
	DeleteIfExists(key string) error
}
