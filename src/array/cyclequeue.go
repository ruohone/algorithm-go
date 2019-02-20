package array

import (
	"errors"
)

type cycleQueue struct {
	Items []int
	Head  int
	Tail  int
	N     int
}

func NewCycleQueue(capacity int) *cycleQueue {
	items := make([]int, capacity)
	return &cycleQueue{Items: items, N: capacity}
}

func (cq *cycleQueue) enqueue(item int) error {
	if (cq.Tail+1)%cq.N == cq.Head {
		return errors.New("queue over")
	}
	cq.Items[cq.Tail] = item
	cq.Tail = (cq.Tail + 1) % cq.N
	return nil
}

func (cq *cycleQueue) dequeue() int {
	if cq.Head == cq.Tail {
		return -1
	}

	ret := cq.Items[cq.Head]
	cq.Head = (cq.Head + 1) % cq.N
	return ret
}
