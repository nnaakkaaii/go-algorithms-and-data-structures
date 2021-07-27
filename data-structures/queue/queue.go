package queue

import "errors"

type Queue struct {
	head int
	tail int
	max  int
	arr  []interface{}
}

// NewQueue initialize Queue
func NewQueue(max int) *Queue {
	queue := new(Queue)
	queue.head = 0
	queue.tail = 0
	queue.max = max
	queue.arr = make([]interface{}, max)
	return queue
}

// IsEmpty if Queue is empty
func (q *Queue) IsEmpty() bool {
	return q.head == q.tail
}

// IsFull if Full is full
func (q *Queue) IsFull() bool {
	return q.head == (q.tail+1)%q.max
}

// Enqueue adds an element
func (q *Queue) Enqueue(v interface{}) error {
	if q.IsFull() {
		return errors.New("queue overflow")
	}
	q.arr[q.tail] = v
	if q.tail+1 == q.max {
		q.tail = 0
	} else {
		q.tail++
	}
	return nil
}

// Dequeue deletes an element
func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue underflow")
	}
	v := q.arr[q.head]
	if q.head == q.max {
		q.head = 0
	} else {
		q.head++
	}
	return v, nil
}

func (q *Queue) Front() interface{} {
	v := q.arr[q.head]
	return v
}
