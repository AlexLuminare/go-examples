package linkedlist

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Insert(value int) {
	newNode := &Node{Value: value}
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

func (l *LinkedList) Delete(value int) {
	if l.head == nil {
		return
	}
	if l.head.Value == value {
		l.head = l.head.Next
		return
	}

	current := l.head
	// for current.Next != nil {
	// 	if current.Next.Value == value {
	// 		current.Next = current.Next.Next
	// 		return
	// 	}
	// 	current = current.Next
	// }
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
		}
		current = current.Next
}

func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}