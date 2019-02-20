package linkedlist

import (
	"fmt"
	"testing"
)

func TestSingleLinked(t *testing.T) {
	sl := singleLinked{}
	err := sl.addToTail(1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = sl.addToTail(2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = sl.addToTail(5)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	sl.addToDataAfter(100, 2)
	fmt.Println(sl.getDataNode(100).Next)
}

func TestLinkedTrack(t *testing.T) {
	ls := linkedStack{}

	err := ls.push(1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = ls.push(2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = ls.push(3)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(ls.pop())
	err = ls.push("gf")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(ls.pop())
	fmt.Println(ls.pop())
	fmt.Println(ls.pop())
	fmt.Println(ls.pop())
}
