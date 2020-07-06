package queue

import (
	"container/list"
)

type StringQueue struct {
	l *list.List
}

func NewStringQueue() *StringQueue {
	return &StringQueue{
		l: list.New(),
	}
}

func (q *StringQueue) Push(s string) {
	q.l.PushBack(s)
}

func (q *StringQueue) Pop() string {
	if q.l.Len() > 0 {
		var e = q.l.Front()
		var v = e.Value.(string)
		q.l.Remove(e)

		return v
	}
	return ""
}

func (q *StringQueue) Len() int {
	return q.l.Len()
}
