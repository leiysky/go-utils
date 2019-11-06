package orderedmap

import (
	"context"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/leiysky/go-utils/assert"
)

func TestSkipList(t *testing.T) {
	assert := assert.New(t)
	m := NewSkipList()

	m.Put([]byte("key1"), []byte("value1"))
	v, _ := m.Get([]byte("key1"))
	assert.Equal(v, []byte("value1"))

	m.Del([]byte("key1"))
	v, err := m.Get([]byte("key1"))
	assert.Equal(err, NotFoundErr)

	m.Put([]byte("key1"), []byte("value1"))
	m.Put([]byte("key1"), []byte("value2"))
	v, _ = m.Get([]byte("key1"))
	assert.Equal(v, []byte("value2"))
}

func generate() []byte {
	length := rand.Intn(1024)
	buff := make([]byte, length)
	for i := 0; i < length; i++ {
		buff[i] = byte(rand.Intn(256))
	}
	return buff
}

func TestSkipListConcurrentWriteRead(t *testing.T) {
	assert := assert.New(t)
	m := NewSkipList()

	ch := make(chan struct {
		Key   []byte
		Value []byte
	}, 65535)

	wg := &sync.WaitGroup{}

	f := func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				k := generate()
				v := generate()
				ch <- struct {
					Key   []byte
					Value []byte
				}{k, v}
				m.Put(k, v)
			}
		}
	}

	data := make(map[string][]byte)

	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 10; i++ {
		go f(ctx)
	}

	go func() {
		wg.Add(1)
		for pair := range ch {
			data[string(pair.Key)] = pair.Value
		}
		wg.Done()
	}()
	time.Sleep(time.Second * 10)
	cancel()
	time.Sleep(time.Second * 5)
	close(ch)
	wg.Wait()

	for k, v := range data {
		value, _ := m.Get([]byte(k))
		assert.Equal(value, v)
	}
}

func BenchmarkSkipList(b *testing.B) {
	m := NewSkipList()
	for i := 0; i < b.N; i++ {
		k := generate()
		v := generate()
		m.Put(k, v)
	}
}

func BenchmarkMap(b *testing.B) {
	m := make(map[string][]byte)
	for i := 0; i < b.N; i++ {
		k := generate()
		v := generate()
		m[string(k)] = v
	}
}

func BenchmarkSyncMap(b *testing.B) {
	m := sync.Map{}
	for i := 0; i < b.N; i++ {
		k := generate()
		v := generate()
		m.Store(string(k), v)
	}
}
