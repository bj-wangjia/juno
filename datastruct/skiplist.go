package datastruct

import (
	"github.com/Mintegral-official/juno/document"
	"github.com/Mintegral-official/juno/helpers"
	"math/rand"
	"time"
)

const (
	DefaultMaxLevel    = 12
	DefaultProbability = 0x3FFF
)

type SkipList struct {
	randSource        rand.Source
	header            *Element
	level             int
	length            int
	previousNodeCache [DefaultMaxLevel]*Element
}

func NewSkipList(level int) *SkipList {
	if level < 0 || level > DefaultMaxLevel {
		level = DefaultMaxLevel
	}
	return &SkipList{
		randSource: rand.New(rand.NewSource(time.Now().UnixNano())),
		level:      level,
		length:     0,
		header:     newNode(0, nil, level),
	}
}

func (sl *SkipList) Add(key document.DocId, value interface{}) {
	if m, ok := sl.findGE(key, true, sl.previousNodeCache); ok && m.key == key {
		h := len(m.next)
		x := newNode(key, value, h)
		for i, n := range sl.previousNodeCache[:h] {
			x.setNext(i, m.Next(i))
			n.setNext(i, x)
		}
		return
	}

	h := sl.randLevel()
	x := newNode(key, value, h)
	for i, n := range sl.previousNodeCache[:h] {
		x.setNext(i, n.Next(i))
		n.setNext(i, x)
	}
	sl.length++
}

func (sl *SkipList) Del(key document.DocId) {
	if x, ok := sl.findGE(key, true, sl.previousNodeCache); ok {
		for i, n := range sl.previousNodeCache[:len(x.next)] {
			if n.Next(i) != nil {
				n.setNext(i, n.Next(i).Next(i))
			}
		}
		sl.length--
	}
}

func (sl *SkipList) Contains(key document.DocId) (ok bool) {
	_, ok = sl.findGE(key, true, sl.previousNodeCache)
	return ok
}

func (sl *SkipList) Get(key document.DocId) (x *Element, err error) {
	if x, ok := sl.findGE(key, true, sl.previousNodeCache); ok {
		return x, err
	}
	return nil, helpers.ElementNotfound
}

func (sl *SkipList) Len() int {
	return sl.length
}

func (sl *SkipList) findGE(key document.DocId, flag bool, element [DefaultMaxLevel]*Element) (*Element, bool) {
	for x, h := sl.header, sl.level-1; h >= 0; {
		if x == nil {
			return nil, false
		}
		next, cmp := x.Next(h), 1
		if next != nil {
			cmp = int(next.key - key)
		}
		if cmp < 0 {
			x = next
		} else {
			if flag {
				element[h] = x
				sl.previousNodeCache[h] = element[h]
			} else if cmp == 0 {
				return next, true
			}
			if h == 0 {
				return next, cmp == 0
			}
			h--
		}
	}
	return nil, false
}

func (sl *SkipList) findLT(key document.DocId) (*Element, bool) {
	for x, h := sl.header, sl.level-1; h >= 0; {
		next := x.Next(h)
		if next == nil || next.key >= key {
			if h == 0 {
				if x == sl.header {
					return nil, false
				}
				return x, true
			}
			h--
		} else {
			x = next
		}
	}
	return nil, false
}

func (sl *SkipList) randLevel() int {
	l := 1
	for ((sl.randSource.Int63() >> 32) & 0xFFFF) < DefaultProbability {
		l++
	}
	if l > DefaultMaxLevel || l <= 0 {
		l = DefaultMaxLevel
	}
	return l
}

func (sl *SkipList) Iterator() *SkipListIterator {
	return NewSkipListIterator(ElementCopy(sl.header))
}
