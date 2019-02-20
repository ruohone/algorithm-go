package array

import (
	"fmt"
	"testing"
)

func TestCycleQueue(t *testing.T) {

	cq := NewCycleQueue(7)
	cq.enqueue(2)
	cq.enqueue(1)
	cq.enqueue(3)
	cq.enqueue(3)
	fmt.Println(cq.dequeue())
	fmt.Println(cq.dequeue())
	cq.enqueue(4)
	cq.enqueue(0)
	cq.enqueue(7)
	cq.enqueue(8)

	fmt.Println(cq.dequeue())
	fmt.Println(cq.dequeue())
	fmt.Println(cq.dequeue())
	fmt.Println(cq.dequeue())
	fmt.Println(cq.dequeue())
	fmt.Println(cq.dequeue())
}
