package linkedlist

type linkedStack struct {
	Items singleLinked
}

func (ls *linkedStack) push(item interface{}) error {
	ls.Items.addToTail(item)
	return nil
}

func (ls *linkedStack) pop() interface{} {
	if ls.Items.Header == nil {
		return nil
	}

	front := ls.Items.Header
	if front.Next == nil {
		ls.Items.Header = nil
		return front.Data
	}

	for front.Next.Next != nil {
		front = front.Next
	}

	d := front.Next
	front.Next = nil
	ls.Items.Tail = front

	return d.Data
}
