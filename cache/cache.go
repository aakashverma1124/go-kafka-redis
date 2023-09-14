package cache

type Cache interface {
	GetKey(key string) (interface{}, error)
	SetKey(key string, value interface{}, ttl int) error
	PutKey(key string, value interface{}) error
	DelKey(key string) error
}
