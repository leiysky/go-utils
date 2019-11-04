package orderedmap

import (
	"bytes"
	"math/rand"
)

const maxLevel = 12

func randomLevel() int {
	return rand.Intn(maxLevel)
}

type skipList struct {
	next     []*node
	curLevel int
}

func (l *skipList) Get(key []byte) ([]byte, error) {
	return nil, nil
}

func (l *skipList) Put([]byte, []byte) error {
	return nil
}

func (l *skipList) Del([]byte) {
}

func (l *skipList) insert(key, value []byte) {
	itr := NewSkipListIterator(l).(*skipListIterator)
	itr.Seek(key)
	if !itr.Valid() {
		return
	}

	cur := itr.get()
	if bytes.Compare(itr.Key(), key) == 0 {
		cur.value = value
		return
	}

	next := cur.next
	level := randomLevel()
	node := &node{
		next: make([]*node, level),
		prev: make([]*node, level),
	}

	if level > l.curLevel {
		for i := l.curLevel; i <= level; i++ {

		}
	}
}

type node struct {
	key   []byte
	value []byte

	prev []*node
	next []*node
}

type skipListIterator struct {
	isValid bool
	cur     *node
	list    *skipList
}

func NewSkipListIterator(m Map) Iterator {
	l, ok := m.(*skipList)
	if !ok {
		return nil
	}

	return &skipListIterator{
		list: l,
	}
}

func (itr *skipListIterator) Valid() bool {
	return itr.isValid
}

func (itr *skipListIterator) Seek(key []byte) {

}

func (itr *skipListIterator) Contains(key []byte) bool {
	itr.Seek(key)
	if itr.isValid {
		if bytes.Compare(key, itr.Key()) == 0 {
			return true
		}
		return false
	}
	return false
}

func (itr *skipListIterator) Next() {

}

func (itr *skipListIterator) Key() []byte {
	if itr.isValid {
		return itr.cur.key
	}
	return nil
}

func (itr *skipListIterator) Value() []byte {
	if itr.isValid {
		return itr.cur.value
	}
	return nil
}

func (itr *skipListIterator) get() *node {
	return itr.cur
}
