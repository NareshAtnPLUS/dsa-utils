package queue

type Queue[T any] interface {
	Enqueue(T)
	Dequeue() (T, bool)
}

/*
		 // Declare intQueue as type Queue[int] and initialize it
	    var intQueue Queue[int] = &GenericQueue[int]{}

		intQueue.Enqueue(10)
	    intQueue.Enqueue(20)


		var stringQueue Queue[string] = &GenericQueue[string]{}

		stringQueue.Enqueue("Hello")
	    stringQueue.Enqueue("World")
*/
type GenericQueue[T any] []T

func (q *GenericQueue[T]) Enqueue(item T) {
	*q = append(*q, item)
}

func (q *GenericQueue[T]) Dequeue() (T, bool) {
	if len(*q) == 0 {
		var zeroValue T // Defaults to zero value of type T
		return zeroValue, false
	}
	firstElement := (*q)[0]
	*q = (*q)[1:]
	return firstElement, true
}
