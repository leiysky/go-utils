package orderedmap

import "errors"

var NotFoundErr = errors.New("Not found")

type Map interface {
	Get([]byte) ([]byte, error)
	Put([]byte, []byte)
	Del([]byte)
}

type Iterator interface {
	Valid() bool
	Seek(key []byte)
	Next()
	Key() []byte
	Value() []byte
}
