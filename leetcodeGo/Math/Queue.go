package Math

type ArrayQueue struct {
	size  int //len
	array []int
}

func (Queue *ArrayQueue) Push(n int) {
	Queue.array = append(Queue.array, n)
	Queue.size++
}
func (Queue *ArrayQueue) Pop() int {
	if Queue.size == 0 {
		panic("Queue Empty")
	}
	Queue.size--
	Newarray := []int{}
	for i := 1; i < Queue.size; i++ {
		Newarray = append(Newarray, Queue.array[i])
	}
	t := Queue.array[0]
	Queue.array = Newarray
	return t
}
func (Queue *ArrayQueue) IsEmpty() bool {
	if Queue.size == 0 {
		return true
	} else {
		return false
	}
}
func Queue() ArrayQueue {
	t := ArrayQueue{}
	return t
}
