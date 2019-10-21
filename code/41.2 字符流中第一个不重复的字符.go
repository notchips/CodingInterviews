package code

import "container/list"

var (
	hash [256]int
	l    = list.New()
)

func InsertChar(c byte) {
	hash[c]++
	l.PushBack(c)
	for e := l.Front(); e != nil; e = e.Next() {
		if hash[e.Value.(byte)] > 0 {
			l.Remove(e)
		}
	}
}

func FirstAppearingOnce() byte {
	if l.Front() == nil {
		return 0
	}
	return l.Front().Value.(byte)
}
