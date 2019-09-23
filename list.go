package list

type Node struct {
	next *Node
	prev *Node
	data interface{}
}

func (n *Node) Value() interface{} {
	return n.data
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) Len() int {
	i := l.head
	length := 0
	for i != nil {
		length += 1
		i = i.next
	}
	return length
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) PushFront(v interface{}) {
	prevHead := l.head
	l.head = &Node{
		next: prevHead,
		prev: nil,
		data: v,
	}
	if prevHead != nil {
		prevHead.prev = l.head
	} else {
		l.tail = l.head
	}
}

func (l *List) PushBack(v interface{}) {
	prevTail := l.tail
	l.tail = &Node{
		next: nil,
		prev: prevTail,
		data: v,
	}
	if prevTail != nil {
		prevTail.next = l.tail
	} else {
		l.head = l.tail
	}
}

func (l *List) Remove(node Node) {
	elm := l.head
	for elm != nil {
		if node == *elm {
			if elm.prev == nil && elm.next == nil {
				l.head = nil
				l.tail = nil
			} else if elm.prev == nil {
				l.head = elm.next
				elm.next.prev = nil
			} else if elm.next == nil {
				l.tail = elm.prev
				elm.prev.next = nil
			} else {
				elm.prev.next = elm.next
				elm.next.prev = elm.prev
			}
			break
		}
		elm = elm.next
	}
}
