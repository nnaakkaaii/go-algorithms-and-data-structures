package queue

import "errors"

type Queue struct {
	head int
	tail int
	max int
	arr []int
}

// NewQueue initialize Queue
func NewQueue(max int) *Queue {
	queue := new(Queue)
	queue.head = 0
	queue.tail = 0
	queue.max = max
	queue.arr = make([]int, max)
	return queue
}

// isEmpty if Queue is empty
func (q *Queue) isEmpty() bool {
	return q.head == q.tail
}

// isFull if Full is full
func (q *Queue) isFull() bool {
	return q.head == (q.tail + 1) % q.max
}

// Enqueue adds an element
func (q *Queue) Enqueue(v int) error {
	if q.isFull() {
		return errors.New("queue overflow")
	}
	q.arr[q.tail] = v
	if q.tail + 1 == q.max {
		q.tail = 0
	} else {
		q.tail++
	}
	return nil
}

// Dequeue deletes an element
func (q *Queue) Dequeue() (int, error) {
	if q.isEmpty() {
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
