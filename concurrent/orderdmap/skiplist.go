package orderedmap

import (
	"bytes"
	"math/rand"
	"sync/atomic"
	"unsafe"
)

const maxLevel = 12

func randomLevel() int {
	return rand.Intn(maxLevel)
}

type up = unsafe.Pointer

type skipList struct {
	head  *node
	level int
}

func NewSkipList() Map {
	return &skipList{
		head:  &node{next: make([]*node, maxLevel)},
		level: 0,
	}
}

func (l *skipList) Get(key []byte) ([]byte, error) {
	itr := NewSkipListIterator(l)
	itr.Seek(key)
	if itr.Valid() && bytes.Equal(itr.Key(), key) {
		return itr.Value(), nil
	}
	return nil, NotFoundErr
}

func (l *skipList) Put(key []byte, value []byte) {
	l.insert(key, value)
}

func (l *skipList) Del(key []byte) {
	l.delete(key)
}

func (l *skipList) insert(key, value []byte) {
	itr := NewSkipListIterator(l).(*skipListIterator)
	itr.Seek(key)
	if itr.Valid() && bytes.Equal(itr.Key(), key) {
		itr.get().value = value
		return
	}

	level := randomLevel()
	n := &node{
		level: level,
		next:  make([]*node, maxLevel),
		key:   key,
		value: value,
	}

	if level > l.level {
		l.level = level
		l.head.level = level
	}

	x := l.head
	for i := level; i >= 0; i-- {
		for {
			if next := (*node)(atomic.LoadPointer((*up)(up(&x.next[i])))); next != nil && bytes.Compare(next.key, n.key) < 0 {
				x = next
			} else {
				break
			}
		}
		n.next[i] = (*node)(atomic.LoadPointer((*up)(up(&x.next[i]))))
		atomic.StorePointer((*up)(up(&x.next[i])), up(n))
	}
}

func (l *skipList) delete(key []byte) {
	itr := NewSkipListIterator(l)
	itr.Seek(key)
	if itr.Valid() && !bytes.Equal(itr.Key(), key) {
		return
	}
	n := itr.(*skipListIterator).get()

	x := l.head
	for i := n.level; i >= 0; i-- {
		for {
			if next := (*node)(atomic.LoadPointer((*up)(up(&x.next[i])))); next != nil && bytes.Compare(next.key, n.key) < 0 {
				x = next
			} else {
				break
			}
		}
		atomic.StorePointer((*up)(up(&x.next[i].next[i])), up(x.next[i]))
	}
}

type node struct {
	level int
	key   []byte
	value []byte

	next []*node
}

type skipListIterator struct {
	cur  *node
	list *skipList
}

func NewSkipListIterator(m Map) Iterator {
	l, ok := m.(*skipList)
	if !ok {
		return nil
	}

	return &skipListIterator{
		list: l,
		cur:  l.head,
	}
}

func (itr *skipListIterator) Valid() bool {
	return itr.cur != nil && itr.cur != itr.list.head
}

func (itr *skipListIterator) Seek(key []byte) {
	x := itr.list.head

	for i := itr.list.level; i >= 0; i-- {
		for {
			if next := (*node)(atomic.LoadPointer((*up)(up(&x.next[i])))); next != nil && bytes.Compare(next.key, key) <= 0 {
				x = next
			} else {
				break
			}
		}
	}

	itr.cur = x
}

func (itr *skipListIterator) Next() {
	if itr.cur == nil {
		return
	}
	itr.cur = itr.cur.next[0]
}

func (itr *skipListIterator) Key() []byte {
	if itr.Valid() {
		return itr.cur.key
	}
	return nil
}

func (itr *skipListIterator) Value() []byte {
	if itr.Valid() {
		return itr.cur.value
	}
	return nil
}

func (itr *skipListIterator) get() *node {
	return itr.cur
}
