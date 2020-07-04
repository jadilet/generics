package queue

// Queue data-structure
type Queue []interface{}

// Push a value to the queue
func (q *Queue) Push(data interface{}) {
	*q = append(*q, data)
}

// Front return a first value of queue
func (q *Queue) Front() interface{} {
	if q.IsEmpty() {
		return nil
	}

	return (*q)[0]
}

// Pop the first element of the queue
func (q *Queue) Pop() interface{} {
	if q.IsEmpty() {
		return nil
	}

	val := q.Front()

	*q  = (*q)[1:]
	return val
}

// IsEmpty checks the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}