package linkedlist

import "errors"

type node struct {
	Data interface{}
	Next *node
}

type singleLinked struct {
	Header *node
	Tail   *node
}

func (sl *singleLinked) addToTail(data interface{}) error {
	node := &node{}
	node.Data = data
	if sl.Tail != nil {
		sl.Tail.Next = node
		sl.Tail = node
		return nil
	}

	if sl.Header == nil {
		sl.Header = node
		sl.Tail = node
		return nil
	}

	return errors.New("linked err")
}

func (sl *singleLinked) addToDataAfter(data interface{}, target interface{}) error {
	node := &node{}
	node.Data = data

	if sl.Header == nil {
		return errors.New("data not exist")
	}

	n := sl.getDataNode(target)
	if n == nil {
		return errors.New("data not exist")
	}

	node.Next = n.Next
	n.Next = node
	return nil
}

func (sl *singleLinked) getDataNode(data interface{}) *node {
	if sl.Header == nil {
		return nil
	}

	p := sl.Header
	for p != nil {
		if p.Data == data {
			return p
		}
		p = p.Next
	}

	return nil
}
