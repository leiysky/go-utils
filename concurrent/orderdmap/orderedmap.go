package orderedmap

type Map interface {
	Get([]byte) ([]byte, error)
	Put([]byte, []byte) error
	Del([]byte)
}

type Iterator interface {
	Valid() bool
	Seek(key []byte)
	Contains(key []byte) bool
	Next()
	Key() []byte
	Value() []byte
}
